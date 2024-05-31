package auth

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/jwt"
	"github.com/eCanteens/backend-ecanteens/src/helpers/smtp"
	"github.com/eCanteens/backend-ecanteens/src/helpers/upload"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/idtoken"
	"gorm.io/gorm"
)

func checkUniqueService(email string, phone string, id ...uint) error {
	sameUser := checkEmailAndPhone(email, phone, id...)

	if len(*sameUser) > 1 {
		return errors.New("email dan nomor telepon sudah digunakan")
	} else if len(*sameUser) == 1 {
		var fields []string

		if (*sameUser)[0].Email == email {
			fields = append(fields, "email")
		}

		if (*sameUser)[0].Phone != nil && phone != "" {
			if *(*sameUser)[0].Phone == phone {
				fields = append(fields, "nomor telepon")
			}
		}

		errMsg := strings.Join(fields, " dan ") + " sudah digunakan"
		return errors.New(errMsg)
	}

	return nil
}

func verifyGoogleToken(idToken string) (*idtoken.Payload, error) {
	ctx := context.Background()
	payload, err := idtoken.Validate(ctx, idToken, os.Getenv("GOOGLE_CLIENT_ID"))
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func registerService(body *registerScheme) error {
	if err := checkUniqueService(body.Email, body.Phone); err != nil {
		return err
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	user := models.User{
		Avatar:   os.Getenv("BASE_URL") + "/public/assets/avatar-user.jpg",
		Name:     body.Name,
		Email:    body.Email,
		Phone:    &body.Phone,
		Password: string(hashed),
	}

	if err := create(&user); err != nil {
		return err
	}

	return nil
}

func loginService(body *loginScheme) (*models.User, *jwt.UserToken, error) {
	var user models.User

	if err := findByEmail(&user, body.Email); err != nil {
		return nil, nil, errors.New("email atau password salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return nil, nil, errors.New("email atau password salah")
	}

	token := jwt.GenerateUserToken(*user.Id, user.RoleId)

	user.Password = ""
	user.Wallet.Pin = ""

	return &user, token, nil
}

func googleService(body *googleScheme) (*models.User, *jwt.UserToken, error) {
	payload, err := verifyGoogleToken(body.IdToken)

	if err != nil {
		return nil, nil, err
	}

	var user models.User

	if err := findByEmail(&user, payload.Claims["email"].(string)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user.Name = payload.Claims["name"].(string)
			user.Email = payload.Claims["email"].(string)
			user.Avatar = payload.Claims["picture"].(string)

			if err := create(&user); err != nil {
				return nil, nil, err
			}

			if err := findByEmail(&user, payload.Claims["email"].(string)); err != nil {
				return nil, nil, err
			}
		} else {
			return nil, nil, err
		}
	}

	token := jwt.GenerateUserToken(*user.Id, user.RoleId)

	user.Password = ""
	user.Wallet.Pin = ""

	return &user, token, nil
}

func setupGoogleService(body *setupScheme, user *models.User) error {
	if err := checkUniqueService(user.Email, body.Phone, *user.Id); err != nil {
		return err
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	user.Phone = &body.Phone
	user.Password = string(hashed)

	if err := save(user); err != nil {
		return err
	}

	user.Password = ""
	user.Wallet.Pin = ""

	return nil
}

func refreshService(body *refreshScheme) (*jwt.UserToken, error) {
	claim, err := jwt.Parse(body.RefreshToken)
	if err != nil {
		return nil, err
	}

	if claim["type"].(string) != "refresh" {
		return nil, errors.New("token tidak valid")
	}

	var user models.User

	if err := findById(&user, uint(claim["sub"].(float64))); err != nil {
		return nil, err
	}

	token := jwt.GenerateUserToken(*user.Id, user.RoleId)

	return token, nil
}

func forgotService(body *forgotScheme) error {
	var user models.User

	if err := findByEmail(&user, body.Email); err != nil {
		return errors.New("pengguna tidak ditemukan")
	}

	tokenString := jwt.GenerateResetToken(*user.Id)

	return smtp.ResetPasswordTemplate([]string{body.Email}, &smtp.ResetPasswordProps{
		LOGO: fmt.Sprintf("%s/public/assets/logo.png", os.Getenv("BASE_URL")),
		URL:  fmt.Sprintf("%s/api/auth/new-password/%s", os.Getenv("BASE_URL"), tokenString),
		NAME: user.Name,
	})
}

func resetService(body *resetScheme) error {
	claim, err := jwt.Parse(body.Token)

	if err != nil {
		return err
	}

	if claim["type"].(string) != "reset" {
		return errors.New("token tidak valid")
	}

	id, ok := claim["sub"].(float64)
	if !ok {
		return errors.New("cant convert claims")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	user := &models.User{Password: string(hashed)}

	return updatePassword(uint(id), user)
}

func updateProfileService(ctx *gin.Context, user *models.User, body *updateScheme) error {
	if err := checkUniqueService(body.Email, body.Phone, *user.Id); err != nil {
		return err
	}

	user.Name = body.Name
	user.Email = body.Email
	user.Phone = &body.Phone

	if body.Avatar != nil {
		filePath := upload.New(&upload.Option{
			Folder:      "avatar/user",
			Filename:    body.Avatar.Filename,
			NewFilename: strconv.FormatUint(uint64(*user.Id), 10),
		})

		if err := ctx.SaveUploadedFile(body.Avatar, filePath.Path); err != nil {
			return err
		}

		user.Avatar = filePath.Url
	}

	if err := save(user); err != nil {
		return err
	}

	user.Password = ""
	user.Wallet.Pin = ""

	return nil
}

func updatePasswordService(user *models.User, body *updatePasswordScheme) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.OldPassword)); err != nil {
		return errors.New("password salah")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(body.NewPassword), bcrypt.DefaultCost)

	user.Password = string(hashed)

	return save(user)
}

func checkPinService(user *models.User, body *checkPinScheme) error {
	if user.Wallet.Pin == "" {
		return errors.New("pin belum di set")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Wallet.Pin), []byte(body.Pin)); err != nil {
		return errors.New("pin salah")
	}

	return nil
}

func updatePinService(user *models.User, body *updatePinScheme) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(body.Pin), bcrypt.DefaultCost)
	if err != nil {
		return errors.New(err.Error())
	}
	user.Wallet.Pin = string(hashed)

	return save(user.Wallet)
}
