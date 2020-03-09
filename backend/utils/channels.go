package utils

import(
	"fmt"
	"math/rand"
)

func GenerateValue(values chan int){
	value := rand.Intn(10)
	fmt.Println("Generated value = {}", value)

	values <- value
}