package auth

import (
	"strconv"
	"strings"
	"time"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/customerror"
	"github.com/eCanteens/backend-ecanteens/src/helpers/jwt"
	"github.com/eCanteens/backend-ecanteens/src/helpers/upload"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	checkUnique(email string, phone string, id ...uint) error
	register(body *registerScheme) error
	login(body *loginScheme) (*models.User, *jwt.UserToken, error)
	logout(body *refreshScheme) error
	refresh(body *refreshScheme) (*jwt.UserToken, error)
	updateProfile(body *updateProfileScheme, user *models.User) error
	updateResto(body *updateRestoScheme, resto *models.Restaurant) error
	updatePassword(user *models.User, body *updatePasswordScheme) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) checkUnique(email string, phone string, id ...uint) error {
	sameUser := s.repo.checkEmailAndPhone(email, phone, id...)

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

func (s *service) register(body *registerScheme) error {
	if err := s.checkUnique(body.Email, body.Phone); err != nil {
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

	if err := s.repo.createResto(&restaurant); err != nil {
		return customerror.GormError(err, "Restoran")
	}

	avatar, err := upload.New(&upload.Option{
		Folder:      "avatar/user",
		File:        body.Avatar,
		NewFilename: strconv.FormatUint(uint64(*user.Id), 10),
	})

	if err != nil {
		return customerror.New("Gagal saat menyimpan file", 500)
	}

	restaurantAvatar, err := upload.New(&upload.Option{
		Folder:      "avatar/restaurant",
		File:        body.RestaurantAvatar,
		NewFilename: strconv.FormatUint(uint64(*restaurant.Id), 10),
	})

	if err != nil {
		return customerror.New("Gagal saat menyimpan file", 500)
	}

	banner, err := upload.New(&upload.Option{
		Folder:      "banner",
		File:        body.Banner,
		NewFilename: strconv.FormatUint(uint64(*restaurant.Id), 10),
	})

	if err != nil {
		return customerror.New("Gagal saat menyimpan file", 500)
	}

	user.Avatar = avatar.Url
	restaurant.Avatar = restaurantAvatar.Url
	restaurant.Banner = banner.Url

	if err := s.repo.updateUser(&user); err != nil {
		return customerror.GormError(err, "Pengguna")
	}

	if err := s.repo.updateResto(&restaurant); err != nil {
		return customerror.GormError(err, "Restoran")
	}

	return nil
}

func (s *service) login(body *loginScheme) (*models.User, *jwt.UserToken, error) {
	var user models.User

	if err := s.repo.findByEmail(&user, body.Email); err != nil {
		return nil, nil, customerror.New("Email atau password salah", 400)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return nil, nil, customerror.New("Email atau password salah", 400)
	}

	token := jwt.GenerateUserToken(*user.Id, user.RoleId)

	go s.repo.createToken(&models.Token{
		UserId:   *user.Id,
		Token:    token.RefreshToken,
		LastUsed: time.Now(),
	})

	return &user, token, nil
}

func (s *service) logout(body *refreshScheme) error {
	if err := s.repo.deleteToken(body.RefreshToken); err != nil {
		return customerror.New("Anda sudah logout", 400)
	}

	return nil
}

func (s *service) refresh(body *refreshScheme) (*jwt.UserToken, error) {
	var refreshToken models.Token

	if err := s.repo.findToken(&refreshToken, body.RefreshToken); err != nil {
		return nil, customerror.New("Refresh token tidak valid", 400)
	}

	if refreshToken.User == nil {
		return nil, customerror.New("Rengguna tidak ditemukan", 400)
	}

	if time.Since(refreshToken.LastUsed) > config.App.Auth.RefreshTokenExpiresIn {
		go s.repo.deleteTokenById(&refreshToken)
		return nil, customerror.New("Refresh token kadaluarsa", 400)
	}

	refreshToken.LastUsed = time.Now()
	token := jwt.GenerateUserToken(*refreshToken.User.Id, refreshToken.User.RoleId)
	refreshToken.Token = token.RefreshToken

	go s.repo.updateToken(&refreshToken)

	return token, nil
}

func (s *service) updateProfile(body *updateProfileScheme, user *models.User) error {
	if err := s.checkUnique(body.Email, body.Phone, *user.Id); err != nil {
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

	if err := s.repo.updateUser(user); err != nil {
		return customerror.GormError(err, "Pengguna")
	}

	return nil
}

func (s *service) updateResto(body *updateRestoScheme, resto *models.Restaurant) error {
	resto.Name = body.Name
	resto.CategoryId = body.CategoryId

	if body.Avatar != nil {
		file, err := upload.New(&upload.Option{
			Folder:      "avatar/restaurant",
			File:        body.Avatar,
			NewFilename: strconv.FormatUint(uint64(*resto.Id), 10),
		})

		if err != nil {
			return customerror.New("Gagal saat menyimpan file", 500)
		}

		resto.Avatar = file.Url
	}

	if body.Banner != nil {
		file, err := upload.New(&upload.Option{
			Folder:      "banner",
			File:        body.Banner,
			NewFilename: strconv.FormatUint(uint64(*resto.Id), 10),
		})

		if err != nil {
			return customerror.New("Gagal saat menyimpan file", 500)
		}

		resto.Banner = file.Url
	}

	if err := s.repo.updateResto(resto); err != nil {
		return customerror.GormError(err, "Restoran")
	}

	return nil
}

func (s *service) updatePassword(user *models.User, body *updatePasswordScheme) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.OldPassword)); err != nil {
		return customerror.New("Password salah", 400)
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(body.NewPassword), bcrypt.DefaultCost)

	user.Password = string(hashed)

	if err := s.repo.updateUser(user); err != nil {
		return customerror.GormError(err, "Pengguna")
	}

	return nil
}
