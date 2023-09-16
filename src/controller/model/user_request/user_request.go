package user_request

type User_request struct {
	Email string `json:"email" binding:"required,min=10,max=150"`
	Password string `json:"password" binding:"required,min=10,max=250"`
	Name string `json:"name" binding:"required,min=2,max=100"`
}