package vessel

type Repository interface {
	Create(*Vessel) error
	Update(*Vessel) error
	Delete(uint64) error
	FindByID(uint64) (*Vessel, error)
	FindAll() ([]*Vessel, error)
	FindByFlagID(uint64) ([]*Vessel, error)
}
