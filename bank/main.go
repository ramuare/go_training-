package main

import (
	"fmt"
	"math/rand"
)

var name, emailid string
var DOB, i int
var passpowrd int
var phonenumber int
var cash, balance, NewAmount, witcash uint
var bal int
var x int

func Accountcreation(name string, emailid string, DOB int, password int, phonenumber int) {

	fmt.Println("enter name:")
	fmt.Scan(&name)

	fmt.Println("enter emailid:")
	fmt.Scan(&emailid)

	fmt.Println("enter DOB:")
	fmt.Scan(&DOB)

	fmt.Println("password:")
	fmt.Scan(&password)

	fmt.Println("phonenumber:")
	fmt.Scan(&phonenumber)

	x := func() {

		RandomIntegerwithinRange := (rand.Intn(9999999-1000000) + 1000000)
		fmt.Println(RandomIntegerwithinRange)
	}
	fmt.Println(x)

}

func Deposit(cash uint, balance uint) {
	fmt.Println("your previouse balace is:", balance)

	NewAmount = cash + balance
	fmt.Println("and your current balance is:", NewAmount)
}

func Withdraw(witcash uint, bal int) {

	fmt.Printf("your ammount is %v initially", bal)

	wit := bal - int(witcash)

	fmt.Println("your current ammount is:", wit)
	if wit < 0 {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("your balance is", wit, "Please add the money")
			}
		}()
		panic("your do not have enough balance")
	}
}

func Details() {

	fmt.Println("enter account numbe:" )


}

func logout() {

}

func main() {
	fmt.Println("welcome to the bank system")
	fmt.Println("1.Accountcreation")
	fmt.Println("2.Deposit")
	fmt.Println("3.Withdraw")
	fmt.Println("4.Details")
	fmt.Println("5.logout")

	fmt.Println("enter your options:", i)
	fmt.Scan(&i)
	switch i {
	case 1:
		Accountcreation(name, emailid, DOB, passpowrd, phonenumber)
		return
	case 2:
		Deposit(cash, balance)
		return
	case 3:
		Withdraw(witcash, bal)
		return
	case 4:
		Details()
	case 5:
		logout()

	}

	var dw string
	fmt.Scanln(&dw)
	var amou uint
	fmt.Println("Enter the ammount")
	fmt.Scanln(&amou)

	if dw == "Deposit" || dw == "deposit" {
		Deposit(amou, 0)

	} else if dw == "Withdraw" || dw == "withdraw" {
		Withdraw(amou, 0)
	}

}
