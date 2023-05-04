package usecase

import "projeto/app/domain/flag"

type FlagUsecase struct {
	flagRepo flag.Repository
}

func NewFlagUsecase(fr flag.Repository) *FlagUsecase {
	return &FlagUsecase{
		flagRepo: fr,
	}
}

func (fu *FlagUsecase) Create(f *flag.Flag) error {
	if err := fu.flagRepo.Create(f); err != nil {
		return err
	}
	return nil
}

func (fu *FlagUsecase) Update(f *flag.Flag) error {
	if _, err := fu.flagRepo.FindByID(f.ID); err != nil {
		return err
	}
	if err := fu.flagRepo.Update(f); err != nil {
		return err
	}

	return nil
}

func (fu *FlagUsecase) Delete(id uint64) error {
	if _, err := fu.flagRepo.FindByID(id); err != nil {
		return err
	}
	if err := fu.flagRepo.Delete(id); err != nil {
		return err
	}

	return nil
}

func (fu *FlagUsecase) FindByID(id uint64) (*flag.Flag, error) {
	return fu.flagRepo.FindByID(id)
}

func (fu *FlagUsecase) FindAll() ([]*flag.Flag, error) {
	return fu.flagRepo.FindAll()
}
