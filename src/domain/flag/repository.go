package flag

type Repository interface {
	Create(*Flag) error
	Update(*Flag) error
	Delete(uint64) error
	FindByID(uint64) (*Flag, error)
	FindAll() ([]*Flag, error)
}
