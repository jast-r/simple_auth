package simpleauth

type User struct {
	Id       string `json:"-"`
	FName    string `json:"fname" binding:"required"`
	SName    string `json:"sname" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
