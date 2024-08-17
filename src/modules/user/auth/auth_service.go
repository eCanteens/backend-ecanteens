package auth

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/customerror"
	"github.com/eCanteens/backend-ecanteens/src/helpers/jwt"
	"github.com/eCanteens/backend-ecanteens/src/helpers/smtp"
	"github.com/eCanteens/backend-ecanteens/src/helpers/upload"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/idtoken"
	"gorm.io/gorm"
)

func checkUniqueService(email string, phone string, id ...uint) error {
	sameUser := checkEmailAndPhone(email, phone, id...)

	if len(*sameUser) > 1 {
		return customerror.New("Email dan nomor telepon sudah digunakan", 400)
	} else if len(*sameUser) == 1 {
		var fields []string

		if (*sameUser)[0].Email == email {
			fields = append(fields, "Email")
		}

		if (*sameUser)[0].Phone != nil && phone != "" {
			if *(*sameUser)[0].Phone == phone {
				fields = append(fields, "nomor telepon")
			}
		}

		errMsg := strings.Join(fields, " dan ") + " sudah digunakan"
		return customerror.New(errMsg, 400)
	}

	return nil
}

func verifyGoogleToken(idToken string) (*idtoken.Payload, error) {
	ctx := context.Background()
	payload, err := idtoken.Validate(ctx, idToken, os.Getenv("GOOGLE_CLIENT_ID"))
	if err != nil {
		return nil, customerror.New(err.Error(), 400)
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
		return customerror.GormError(err, "Pengguna")
	}

	return nil
}

func loginService(body *loginScheme) (*models.User, *jwt.UserToken, error) {
	var user models.User

	if err := findByEmail(&user, body.Email); err != nil {
		return nil, nil, customerror.New("Email atau password salah", 400)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return nil, nil, customerror.New("Email atau password salah", 400)
	}

	token := jwt.GenerateUserToken(*user.Id, user.RoleId)

	go create(&models.Token{
		UserId:   *user.Id,
		Token:    token.RefreshToken,
		LastUsed: time.Now(),
	})

	user.Password = ""
	user.Wallet.IsPinSet = user.Wallet.Pin != ""
	user.Wallet.Pin = ""

	return &user, token, nil
}

func logoutService(body *refreshScheme) error {
	if err := deleteToken(body.RefreshToken); err != nil {
		return customerror.New("Anda sudah logout", 400)
	}

	return nil
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
				return nil, nil, customerror.GormError(err, "Pengguna")
			}

			if err := findByEmail(&user, payload.Claims["email"].(string)); err != nil {
				return nil, nil, customerror.GormError(err, "Pengguna")
			}
		} else {
			return nil, nil, err
		}
	}

	token := jwt.GenerateUserToken(*user.Id, user.RoleId)

	go create(&models.Token{
		UserId:   *user.Id,
		Token:    token.RefreshToken,
		LastUsed: time.Now(),
	})

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

	if err := update(user); err != nil {
		return customerror.GormError(err, "Pengguna")
	}

	user.Password = ""
	user.Wallet.Pin = ""

	return nil
}

func refreshService(body *refreshScheme) (*jwt.UserToken, error) {
	var refreshToken models.Token

	if err := findToken(&refreshToken, body.RefreshToken); err != nil {
		return nil, customerror.New("Refresh token tidak valid", 400)
	}

	if refreshToken.User == nil {
		return nil, customerror.New("Pengguna tidak ditemukan", 404)
	}

	if time.Since(refreshToken.LastUsed) > config.App.Auth.RefreshTokenExpiresIn {
		go deleteById(&refreshToken)
		return nil, customerror.New("Refresh token kadaluarsa", 400)
	}

	refreshToken.LastUsed = time.Now()
	token := jwt.GenerateUserToken(*refreshToken.User.Id, refreshToken.User.RoleId)
	refreshToken.Token = token.RefreshToken

	go update(&refreshToken)

	return token, nil
}

func forgotService(body *forgotScheme) error {
	var user models.User

	if err := findByEmail(&user, body.Email); err != nil {
		return customerror.New("pengguna tidak ditemukan", 404)
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
		return customerror.New("Token tidak valid", 400)
	}

	if claim["type"].(string) != "reset" {
		return customerror.New("Token tidak valid", 400)
	}

	id, ok := claim["sub"].(float64)
	if !ok {
		return customerror.New("Token tidak valid", 400)
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	user := &models.User{Password: string(hashed)}

	if err := updatePassword(uint(id), user); err != nil {
		return customerror.GormError(err, "Pengguna")
	}

	return nil
}

func updateProfileService(user *models.User, body *updateScheme) error {
	if err := checkUniqueService(body.Email, body.Phone, *user.Id); err != nil {
		return err
	}

	user.Name = body.Name
	user.Email = body.Email
	user.Phone = &body.Phone

	if body.Avatar != nil {
		filePath, err := upload.New(&upload.Option{
			Folder:      "avatar/user",
			File:        body.Avatar,
			NewFilename: strconv.FormatUint(uint64(*user.Id), 10),
		})

		if err != nil {
			return customerror.New("Gagal saat menyimpan file", 500)
		}

		user.Avatar = filePath.Url
	}

	if err := update(user); err != nil {
		return customerror.GormError(err, "Pengguna")
	}

	user.Password = ""
	user.Wallet.Pin = ""

	return nil
}

func updatePasswordService(user *models.User, body *updatePasswordScheme) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.OldPassword)); err != nil {
		return customerror.New("Password salah", 400)
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(body.NewPassword), bcrypt.DefaultCost)

	user.Password = string(hashed)

	if err := update(user); err != nil {
		return customerror.GormError(err, "Pengguna")
	}

	return nil
}

func checkPinService(user *models.User, body *checkPinScheme) error {
	if user.Wallet.Pin == "" {
		return customerror.New("Pin anda belum di set", 400)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Wallet.Pin), []byte(body.Pin)); err != nil {
		return customerror.New("Pin yang anda masukkan salah", 400)
	}

	return nil
}

func updatePinService(user *models.User, body *updatePinScheme) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(body.Pin), bcrypt.DefaultCost)
	if err != nil {
		return customerror.New("Gagal saat hashing pin", 500)
	}
	user.Wallet.Pin = string(hashed)

	if err := update(user.Wallet); err != nil {
		return customerror.GormError(err, "Wallet")
	}

	return nil
}
