package api

import (
	"github.com/nightlegend/consulapi/core/data/userdata"
)

func Login(account *userdata.Accounts) bool {
	res := userdata.FindOne(account)
	return res
}

func Register(account *userdata.Accounts) bool {
	res := userdata.Add(account)
	return res
}
