package repository

type Company interface {
}

type Customer interface {
}

type General interface {
}

type Repository struct {
	Company
	Customer
	General
}

func NewRepo() *Repository {
	return &Repository{
		Company:  newCompany(),
		Customer: newCustomer(),
		General:  newGeneral(),
	}
}
