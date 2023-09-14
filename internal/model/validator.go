package model

import (
	"log"
	"regexp"
)

func (u Users) Validate() map[string]string {
	errs := make(map[string]string)
	if len(u.Login) == 0 {
		errs["login"] = "Введите логин"
	}
	if _, ok := errs["login"]; !ok {
		if len(u.Login) > 16 {
			errs["login"] = "Логин не должен быть больше 16 символов"
		}
		if len(u.Login) < 3 {
			errs["login"] = "Логин не должен быть меньше 3 символов"
		}
	}
	if _, ok := errs["login"]; !ok {
		m, err := regexp.MatchString("^[a-z0-9_-]+$", u.Login)
		if err != nil {
			log.Fatal(err)
		}
		if !m {
			errs["login"] = "Логин должен содержать a-z, 0-9, символы _,-"
		}
	}

	if len(u.Password) == 0 {
		errs["password"] = "Введите пароль"
	}
	if _, ok := errs["password"]; !ok && len(u.Password) <= 6 {
		errs["password"] = "Минимальная длина 6"
	}
	if u.RoleID == 0 {
		errs["password"] = "Выберите роль"
	}

	if len(u.WorkPlace) == 0 && u.RoleID == 4 {
		errs["company_ids"] = "Выберите предприятие"
	}
	return errs
}
