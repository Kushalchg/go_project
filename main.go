package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	advancefunction "github.com/kushalchg/go_project/advanceFunction"
	structure "github.com/kushalchg/go_project/struct"
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
	// DataArrayPractice()
	// data, err := ExtractJsonData()
	// if err != nil {

	// 	log.Printf("error on data extract %v", err)

	structure.StructConcept()
	fmt.Println(advancefunction.Aggrigate(2, 3, 4, advancefunction.Add))

	fmt.Println(advancefunction.Aggrigate(2, 3, 4, advancefunction.Mul))
	fmt.Println(advancefunction.Aggrigate(2, 3, 4, advancefunction.Avg))
	// }
	// fmt.Printf("data from main fumction %v", data)
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

func ExtractJsonData() ([]CarDetail, error) {

	var CarObj []CarDetail
	// reading the json file
	jsonfile, err := os.ReadFile("car_details.json")
	if err != nil {
		log.Fatal("error while opening json file")

	}
	// it print the json file in string format i.e it convert byte code into stirng and print in json format

	// fmt.Printf("first object of json file %v \n ", string(jsonfile))
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
	// fmt.Println("cars array=", CarObj)

	// iterate over array of object
	// for index, value := range CarObj {
	// 	fmt.Println(value.CarName)
	// 	fmt.Println(index)
	// }
	// for _, row := range CarObj {
	// 	counter[row]++
	// }

	// GroupVehicle is map contain the array of CarObj for unique cars by car name
	GroupVehicle := make(map[string][]CarDetail)
	// iterating through whole data of cars in struct format and group by their car name
	for _, p := range CarObj {
		GroupVehicle[p.CarName] = append(GroupVehicle[p.CarName], p)
	}
	fmt.Printf("Group vehicle by their barad name:= %v \n", GroupVehicle)

	//count the number of cars with same brand name

	CountVehicle := make(map[string]int)

	// method 1(it require extra maps)
	// for _, p := range CarObj {
	// 	CountVehicle[p.CarName] = len(append(GroupVehicle[p.CarName], p))
	// }

	// method 2 (more efficient method and short one)
	for _, cars := range CarObj {
		CountVehicle[cars.CarName]++
	}

	fmt.Printf("No of cars by brand name:=%v\n", CountVehicle)
	return CarObj, err

}
