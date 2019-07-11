package main

import (
	"fmt"
)

// Mutex for AllStudentsRecord and TotalStudentCount
//var StudentMutex = &sync.Mutex{}
func main() {


	var reTry string = "Y"
	for reTry == "Y" || reTry == "y" {
		reTry="n"

		fmt.Print("\n( Enter n to exist). \t \t Please enter Choice. \n 1. Concurrency n Go Routine \n 2. Channel :")
		var choice int
		fmt.Scanf("%d",&choice)
			switch choice {
			case 1:
				concurrencyMain()

			case 2:
				ChannelTest()
			default:
				fmt.Println("You have not Spesified any Choice")

			}

		fmt.Printf("Do you wan to continue (y/n):")
		fmt.Scanf("%s",&reTry)
	}


	//StudentMutex.Unlock()

}

