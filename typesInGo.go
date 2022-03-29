package main

import "fmt"

//Insert variables declarations here based on activity
const (
	sentence         = "The following is the account information"
	fName            = "Luke"
	lName            = "Skywalkter"
	age              = 20
	weight           = 73.0
	height           = 1.72
	remainingCredits = 123.55
	accountName      = "admin"
	accountPW        = "password"
	subscribed       = true
)

func tellMeTypes() {
	//insert code here to print out types of variables
	fmt.Printf("%#v is of type %T\n", sentence, sentence)
	fmt.Printf("%#v is of type %T\n", fName, fName)
	fmt.Printf("%#v is of type %T\n", lName, lName)
	fmt.Printf("%d is of type %T\n", age, age)
	fmt.Printf("%.2f is of type %T\n", weight, weight)
	fmt.Printf("%.2f is of type %T\n", height, height)
	fmt.Printf("%.2f is of type %T\n", remainingCredits, remainingCredits)
	fmt.Printf("%#v is of type %T\n", accountName, accountName)
	fmt.Printf("%#v is of type %T\n", accountPW, accountPW)
	fmt.Printf("%t is of type %T\n", subscribed, subscribed)

}

func main() {
	tellMeTypes()
}
