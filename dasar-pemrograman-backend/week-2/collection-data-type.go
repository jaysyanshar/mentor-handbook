package main

import "fmt"

func main() {
	// array ["A", "", "Hello", "", "X"] -> length == capacity (capacity harus didefine)
	arr := [5]string{}
	// arr := make([]string, 10, 10)
	arr[0] = "A"
	arr[4] = "X"
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	for i, v := range arr {
		fmt.Println(i, v)
	}

	// slice ["A", "B", "D", "C", ""] -> length <= capacity (capacity kalau ga didefine jadi unlimited)
	slc := make([]string, 0)
	slc = append(slc, "X", "Y", "Z")
	slc = append([]string{"A", "B"}, slc...) // [0:"A", 1:"B", 2:"D", 3:"Hello", 4:"Z"]
	fmt.Println(slc)
	fmt.Println(len(slc))
	fmt.Println(cap(slc))
	for i, v := range slc {
		fmt.Println(i, v)
	}

	// map {"Name":"Jays", "Gender":"Male"} -> key bisa fleksibel
	m := make(map[string]string)
	m["Name"] = "Jays"
	m["Gender"] = "Male"
	m["Name"] = "Hello"
	fmt.Println(m)
	for i, v := range m {
		fmt.Println(i, v)
	}

	m2 := make(map[string][]any)
	m2["Group 1"] = []any{1, 2, 3, "4"}
	m2["Group 2"] = []any{5, 6, 7, 8, 9, 0}
	fmt.Println(m2)
}

// "Jays,Present,Absent,Sick,Present,Present,Fikri,Present,Present,Present,Present,Sick"
// map[Jays:map[Present:3, Absent:1, Sick:1], Fikri:map[Present:4, Sick:1]]
