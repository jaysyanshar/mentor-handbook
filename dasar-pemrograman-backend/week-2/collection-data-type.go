package main

import "fmt"

func main() {
	// array ["A", "B", "C", "", ""] -> length == capacity (capacity harus didefine)
	arr := [5]string{"A", "B", "C"}
	// arr := make([]string, 10, 10)
	arr[4] = "X"
	fmt.Println(arr)

	// slice ["A", "B", "D"] -> length <= capacity (capacity kalau ga didefine jadi unlimited)
	slc := make([]string, 0)
	slc = append(slc, "A", "B", "D", "Hello", "Z") // [0:"A", 1:"B", 2:"D", 3:"Hello", 4:"Z"]
	fmt.Println(slc)

	// map {"Name":"Jays", "Gender":"Male"} -> key bisa fleksibel
	m := make(map[string]string)
	m["Name"] = "Jays"
	m["Gender"] = "Male"
	m["Name"] = "Hello"
	fmt.Println(m)
}

// "Jays,Present,Absent,Sick,Present,Present,Fikri,Present,Present,Present,Present,Sick"
// map[Jays:map[Present:3, Absent:1, Sick:1], Fikri:map[Present:4, Sick:1]]
