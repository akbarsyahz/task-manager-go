package authn

type Login struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type UserRegister struct {
	NameFirst   string `form:"name_first" binding:"required"`
	NameLast    string `form:"name_last" binding:"required"`
	Age         uint   `form:"age" binding:"required"`
	DateOfBirth string `form:"date_of_birth" binding:"required"`
	PlaceBirth  string `form:"place_of_birth" binding:"required"`
	Username    string `form:"username" binding:"required"`
	Password    string `form:"password" binding:"required"`
}