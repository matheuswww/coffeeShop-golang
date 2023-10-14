package admin_auth_request

type Admin_signIn_request struct {
	Email    string `json:"email" binding:"required,min=10,max=150"`
	Password string `json:"password" binding:"required,min=10,max=256"`
}

type Admin_signUp_request struct {
	Email    string `json:"email" binding:"required,min=10,max=150"`
	Password string `json:"password" binding:"required,min=10,max=256"`
}

type Admin_InsertProduct_request struct {
	Email    string `json:"email" binding:"required,min=10,max=150"`
	Password string `json:"password" binding:"required,min=10,max=256"`
}
