package main

import (
	"bufio"
	"challenge2016/internal/distributor"
	"fmt"
	"os"
	"strings"
)

func SplitRegion(input string) []string {
	regions := strings.Split(input, " ")
	for i, region := range regions {
		regions[i] = strings.ToUpper(region)
	}
	return regions
}

func CreateNewDistributor() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Distributor Name")
	var name string
	fmt.Scanln(&name)
	var isSubDistributor string

	// var regionsIncludes string
	fmt.Println("Enter Include region like region1 region2")
	scanner.Scan()
	var includes = SplitRegion(scanner.Text())

	// var regionsExclude string
	fmt.Println("Enter Exclude region like region1 region2")
	scanner.Scan()
	var excludes = SplitRegion(scanner.Text())

	var d = distributor.NewDistributor(name, includes, excludes)

	fmt.Println("Is Sub Distributor (enter yes or no)")
	fmt.Scanln(&isSubDistributor)
	if isSubDistributor == "yes" {
		var subDistributorName string
		fmt.Println("Enter Perent Distributor Name")
		fmt.Scanln(&subDistributorName)
		distributor.AddDistributorAsSubDistributor(d, subDistributorName)
	} else {
		distributor.AddDistributor(d)
	}
}

func checkPermission() {
	fmt.Println("Enter Distributor Name")
	var name string
	fmt.Scanln(&name)

	var region string
	fmt.Println("Enter region name")
	fmt.Scanln(&region)

	distributor.CheckDistributorPermission(name, region)

}

func main() {
	// var d1 = distributor.NewDistributor("DISTRIBUTOR1", []string{"INDIA", "UNITEDSTATES"}, []string{"KARNATAKA-INDIA", "CHENNAI-TAMILNADU-INDIA"})
	// var d2 = distributor.NewDistributor("DISTRIBUTOR2", []string{"INDIA"}, []string{"TAMILNADU-INDIA"})
	// distributor.AddDistributor(d1)
	// distributor.AddDistributorAsSubDistributor(d2, "DISTRIBUTOR1")
	// distributor.PrintAllDistributor()
	// distributor.CheckDistributorPermission("DISTRIBUTOR2", "CHICAGO-ILLINOIS-UNITEDSTATES")

	var num int
	for {
		fmt.Println("To Create Distributor: 1")
		fmt.Println("To Check Permission: 2")
		fmt.Println("To Print All Distributor: 3")
		fmt.Println("To Exit: 4")

		fmt.Scanln(&num)

		if num == 1 {
			CreateNewDistributor()
		} else if num == 2 {
			checkPermission()
		} else if num == 3 {
			distributor.PrintAllDistributor()
		} else {
			break
		}
	}
}
