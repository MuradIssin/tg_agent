package main

import (
	"fmt"
	"freedomAgent/account"

	"github.com/fatih/color"
)

var menuVariants = []string{
	"1. создать агента",
	"2. найти агента",
	"3. удалить агента",
	"4. выход",
	"Выберите вариант",
}

var menu = map[string]func(*account.Vault){
	"1": createAccount,
	"2": findAccount,
	// "3": findAccountByLogin,
	// "4": deleteAccount,
}

func main() {

	fmt.Println("<<<<<- Контроль агенских записей ->>>>>")
	color.Cyan("-------------------")

	vault := account.NewVault()

	// vault := account
	// account.NewVault(files.NewJsonDb("data.vault"), *encrypter.NewEncrypter())
Menu:
	for {
		variant := promptData(menuVariants...)
		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
	}
}

func promptData(prompt ...string) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}

func createAccount(vault *account.Vault) {
	name := promptData("Введите ФИО ")
	mobile := promptData("Введите мобильный номер агента")
	fmt.Println("был создан аккаунт")
	myAccount, err := account.NewAccount(name, mobile)
	if err != nil {
		fmt.Println(err)
		return
	}
	// vault := account.NewVault()
	vault.AddAccount(*myAccount)
	myAccount.Output()
}

func findAccount(vault *account.Vault) {
	// поиск
	// вывод
	nameAgent := promptData("введите имя агента")
	fmt.Println(nameAgent)
	accounts := vault.FindAccount(nameAgent)
	if len(accounts) == 0 {
		color.Red("не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}
