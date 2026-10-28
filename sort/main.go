package main

import (
	"fmt"
	"sort"
)

var (
	intValue    = []int{3, 2, 4, 1, 5}
	floatValue  = []float64{3.1, 2.2, 4.4, 1.1, 5.5}
	stringValue = []string{"c", "b", "a"}
	family      = []Person{
		{"Alice", 23},
		{"Eve", 2},
		{"David", 2},
		{"Bob", 25},
	}
)

type Person struct {
	Name string
	Age  int
}

func sortBasic() {
	sort.Ints(intValue)
	sort.Float64s(floatValue)
	sort.Strings(stringValue)
	fmt.Printf("%v", intValue)
	fmt.Printf("%v", floatValue)
	fmt.Printf("%v", stringValue)
}

func sortSlice() {
	sort.Slice(family, func(i, j int) bool { return family[i].Age < family[j].Age })
	fmt.Printf("%v", family)
}

func reverseSlice() {
	sort.Sort(sort.Reverse(sort.Float64Slice(floatValue)))
}

func main() {
	sortBasic()
	//fmt.Println("")
	//sortSlice()
	reverseSlice()
	//fmt.Printf("\n%v", floatValue)
}
