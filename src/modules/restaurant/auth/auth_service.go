package auth

import (
	"errors"
	"strconv"
	"strings"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/jwt"
	"github.com/eCanteens/backend-ecanteens/src/helpers/upload"
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

func registerService(body *registerScheme) error {
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

	if err := create(&restaurant); err != nil {
		return err
	}

	avatar, err := upload.New(&upload.Option{
		Folder:      "avatar/user",
		File:        body.Avatar,
		NewFilename: strconv.FormatUint(uint64(*user.Id), 10),
	})

	if err != nil {
		return err
	}

	restaurantAvatar, err := upload.New(&upload.Option{
		Folder:      "avatar/restaurant",
		File:        body.RestaurantAvatar,
		NewFilename: strconv.FormatUint(uint64(*restaurant.Id), 10),
	})

	if err != nil {
		return err
	}

	banner, err := upload.New(&upload.Option{
		Folder:      "banner",
		File:        body.Banner,
		NewFilename: strconv.FormatUint(uint64(*restaurant.Id), 10),
	})

	if err != nil {
		return err
	}

	user.Avatar = avatar.Url
	restaurant.Avatar = restaurantAvatar.Url
	restaurant.Banner = banner.Url

	if err := update(&user); err != nil {
		return err
	}

	return update(&restaurant)
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

func updateProfileService(body *updateProfileScheme, user *models.User) error {
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
			return err
		}

		user.Avatar = filePath.Url
	}

	if err := update(user); err != nil {
		return err
	}

	user.Password = ""
	user.Wallet.Pin = ""

	return nil
}
