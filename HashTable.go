package main

import ("fmt")


func main() {

	salary := make(map[string]map[string]int)
	salary["aditya"] = make(map[string]int)
	salary["aditya"]["pn"] = 123142312312
	fmt.Println(salary["aditya"]["pn"])
}