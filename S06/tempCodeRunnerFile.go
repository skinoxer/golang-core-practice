package main

import (
	"fmt"
)

func main() {
	var password string

	fmt.Print("please enter your password: ")
	fmt.Scanln(&password)

	if len(password) < 8 {
		fmt.Println("Error: Your password must be at least 8 characters.")
		return
	}

	hasLetter := false
	hasNumber := false

	for i := 0; i < len(password); i++ {
		char := password[i]

		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			hasLetter = true
		}

		if char >= '0' && char <= '9' {
			hasNumber = true
		}
	}

	if hasLetter && hasNumber {
		fmt.Println("✅ Success: You can login successfully!")
	} else if !hasLetter {
		fmt.Println("❌ Error: Your password must contain at least one letter (A-Z or a-z).")
	} else if !hasNumber {
		fmt.Println("❌ Error: Your password must contain at least one number (0-9).")
	}
}
