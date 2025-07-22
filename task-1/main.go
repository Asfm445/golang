package main

import (
	"fmt"
)

type subject struct {
	sub_name string
	grade    float64
}

func calculate_grade(subjects []subject, sub_num int) float64 {
	var total float64
	for _, val := range subjects {
		total = total + val.grade
	}
	return total / float64(sub_num)
}

func main() {
	fmt.Print("Enter your name: ")
	var name string
	var num_sub int
	fmt.Scanln(&name)
	fmt.Print("Enter number of subjects: ")
	fmt.Scanln(&num_sub)
	subjects := make([]subject, num_sub)
	for i := 0; i < num_sub; i++ {
		fmt.Print("Enter subject name: ")
		var sub_name string
		var grade float64
		fmt.Scan(&sub_name)
		fmt.Print("Enter subject grade: ")
		fmt.Scan(&grade)
		for grade < 0 || grade > 4 {
			fmt.Println("invalid grade enter again!")
			fmt.Scan(&grade)
		}
		subjects[i] = subject{sub_name: sub_name, grade: grade}
	}
	grade := calculate_grade(subjects, num_sub)
	fmt.Println("\nsubjects     grade ")
	fmt.Println("____________________________________")
	for _, val := range subjects {
		fmt.Printf("%-15v %0.1f \n", val.sub_name, val.grade)
	}
	fmt.Printf("%-15v %0.1f ", "total", grade)
}
