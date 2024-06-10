package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Expenses struct {
	Name     string `json:name`
	Category string `json:category`
	Value    int    `json:value`
}
type CarDetail struct {
	Id      int16  `json:id`
	CarName string `json:carName`
	Modal   string `json:modal`
	Cost    int16  `json:cost`
}

func main() {
	var t0 = time.Now()
	DataArrayPractice()
	ExtractJsonData()
	fmt.Println(time.Since(t0))

}

func DataArrayPractice() {
	var dataArray = make([]Expenses, 0, 10)

	// append 10 data in slice
	for i := 0; i < 10; i++ {
		dataArray = append(dataArray, Expenses{Name: "kushal", Category: "food", Value: 200})
	}

	// transformedArray, err := json.Marshal(dataArray)

	// if err != nil {
	// 	log.Fatal("error while opening json file")

	// }

	// fmt.Printf("%s", transformedArray)
	StringArray := append(make([]string, 0, 10), "kusahlA")
	fmt.Println(dataArray[0].Name)
	fmt.Println(StringArray)
}

func ExtractJsonData() {

	var CarObj []CarDetail
	// reading the json file
	jsonfile, err := os.ReadFile("car_details.json")
	if err != nil {
		log.Fatal("error while opening json file")

	}

	// jsonData, err := json.Marshal(jsonFile)

	// if err != nil {
	// 	log.Fatal("error while opening json file")
	// }

	// println json file give bytecode of josn and printing as string give actual code
	// fmt.Println(jsonfile)
	// fmt.Printf("%s\n", jsonfile)

	// umnarshal the json file and convert it into the structure format
	err = json.Unmarshal(jsonfile, &CarObj)
	if err != nil {
		log.Fatal("error while unmarshal", err)

	}
	fmt.Println(CarObj)

	fmt.Println(CarObj[0].CarName)
	// iterate over array of object
	// for index, value := range CarObj {
	// 	fmt.Println(value.CarName)
	// 	fmt.Println(index)
	// }
	// for _, row := range CarObj {
	// 	counter[row]++
	// }

	GroupVehicle := make(map[string][]CarDetail)
	fmt.Println(GroupVehicle)
	// iterating through whole data of cars in struct format and group by their car name
	for _, p := range CarObj {
		GroupVehicle[p.CarName] = append(GroupVehicle[p.CarName], p)

	}
	fmt.Println(GroupVehicle)

}
