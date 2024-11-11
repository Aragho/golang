package main
import "fmt"
func main(){
	 var(milesDriven,gallonUsed int
		totalMiles,totalGallon float64
	)
	for{
		fmt.Print("Enter miles driven(-1 to quit):")
		fmt.Scan(&milesDriven)
		if milesDriven == -1{
			break
		}
		fmt.Print("Enter gallon used:")
		fmt.Scan(&gallonUsed)
		milesPerGallon := float64(milesDriven) / float64(gallonUsed)
		fmt.Printf("Miles per gallon for this trip: %.2f\n",milesPerGallon)

		totalMiles += float64(milesDriven)
		totalGallon += float64(gallonUsed)

		combinedMilesPerGallon := totalMiles / totalGallon
		fmt.Printf("combined miles oer gallon so far: %.2f\n\n",combinedMilesPerGallon)

	}
}