package forms

type LoginForm struct {
	UserName string `form:"username" json:"username" binding:"required"`
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
}
