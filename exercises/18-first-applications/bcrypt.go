package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := `mystring`

	fmt.Println("Password: ", password)
	sb, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error while generating password: ", err)
		os.Exit(1)
	}
	fmt.Printf("%v", sb)

	loginPassword := "mystring"
	err = bcrypt.CompareHashAndPassword(sb, []byte(loginPassword))
	if err != nil {
		fmt.Println("Unable to compare:", err)
		os.Exit(1)
	}

	fmt.Println("good")
}
