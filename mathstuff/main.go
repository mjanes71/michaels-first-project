package main

func addOrSubtract(operation string, num1, num2 int) int {
	var result int
	if operation == "add" {
		result = num1 + num2
	} else {
		result = num1 - num2
	}
	return result
}

//TODO - can you explain what this function does? Do you see any potential problems
//with this function? How could we fix them?
