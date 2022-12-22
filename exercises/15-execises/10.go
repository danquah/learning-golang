package main

import "fmt"

func getSecretsHandler() func() string {
	secret := "shhh"

	return func() string {
		return secret
	}
}

func main() {
	/*
		Closure is when we have “enclosed” the scope of a variable in some code block.
		For this hands-on exercise, create a func which “encloses” the scope of a variable:
	*/

	secretsHandler := getSecretsHandler()

	fmt.Printf("The secret is %v", secretsHandler())
}
