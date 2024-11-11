package main
import "fmt"
func main(){
	var counter,number,largest int
	largest = -1 << 31 


	for counter = 1; counter <= 10; counter++{
		fmt.Printf("Enter number %d:",counter)
		fmt.Scan(&number)
		if number > largest{
			largest = number
		}
	}
	fmt.Printf("The largest number is: %d\n", largest)
}	

