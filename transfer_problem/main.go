package main

import (
    "fmt"
)

type Transaction struct {
    SupplierIndex int
    HowMuch int
}

func main() {
    // ______|customers
    // suppl1
    // suppl2
    // suppl3
    routers := [][]int{
        []int{-1, 2, 3},
        []int{3, -1, -1},
        []int{5, 6, 8},
    }

    customers := []int{12, 87, 45}
    suppliers := []int{4, 7, 34}
    findSolution(routers, suppliers, customers)
}

func findSolution(routers [][]int, suppliers []int, customers []int) {
    solutions := make(map[int][]Transaction)
    n := len(suppliers)
    for i := 0; i<n; i++ {
        solutions[i] = []Transaction{}
    }
    for supplIndex, suplQuantity := range suppliers {
        for j := 0; j < n; j++ {
            cI := findMinRoute(routers[supplIndex])
            if cI == -1 {
                break
            }
            deliverQuantity := customers[cI]
            deliverQuantity -= suplQuantity
            if deliverQuantity <= 0 {
                deliverQuantity = customers[cI]
            }
            customers[cI] -= deliverQuantity
            if customers[cI] == 0 {
                deleteCustomerFromRouters(cI, routers)
            }
            suppliers[supplIndex] -= deliverQuantity
            solutions[cI] = append(solutions[cI], Transaction{SupplierIndex: supplIndex, HowMuch: deliverQuantity})
        }
    }
    fmt.Println("Suppliers", suppliers)
    fmt.Println(solutions)
}

func findMinRoute(routers []int) int {
    var min int
    maxIndex := len(routers) + 1
    for i, route := range routers {
        if route != -1 {
            min = route
            maxIndex = i
            break
        }
    }

    for i, route := range routers {
        if route != -1 && route < min {
            min = route
            maxIndex = i
        }
    }
    if maxIndex > len(routers) {
        return -1
    }
    return maxIndex
}

func deleteCustomerFromRouters(customerIndex int, routers [][]int) {
    for i := 0; i < len(routers); i++ {
        routers[i][customerIndex] = -1
    }
}
