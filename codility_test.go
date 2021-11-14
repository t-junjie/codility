package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	output := Solution(2, "1A 2F 1C")
	fmt.Println(output)
}

var seatToIndex = map[string]int{
	"A": 0,
	"B": 1,
	"C": 2,
	"D": 3,
	"E": 4,
	"F": 5,
	"G": 6,
	"H": 7,
	"J": 8,
	"K": 9,
}

func Solution(N int, S string) int {
	seats := setupSeatingPlan(N, S)
	sum := 0
	for _, row := range seats {
		sum += findMaxInRow(row)
		fmt.Printf("row: %v\n", row)
		fmt.Println("sum:", sum)
	}
	return sum
}

// given a row, find max no of 4-member family that can be seated
func findMaxInRow(row []bool) int {
	// note only 3 possibilties per row
	// 2 if BCDE,FGHJ are not taken
	takenBCDE := row[seatToIndex["B"]] || row[seatToIndex["C"]] || row[seatToIndex["D"]] || row[seatToIndex["E"]]
	takenFGHJ := row[seatToIndex["F"]] || row[seatToIndex["G"]] || row[seatToIndex["H"]] || row[seatToIndex["J"]]
	fmt.Printf("BCDE taken: %v, FGHJ taken:%v\n", takenBCDE, takenFGHJ)
	if takenBCDE && takenFGHJ {
		// test if D & G are both taken - assume BCDE and FGHJ are taken if so
		return 0
	}
	// 1 if DEFG not taken
	if takenBCDE || takenFGHJ {
		// assume either BCDE or FGHJ are taken
		return 1
	}
	// 0 otherwise, all taken
	return 0
}

func setupSeatingPlan(N int, S string) map[int][]bool {
	// Create a seating plan with N rows and 10 seats (A-K)
	seats := make(map[int][]bool)
	r := []bool{}
	// create a row of seats
	for i := 0; i < 10; i++ {
		r = append(r, false)
	}

	for i := 1; i <= N; i++ {
		newrow := make([]bool, 10)
		copy(newrow, r)
		seats[i] = newrow
	}

	// fill in reserved seats
	reserved := strings.Split(S, " ")
	for _, v := range reserved {
		seat := string(v[len(v)-1])
		row, _ := strconv.Atoi(v[:len(v)-1])
		if index, ok := seatToIndex[seat]; ok {
			seats[row][index] = true
		}
	}

	return seats
}

func test(A []int) int {
	// sort array and split into -ve/+ve array using mid as an indicator
	sort.Ints(A)

	// mid := 0
	// for k, v := range A {
	// 	if v > 0 {
	// 		mid = k
	// 		break
	// 	}
	// }

	found := 0
	i, j := 0, len(A)-1
	for i < j {
		if -1*A[i] == A[j] {
			found = A[j]
			break
		}
		if -1*A[i] < A[j] {
			j--
		} else {
			i++
		}
	}

	return found
}
