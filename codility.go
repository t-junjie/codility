package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	output := Solution(4, "3D")
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
		return 0
	}
	// 1 if DEFG not taken
	if takenBCDE || takenFGHJ {
		return 1
	}
	// 0 otherwise, all taken
	return 2
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
		if v == "" {
			break
		}
		seat := string(v[len(v)-1])
		row, _ := strconv.Atoi(v[:len(v)-1])
		if index, ok := seatToIndex[seat]; ok {
			seats[row][index] = true
		}
	}

	return seats
}
