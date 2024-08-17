package admin

import (
	"os"
	"strconv"
	"strings"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/enums"
	"github.com/eCanteens/backend-ecanteens/src/helpers/customerror"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
	"github.com/eCanteens/backend-ecanteens/src/helpers/upload"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	adminLogin(body *adminLoginScheme) (*models.User, *string, error)
	dashboard() (map[string]interface{}, error)
	checkWallet(phone string) (*models.User, error)
	topupWithdraw(phone string, body *topupWithdrawScheme, tipe string) (*models.Transaction, error)
	transaction(id string) (*models.Transaction, error)
	mutasi(query *mutationQS) (*pagination.Pagination[models.Transaction], error)
	updateAdminProfile(user *models.User, body *updateAdminProfileScheme) error
	updateAdminPassword(user *models.User, body *updateAdminPasswordScheme) error
	checkUnique(email string, id ...uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) adminLogin(body *adminLoginScheme) (*models.User, *string, error) {
	var user models.User

	if err := s.repo.findAdminEmail(&user, body.Email); err != nil {
		return nil, nil, customerror.New("email admin salah", 400)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return nil, nil, customerror.New("password admin salah", 400)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   *user.Id,
		"email": user.Email,
		"exp":   0,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return nil, nil, err
	}

	user.Password = ""

	return &user, &tokenString, nil
}

func (s *service) dashboard() (map[string]interface{}, error) {
	var userCount int64
	var restaurantCount int64

	if err := s.repo.count("users", &userCount); err != nil {
		return nil, customerror.GormError(err, "Pengguna")
	}

	if err := s.repo.count("restaurants", &restaurantCount); err != nil {
		return nil, customerror.GormError(err, "Restoran")
	}

	data := map[string]interface{}{
		"users":       userCount,
		"restaurants": restaurantCount,
		"total":       userCount + restaurantCount,
	}

	return data, nil
}

func (s *service) checkWallet(phone string) (*models.User, error) {
	var user models.User

	if err := s.repo.findUser(&user, phone); err != nil {
		return nil, customerror.GormError(err, "Pengguna")
	}

	return &user, nil
}

func (s *service) topupWithdraw(phone string, body *topupWithdrawScheme, tipe string) (*models.Transaction, error) {
	var user models.User

	if err := s.repo.findUser(&user, phone); err != nil {
		return nil, customerror.GormError(err, "Pengguna")
	}

	if err := s.repo.topupWithdraw(body.Amount, &user, tipe); err != nil {
		return nil, err
	}

	data, err := s.repo.createTransaction(&user, body.Amount, enums.TransactionType(tipe))

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *service) transaction(id string) (*models.Transaction, error) {
	var transaction models.Transaction

	data, err := s.repo.findTransaction(&transaction, id)

	if err != nil {
		return nil, customerror.GormError(err, "Transaksi")
	}

	return data, nil
}

func (s *service) mutasi(query *mutationQS) (*pagination.Pagination[models.Transaction], error) {
	var result = pagination.New(models.Transaction{})

	if err := s.repo.findMutasi(result, query); err != nil {
		return nil, customerror.GormError(err, "Mutasi")
	}

	return result, nil
}

func (s *service) updateAdminProfile(user *models.User, body *updateAdminProfileScheme) error {
	if err := s.checkUnique(body.Email, *user.Id); err != nil {
		return err
	}

	user.Name = body.Name
	user.Email = body.Email

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

	if err := s.repo.save(user); err != nil {
		return customerror.GormError(err, "Pengguna")
	}

	user.Password = ""

	return nil
}

func (s *service) updateAdminPassword(user *models.User, body *updateAdminPasswordScheme) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.OldPassword)); err != nil {
		return customerror.New("Password salah", 400)
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(body.NewPassword), bcrypt.DefaultCost)

	user.Password = string(hashed)

	if err := s.repo.save(user); err != nil {
		return customerror.GormError(err, "Pengguna")
	}

	return nil
}

func (s *service) checkUnique(email string, id ...uint) error {
	sameUser := s.repo.checkEmail(email, id...)

	if len(*sameUser) > 1 {
		return customerror.New("Email dan nomor telepon sudah digunakan", 400)
	} else if len(*sameUser) == 1 {
		var fields []string

		if (*sameUser)[0].Email == email {
			fields = append(fields, "email")
		}

		errMsg := strings.Join(fields, " dan ") + " sudah digunakan"
		return customerror.New(errMsg, 400)
	}

	return nil
}
