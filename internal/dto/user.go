package dto

type UserSingUpInput struct {
	Name     string `json:"name" binding:"required,min=2,max=64"`
	LastName string `json:"lastName" binding:"required,min=2,max=64"`
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}
