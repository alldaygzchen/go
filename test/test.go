package main

import "fmt"

// type emptyMap map[string]string

func main(){
	// fmt.Println("This is a","book")	
	// fmt.Print("This is a","book")	
	// result:= someFunction(1)
	// fmt.Println("the result is:",result)
	// var test = []byte("Hello world")
	// fmt.Println(test)
	// writeMyFile("hello world\n")
	// writeMyFile("abc\n")
	// var testString0 = string("hello world")
	// var testString1 = string(35)
	// var testString2 = string("35")
	// fmt.Println(testString0,testString1,testString2)


	// fmt.Println("emptyMap{}", emptyMap{})
	// fmt.Println("emptyMap{}", emptyMap{}==nil)

	// makeSlice := make([]int, 5)
	// fmt.Println("makeSlice", makeSlice)

	// var a int
	// fmt.Println(a==0)
	// fmt.Println(a==nil)
	var something string
	fmt.Println("Enter something")
	fmt.Scanln(&something)
	fmt.Println("Print something",something)



	

}



// func someFunction(input_int int) (int) {
//     // Some code here
// 	var result int;

//     if input_int==1 {
//         return result// Exit early
//     }

//     // More code here

// 	result = input_int
//     return result 

// }


// func writeMyFile(content string){

// 	file,err := os.OpenFile("contents.txt",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
// 	if err != nil{
// 		fmt.Println("Error opening file",err)
// 		panic("### Error opening file")
// 	}

// 	defer file.Close()

// 	_,err = file.WriteString(content)
// 	if err !=nil {
// 		fmt.Println("Error writing to file",err)
// 		return
// 	}

// }
