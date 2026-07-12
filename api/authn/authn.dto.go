package authn

// LoginDto defines the field(s) used to DTO input login.
type LoginDto struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// UserRegisterDto defines the field(s) used to DTO input User Register DTO.
type UserRegisterDto struct {
	NameFirst   string `form:"name_first" binding:"required"`
	NameLast    string `form:"name_last" binding:"required"`
	Age         uint   `form:"age" binding:"required"`
	DateOfBirth string `form:"date_of_birth" binding:"required"`
	PlaceBirth  string `form:"place_of_birth" binding:"required"`
	Username    string `form:"username" binding:"required"`
	Password    string `form:"password" binding:"required"`
}
