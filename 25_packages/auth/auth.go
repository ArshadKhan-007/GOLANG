package auth

import "fmt"

// Capital character is a convention to use this function outside of the package.
// If we don't give first character capital then we can't use it outside of the package.

func LoginWithCredentials(username string, password string) {
	fmt.Println("Logging you in", username, password)
}
