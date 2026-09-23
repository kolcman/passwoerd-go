package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*?-_+=")

type account struct {
	login    string
	password string
	url      string
}

func (account *account) generatePassword(n int) {
	password := make([]rune, n)
	for i := range password {
		password[i] = alphabet[rand.IntN(len(alphabet))]
	}
	account.password = string(password)
}

func newAccount(login string, password string, urlString string) (*account, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("некорректный адрес")
	}
	return &account{
		login:    login,
		password: password,
		url:      urlString,
	}, nil
}

func main() {
	login := promptData("Введите логин: ")
	urlString := promptData("Введите URL: ")

	myAccount, err := newAccount(login, "", urlString)
	if err != nil {
		fmt.Println(err)
		return
	}
	myAccount.generatePassword(10)
	fmt.Print(myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)
	return res
}
