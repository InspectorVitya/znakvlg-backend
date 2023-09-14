package model

import "strings"

func (u *Users) Prepare() {
	u.Login = strings.ToLower(u.Login)
	u.Email = strings.ToLower(u.Email)
}
