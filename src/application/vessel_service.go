package application

import (
	"errors"

	"projeto/app/domain/flag"
	"projeto/app/domain/vessel"
)

type VesselService struct {
	vesselRepository vessel.Repository
	flagRepository   flag.Repository
}

func NewVesselService(vr vessel.Repository, fr flag.Repository) *VesselService {
	return &VesselService{
		vesselRepository: vr,
		flagRepository:   fr,
	}
}

func (vs *VesselService) Create(v *vessel.Vessel) error {
	err := vs.validateFlag(v.FlagID)
	if err != nil {
		return err
	}

	return vs.vesselRepository.Create(v)
}

func (vs *VesselService) Update(v *vessel.Vessel) error {
	err := vs.validateFlag(v.FlagID)
	if err != nil {
		return err
	}

	return vs.vesselRepository.Update(v)
}

func (vs *VesselService) Delete(id uint64) error {
	return vs.vesselRepository.Delete(id)
}

func (vs *VesselService) FindByID(id uint64) (*vessel.Vessel, error) {
	v, err := vs.vesselRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (vs *VesselService) FindAll() ([]*vessel.Vessel, error) {
	v, err := vs.vesselRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (vs *VesselService) validateFlag(id uint64) error {
	f, err := vs.flagRepository.FindByID(id)
	if err != nil {
		return err
	}

	if f == nil {
		return errors.New("Flag not found")
	}

	return nil
}
