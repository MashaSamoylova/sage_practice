package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

var Routers = [][]int{
	[]int{8, 3, 1},
	[]int{4, 7, 4},
	[]int{5, 2, 6},
}

var PotentialSumms = [][]int{
	[]int{0, 0, 0},
	[]int{0, 0, 0},
	[]int{0, 0, 0},
	[]int{0, 0, 0},
}

var Plan = [][]int{
	[]int{-1, -1, -1},
	[]int{-1, -1, -1},
	[]int{-1, -1, -1},
	[]int{-1, -1, -1},
}

var Suppliers = []int{30, 90, 50}
var Customers = []int{70, 60, 30}

var U = []int{0xff, 0xff, 0xff}
var V = []int{0xff, 0xff, 0xff, 0xff}

func main() {
	A := 0
	for _, q := range Suppliers {
		A += q
	}
	B := 0
	for _, q := range Customers {
		B += q
	}
	if A > B {
		addFakeCustomer(A - B)
	}
	makeScratchPlan()
	fmt.Println("Price: ", calculatePlanPrice())
	calculateTransportPotential()
	calculateSummPotentials()
	fmt.Println("=================")
	spew.Dump(Routers)
	fmt.Println("=================")
	spew.Dump(PotentialSumms)

}

func addFakeCustomer(diff int) {
	Routers = append(Routers, make([]int, len(Routers[0])))
	Customers = append(Customers, diff)
}

func makeScratchPlan() {
	for !allCustomersAreSatisfy() {
		i, j := findMinRoute()
		amount := calculateHowMuch(Suppliers[j], Customers[i])
		Customers[i] -= amount
		Suppliers[j] -= amount
		Plan[i][j] = amount
	}
	j := findRemain()
	i := len(Plan) - 1
	amount := calculateHowMuch(Suppliers[j], Customers[i])
	Customers[i] -= amount
	Suppliers[j] -= amount
	Plan[i][j] = amount
}

func findMinRoute() (int, int) {
	min := Routers[0][0]
	var minI, minJ int
	for i, _ := range Routers {
		for j, _ := range Routers[i] {
			// ignore fake customer
			if Routers[i][j] < min && Routers[i][j] != 0 && Suppliers[j] > 0 && Customers[i] > 0 {
				min = Routers[i][j]
				minI, minJ = i, j
			}
		}
	}
	return minI, minJ
}

func calculateHowMuch(SuplQuantity int, CustomerNeed int) int {
	tmp := CustomerNeed - SuplQuantity
	if tmp > 0 {
		return SuplQuantity
	}
	return CustomerNeed
}

func allCustomersAreSatisfy() bool {
	for i := 0; i < len(Customers)-1; i++ {
		if Customers[i] > 0 {
			return false
		}
	}
	return true
}

func findRemain() int {
	for i, q := range Suppliers {
		if q > 0 {
			return i
		}
	}
	return -1
}

func calculatePlanPrice() int {
	price := 0
	for i := 0; i < len(Plan); i++ {
		for j := 0; j < len(Plan[0]); j++ {
			if Plan[i][j] != -1 {
				price += Plan[i][j] * Routers[i][j]
			}
		}
	}
	return price
}

func calculateTransportPotential() {
	for !allPotentialAreCalculated() {
		i, j := findMaxRoute()
		summ := Routers[i][j]
		if U[j] == 0xff && V[i] == 0xff {
			U[j] = summ / 2
			V[i] = summ - U[j]
			continue
		}
		if U[j] == 0xff {
			U[j] = summ - V[i]
			continue
		}
		if V[i] == 0xff {
			V[i] = summ - U[j]
			continue
		}
	}

}

func findMaxRoute() (int, int) {
	max := -1
	var maxI, maxJ int
	for i := 0; i < len(Routers); i++ {
		for j := 0; j < len(Routers[i]); j++ {
			// ignore fake customer
			if Routers[i][j] > max && Plan[i][j] != -1 && (U[j] == 0xff || V[i] == 0xff) {
				max = Routers[i][j]
				maxI, maxJ = i, j
			}
		}
	}
	return maxI, maxJ
}

func allPotentialAreCalculated() bool {
	for i := 0; i < len(Routers); i++ {
		for j := 0; j < len(Routers[i]); j++ {
			if (U[j] == 0xff || V[i] == 0xff) && Plan[i][j] != -1 {
				return false
			}
		}
	}
	return true
}

func calculateSummPotentials() {
	for i := 0; i < len(Routers); i++ {
		for j := 0; j < len(Routers[i]); j++ {
			PotentialSumms[i][j] = V[i] + U[j]
		}
	}
}
