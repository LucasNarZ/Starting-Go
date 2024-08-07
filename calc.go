package main

import (
    "fmt"
)

type Operation func(float32, float32) (float32, error)

func main(){
	operations := map[string] Operation{
		"add":func (num1 float32, num2 float32) (float32, error){
			return num1 + num2, nil
		},
		"sub":func (num1 float32, num2 float32) (float32, error){
			return num1 - num2, nil
		},
		"mul":func (num1 float32, num2 float32) (float32, error){
			return num1 * num2, nil
		},
		"div":(func (num1 float32, num2 float32) (float32, error){
			if num2 == 0 {
                return 0, fmt.Errorf("division by zero")
            }
			return num1 / num2, nil
		}),
	}

	var num1 float32
	var num2 float32
	var operation string
	for {
		fmt.Print("Type the first number: ")
		fmt.Scanln(&num1)
		fmt.Print("Type the second number: ")
		fmt.Scanln(&num2)
		fmt.Print("Type the operation(add ,sub, mul, div): ")
		fmt.Scanln(&operation)
		opFunc, exists := operations[operation];
		if !exists {
			fmt.Print("Please enter a valid operation!\n")
			continue
		}

		result, err := opFunc(num1, num2);
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("The result is %.1f\n", result);
	}
}