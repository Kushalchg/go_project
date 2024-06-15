package structure

import (
	"errors"
	"fmt"
)

var Struct = "kushal"

type Info struct {
	name   string
	age    int32
	gender string
}

func CreateNewStruct(age int32) (Info, error) {
	if age < 0 {
		return Info{}, errors.New("age cant be in negative")
	}
	s := Info{age: age}
	s.gender = "new struct gender"
	s.name = "new struct name"
	return s, nil
}

func ModifyStructWP(name, gender string, age int32) (*Info, error) {

	if age < 18 {
		errror := errors.New("you are unable to vote")
		return &Info{}, errror

	}
	s := Info{name, age, gender}
	// the name is updated in s structure and it doesnot create new structure and save memory consumption
	s.name = "updated Kushal chapagain"
	return &s, nil

}

func StructConcept() {

	newData, err := ModifyStructWP("kushal chapagain", "male", 19)

	fmt.Printf("the value from pointer fo the struct is %v\n", newData)
	if err != nil {
		fmt.Println("error occured while creating new struct through pointer", err)
	}

	data, err := CreateNewStruct(-2)
	if err != nil {
		fmt.Println("error occured while creating new struct", err)
	}
	fmt.Printf("the value  of new struct is here %v", data)
	fmt.Printf("empty structure default= %v\n", Info{})

	fmt.Printf("the structure with single value:= %v\n", Info{"kushal", 23, "male"})
	fmt.Printf("the structure with single value:= %v\n", Info{gender: "kushal", age: 23, name: "male"})
}
