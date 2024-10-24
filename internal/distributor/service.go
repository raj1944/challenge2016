package distributor

import (
	"fmt"
	"strings"
)

var Distributors []*Distributor

func GetDistributorByName(name string) *Distributor {
	for _, distributor := range Distributors {
		if name == distributor.Name {
			return distributor
		}
	}
	return nil
}

func AddDistributor(distributor *Distributor) {
	if CheckDistributor(distributor) {
		Distributors = append(Distributors, distributor)
	} else {
		fmt.Println("Location is already assigned to other distributor")
	}
}

func AddDistributorAsSubDistributor(distributor *Distributor, parentDistribotorName string) {
	var parentDistributor = GetDistributorByName(parentDistribotorName)
	if parentDistributor == nil {
		fmt.Println("No Distributor found with name ", parentDistribotorName)
		return
	}

	for _, include := range distributor.Includes {
		if !parentDistributor.HasPermission(include) {
			fmt.Println("Perent Distributor has no permission for location ", include)
			return
		}
	}

	if !CheckSubDistributor(distributor, parentDistributor) {
		fmt.Println("Location is already assigned to other sub distributor")
	}
	parentDistributor.AddSubDistributor(distributor)
	Distributors = append(Distributors, distributor)
}

func CheckDistributor(newDistributor *Distributor) bool {
	for _, distributor := range Distributors {
		for _, newInclude := range newDistributor.Includes {
			for _, exclude := range distributor.Excludes {
				if strings.HasSuffix(newInclude, exclude) {
					return true
				}
			}

			for _, include := range distributor.Includes {
				if strings.HasSuffix(newInclude, include) {
					return false
				}
			}
		}
	}
	return true
}

func CheckSubDistributor(newDistributor *Distributor, parentDistributor *Distributor) bool {
	for _, distributor := range parentDistributor.SubDistributor {
		for _, newInclude := range newDistributor.Includes {
			for _, exclude := range distributor.Excludes {
				if strings.HasSuffix(newInclude, exclude) {
					fmt.Println(newInclude, exclude)
					return true
				}
			}

			for _, include := range distributor.Includes {
				if strings.HasSuffix(newInclude, include) {
					return false
				}
			}
		}
	}
	return true
}

func CheckDistributorPermission(distributorName string, regionName string) {
	var distributor = GetDistributorByName(distributorName)
	fmt.Print("Distribution Permission: ")
	if distributor.HasPermission(regionName) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func PrintAllDistributor() {
	fmt.Println("#################################################################################")
	for _, distributor := range Distributors {
		fmt.Println("Name: ", distributor.Name)

		fmt.Print("Includes: ")
		for _, include := range distributor.Includes {
			fmt.Print(include, ", ")
		}
		fmt.Println()

		fmt.Print("Excludes: ")
		for _, exclude := range distributor.Excludes {
			fmt.Print(exclude, ", ")
		}
		fmt.Println()
		fmt.Print("Sub Distributor: ")
		for _, distributor := range distributor.SubDistributor {
			fmt.Print(distributor.Name, ", ")
		}
		fmt.Println()
		fmt.Println("#################################################################################")
	}
}
