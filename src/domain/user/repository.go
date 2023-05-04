package user

type Repository interface {
	Create(*User) error
	Update(*User) error
	Delete(uint64) error
	FindByID(uint64) (*User, error)
	FindAll() ([]*User, error)
	FindByEmail(string) (*User, error)
}
