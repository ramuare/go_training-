package main

import (
	"fmt"
)

var name,emailid string
var DOB,i int
var passpowrd int
var phonenumber int
var cash, balance, NewAmount, witcash uint
var bal int

func welcomeoptions(){

	

}



func Accountcreation(name string, emailid string ,DOB int,password int,phonenumber int){

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


}

func DepositFun(cash uint, balance uint) {
	fmt.Println("your previouse balace is:", balance)

		NewAmount = cash + balance
	fmt.Println("and your current balance is:", NewAmount)
}

func Withdraw(witcash uint , bal int) {

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

func Details(){


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
		switch i{
		case 1:
			fmt.Println("")
		case 2:

		case 3:
			
		case 4:

		case 5:	
		
		}

		






	

	

	var dw string
	fmt.Scanln(&dw)
	var amou uint
	fmt.Println("Enter the ammount")
	fmt.Scanln(&amou)

	if dw == "Deposit" || dw == "deposit" {
		DepositFun(amou, 0)
		
	} else if dw == "Withdraw" || dw == "withdraw" {
		Withdraw(amou, 0)
	}

}
