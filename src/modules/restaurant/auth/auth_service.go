package auth

import (
	"errors"
	"strconv"
	"strings"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/jwt"
	"github.com/eCanteens/backend-ecanteens/src/helpers/upload"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

func registerService(ctx *gin.Context, body *registerScheme) error {
	if err := checkUniqueService(body.Email, body.Phone); err != nil {
		return err
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Phone:    &body.Phone,
		Password: string(hashed),
		RoleId:   3,
	}

	restaurant := models.Restaurant{
		Name:       body.RestaurantName,
		CategoryId: body.CategoryId,
		Owner:      &user,
	}

	avatar := upload.New(&upload.Option{
		Folder: "avatar/user",
		Filename: body.Avatar.Filename,
		NewFilename: strconv.FormatUint(uint64(*user.Id.Id), 10),
	})

	restaurantAvatar := upload.New(&upload.Option{
		Folder: "avatar/restaurant",
		Filename: body.RestaurantAvatar.Filename,
		NewFilename: strconv.FormatUint(uint64(*user.Id.Id), 10),
	})

	banner := upload.New(&upload.Option{
		Folder: "banner",
		Filename: body.Avatar.Filename,
		NewFilename: strconv.FormatUint(uint64(*user.Id.Id), 10),
	})

	if err := ctx.SaveUploadedFile(body.Avatar, avatar.Path); err != nil {
		return err
	}

	if err := ctx.SaveUploadedFile(body.RestaurantAvatar, restaurantAvatar.Path); err != nil {
		return err
	}

	if err := ctx.SaveUploadedFile(body.Banner, banner.Path); err != nil {
		return err
	}

	user.Avatar = &avatar.Url
	restaurant.Avatar = restaurantAvatar.Url
	restaurant.Banner = banner.Url

	return create(&user)
}

func loginService(body *loginScheme) (*models.User, *jwt.UserToken, error) {
	var user models.User

	if err := findByEmail(&user, body.Email); err != nil {
		return nil, nil, errors.New("email atau password salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return nil, nil, errors.New("email atau password salah")
	}

	token := jwt.GenerateUserToken(*user.Id.Id, user.RoleId)

	user.Password = ""
	user.Wallet.Pin = ""

	return &user, token, nil
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

	token := jwt.GenerateUserToken(*user.Id.Id, user.RoleId)

	return token, nil
}
