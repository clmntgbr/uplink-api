package dto

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterInput struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	FirstName string `json:"firstName" validate:"required,min=2"`
	LastName  string `json:"lastName" validate:"required,min=2"`
}

type LoginOutput struct {
	Token string     `json:"token"`
	User  UserOutput `json:"user"`
}

type RegisterOutput struct {
	Token string     `json:"token"`
	User  UserOutput `json:"user"`
}
