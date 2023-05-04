package application

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"projeto/app/domain/user"
)

type UserService struct {
	userRepository user.Repository
}

func NewUserService(ur user.Repository) *UserService {
	return &UserService{
		userRepository: ur,
	}
}

func (us *UserService) Create(u *user.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	return us.userRepository.Create(u)
}

func (us *UserService) Update(u *user.User) error {
	return us.userRepository.Update(u)
}

func (us *UserService) Delete(id uint64) error {
	return us.userRepository.Delete(id)
}

func (us *UserService) FindByID(id uint64) (*user.User, error) {
	u, err := us.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (us *UserService) FindAll() ([]*user.User, error) {
	u, err := us.userRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (us *UserService) Authenticate(email string, password string) (*user.User, error) {
	u, err := us.userRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return nil, errors.New("Invalid password")
	}

	return u, nil
}
