package userdata

type Accounts struct {
	USERNAME string `json:"userName" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
}
