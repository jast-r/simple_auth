package repository

type Authorization struct {
}

type CompanyList struct {
}

type Repository struct {
	Authorization
	CompanyList
}

func NewRepository() *Repository {
	return &Repository{}
}
