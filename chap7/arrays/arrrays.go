package main

import "fmt"

func main() {
	// var productNames [4]string = [4]string{"A Book"}
	productNames:= [4]string{"A Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

	fmt.Println(productNames)

	fmt.Println("prices",prices)
	

	// slice share same memory address
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	featuredPrices := prices[1:3]
	fmt.Println("featurePrices",featuredPrices)
	featuredPrices[0]=100
	fmt.Println("modify prices",prices)
	fmt.Println("modify featurePrices",featuredPrices)

	fmt.Println("prices[0] memory address",&prices[1])
	fmt.Println("prices[1] memory address",&prices[2])
	fmt.Println("featuredPrices[0] memory address",&featuredPrices[0])
	fmt.Println("featuredPrices[1] memory address",&featuredPrices[1])

	
	// slice length and capacity
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	
	fmt.Println("len of 1234",len("1234"))
	// can resize to the right not to the left
	fmt.Println("len and cap of featuredPrices",len(featuredPrices),cap(featuredPrices))
	highLightedPrice := featuredPrices[:1]
	fmt.Println("highLightedPrice",highLightedPrice)
	fmt.Println("len and cap of highLightedPrice",len(highLightedPrice),cap(highLightedPrice))
	// can resize to the right not to the left
	highLightedPrice =highLightedPrice[:3]
	fmt.Println("since highLightedPrice capacity is 3",highLightedPrice)


	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	
	// go create a slice and undelying array in the background. When appending a value, create a new underlying array if the capacity exceeded
	dynamicPrices := []float64{10.1,9.1}
	fmt.Println("dynamicPrices",dynamicPrices)
	updatedDynamicPrices:=append(dynamicPrices,8.0)
	fmt.Println("updatedDynamicPrices",updatedDynamicPrices)
	fmt.Println("dynamicPrices0] memory address",&dynamicPrices[0])
	fmt.Println("dynamicPrices[1] memory address",&dynamicPrices[1])
	fmt.Println("updatedDynamicPrices[0] memory address",&updatedDynamicPrices[0])
	fmt.Println("updatedDynamicPrices[1] memory address",&updatedDynamicPrices[1])
	updatedDynamicPrices=updatedDynamicPrices[:]
	fmt.Println("updatedDynamicPrices[0] memory address",&updatedDynamicPrices[0])
	fmt.Println("updatedDynamicPrices[1] memory address",&updatedDynamicPrices[1])

	updatedDynamicPrices=updatedDynamicPrices[1:]
	fmt.Println("updatedDynamicPrices[0] memory address",&updatedDynamicPrices[0])
	fmt.Println("cap of updatedDynamicPrices",cap(updatedDynamicPrices))


	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	

	testArray := [2]float64{10.1,9.1}
	testArray2 :=testArray 

	fmt.Println("testArray[0] memory address",&testArray[0])
	fmt.Println("testArray[1] memory address",&testArray[1])

	fmt.Println("testArray2[0] memory address",&testArray2[0])
	fmt.Println("testArray2[1] memory address",&testArray2[1])

}

