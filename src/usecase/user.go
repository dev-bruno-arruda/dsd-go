package usecase

import (
	"errors"

	"DSD/projeto/domain/user"
)

type UserUsecase struct {
	userRepo user.Repository
}

func NewUserUsecase(ur user.Repository) *UserUsecase {
	return &UserUsecase{
		userRepo: ur,
	}
}

func (uu *UserUsecase) Create(u *user.User) error {
	if _, err := uu.userRepo.FindByEmail(u.Email); err == nil {
		return errors.New("user already exists")
	}

	if err := uu.userRepo.Create(u); err != nil {
		return err
	}

	return nil
}

func (uu *UserUsecase) Update(u *user.User) error {
	if _, err := uu.userRepo.FindByID(u.ID); err != nil {
		return err
	}

	if err := uu.userRepo.Update(u); err != nil {
		return err
	}

	return nil
}

func (uu *UserUsecase) Delete(id uint64) error {
	if _, err := uu.userRepo.FindByID(id); err != nil {
		return err
	}

	if err := uu.userRepo.Delete(id); err != nil {
		return err
	}

	return nil
}

func (uu *UserUsecase) FindByID(id uint64) (*user.User, error) {
	return uu.userRepo.FindByID(id)
}

func (uu *UserUsecase) FindAll() ([]*user.User, error) {
	return uu.userRepo.FindAll()
}
