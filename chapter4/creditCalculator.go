package main 
import "fmt"

func main(){
	var accountNumber int
	var balanceAtStart,totalCharges,totalCredits, creditLimit,newBalance float64

	fmt.Print("Enter account number: ")
	fmt.Scan(&accountNumber)

	fmt.Print("Enter balance at the beginning of the month: ")
	fmt.Scan(&balanceAtStart)

	fmt.Print("Enter total charges this month: ")
	fmt.Scan(&totalCharges)

	fmt.Print("Enter total credits this month: ")
	fmt.Scan(&totalCredits)

	fmt.Print("Enter allowed credit limit: ")
	fmt.Scan(&creditLimit)

	newBalance = balanceAtStart + totalCharges - totalCredits

	fmt.Printf("New balance for account %d: %.2f\n", accountNumber, newBalance)

	if newBalance > creditLimit {
		fmt.Println("Credit limit exceeded!")
	} else {
		fmt.Println("Credit limit not exceeded.")
	}

}