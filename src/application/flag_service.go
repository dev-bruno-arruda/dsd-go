package application

import (
	"errors"

	"projeto/app/domain/flag"
	"projeto/app/domain/vessel"
)

type FlagService struct {
	flagRepository   flag.Repository
	vesselRepository vessel.Repository
}

func NewFlagService(fr flag.Repository, vr vessel.Repository) *FlagService {
	return &FlagService{
		flagRepository:   fr,
		vesselRepository: vr,
	}
}

func (fs *FlagService) Create(f *flag.Flag) error {
	return fs.flagRepository.Create(f)
}

func (fs *FlagService) Update(f *flag.Flag) error {
	return fs.flagRepository.Update(f)
}

func (fs *FlagService) Delete(id uint64) error {
	// check if the flag is assigned to a vessel before deleting
	vessels, err := fs.vesselRepository.FindByFlagID(id)
	if err != nil {
		return err
	}
	if len(vessels) > 0 {
		return errors.New("Flag is assigned to a vessel and cannot be deleted")
	}

	return fs.flagRepository.Delete(id)
}

func (fs *FlagService) FindByID(id uint64) (*flag.Flag, error) {
	f, err := fs.flagRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs *FlagService) FindAll() ([]*flag.Flag, error) {
	f, err := fs.flagRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return f, nil
}
