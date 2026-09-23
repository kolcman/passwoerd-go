package main

import "fmt"

func main() {
	login := promptData("Введите логи")
	password := promptData("Введите логи")
	url := promptData("Введите URL")

	fmt.Print(login, password, url)
}

func promptData(prompt string) string {

	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)
	return res
}
