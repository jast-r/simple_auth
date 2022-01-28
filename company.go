package simpleauth

type Company struct {
	Id   string `json:"-"`
	Name string `json:"name"`
	INN  string `json:"INN"`
}
