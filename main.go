package main

import (
	"fmt"
	"math/rand/v2"
)

var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*?-_+=")

type account struct {
	login    string
	password string
	url      string
}

func main() {
	fmt.Println(generatePassword(10))
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url := promptData("Введите URL: ")

	myAccount := account{
		login,
		password,
		url,
	}
	fmt.Print(myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)
	return res
}

func generatePassword(n int) string {
	password := make([]rune, n)
	for i := range password {
		password[i] = alphabet[rand.IntN(len(alphabet))]
	}
	return string(password)
}
