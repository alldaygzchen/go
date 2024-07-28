package main

import "fmt"

func exercise3(){

	prices:=[]int{1,2,3,4,5}
	discoutPrices:=  []int{6,7,8}

	prices = append(prices,discoutPrices...)
	fmt.Println("prices",prices)


}