package main

import (
	"PrimeMultiplicationTable/prime"
	"fmt"
)

type PrimeTable struct {
	tableRows   []string
	tableHeader string
}

func main() {
	table := new(PrimeTable)
	table.GenerateRows(10)
	table.Print()
}

func (t *PrimeTable) Print() {
	fmt.Println(t.tableHeader)
	for _, i := range t.tableRows {
		fmt.Println(i)
	}
}

func (t *PrimeTable) GenerateRows(max int) {
	t.tableHeader = "    | "

	primeNumbers := prime.GetPrimeNumbers(max)
	for _, i := range primeNumbers {
		t.tableHeader = t.tableHeader + formatHeader(i)
		row := formatRow(i)
		for _, j := range primeNumbers {
			product := i * j
			row = row + formatRow(product)
		}
		t.tableRows = append(t.tableRows, row)
	}
}

func formatRow(val int) string {
	var row string
	if val < 10 {
		row = fmt.Sprintf("%d   | ", val)
	} else if val > 99 {
		row = fmt.Sprintf("%d | ", val)
	} else {
		row = fmt.Sprintf("%d  | ", val)
	}

	return row
}

func formatHeader(val int) string {
	var header string
	if val >= 11 {
		header = fmt.Sprintf("%d  | ", val)
	} else {
		header = fmt.Sprintf("%d   | ", val)
	}

	return header

}
