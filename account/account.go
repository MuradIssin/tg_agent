package account

import (
	"errors"
	"math/rand/v2"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("1234567890")

type Account struct {
	NameFIO    string    `json:"name"`
	IdTelegram string    `json:"idTelegram"`
	Mobile     string    `json:"mobile"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func NewAccount(name, mobile string) (*Account, error) {
	if name == "" || mobile == "" {
		return nil, errors.New("INVALID_NAME or MOBILE")
	}

	newAcc := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		NameFIO:   name,
		Mobile:    mobile,
		// Password:  password,
	}
	if newAcc.Password == "" {
		newAcc.generatePassword(4)
	}
	return newAcc, nil
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(res)
}

func (acc *Account) Output() {
	color.Cyan("ФИО" + acc.NameFIO)
	color.Cyan("Моб." + acc.Mobile)
	color.Cyan("PIN." + acc.Password)
	color.Green("-----------------------")
}
