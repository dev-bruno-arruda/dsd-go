package usecase

import (
	"errors"
	"projeto/domain/vessel"
)

type VesselUsecase struct {
	vesselRepo vessel.Repository
}

func NewVesselUsecase(vr vessel.Repository) *VesselUsecase {
	return &VesselUsecase{
		vesselRepo: vr,
	}
}

func (vu *VesselUsecase) Create(v *vessel.Vessel) error {
	if _, err := vu.vesselRepo.FindByFlagID(v.FlagID); err == nil {
		return errors.New("vessel already assigned to a flag")
	}

	if err := vu.vesselRepo.Create(v); err != nil {
		return err
	}

	return nil
}

func (vu *VesselUsecase) Update(v *vessel.Vessel) error {
	if _, err := vu.vesselRepo.FindByID(v.ID); err != nil {
		return err
	}

	if err := vu.vesselRepo.Update(v); err != nil {
		return err
	}
	return nil
}

func (vu *VesselUsecase) Delete(id uint64) error {
	if _, err := vu.vesselRepo.FindByID(id); err != nil {
		return err
	}
	if err := vu.vesselRepo.Delete(id); err != nil {
		return err
	}

	return nil
}

func (vu *VesselUsecase) FindByID(id uint64) (*vessel.Vessel, error) {
	return vu.vesselRepo.FindByID(id)
}

func (vu *VesselUsecase) FindAll() ([]*vessel.Vessel, error) {
	return vu.vesselRepo.FindAll()
}

func (vu *VesselUsecase) FindByFlagID(flagID uint64) ([]*vessel.Vessel, error) {
	return vu.vesselRepo.FindByFlagID(flagID)
}
