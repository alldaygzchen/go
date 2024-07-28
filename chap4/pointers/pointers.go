package main

import "fmt"


func main() {
	pointerEx1()
	fmt.Print("#################\n")
	pointerEx2()
}


// Directly mutate values
func pointerEx2() {
	age := 32 // regular variable
	agePointer := &age

	fmt.Println("Age Pointer:", *agePointer)

	editAdultYears(agePointer)
	fmt.Println("After *agePointer", *agePointer)
	fmt.Println("After age", age)
}

func editAdultYears(agePointer *int){
	*agePointer = *agePointer - 18
}

// Avoid unnecessary value copies
func pointerEx1() {
	age := 32 
	agePointer := &age
	fmt.Println("Age Pointer:", *agePointer)

	adultYears := getAdultYearsEx1(agePointer)
	fmt.Println("After adultYears", adultYears)
	fmt.Println("After *agePointer:", *agePointer )
	fmt.Println("After age", age)
}

func getAdultYearsEx1(agePointer *int) int {
	return *agePointer - 18
}
