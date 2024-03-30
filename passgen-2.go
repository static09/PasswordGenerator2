package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.design/x/clipboard"
)

// create bytes to use for password string
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// convert password from randomly generated bytes
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.RawStdEncoding.EncodeToString(b), err
}

func main() {
	//Define paramaters
	username := ""
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	passLength := 16
	fmt.Print("Password Length: ")
	fmt.Scanln(&passLength)
	//Generate password
	token, err := GenerateRandomString(passLength)
	if err != nil {
		fmt.Println("Error")
	}
	//display paramaters
	fmt.Println("User: " + username)
	fmt.Println("Pass: " + token)
	//prepare combined paramaters to be sent to clipboard
	var creds string = username + "\r" + token
	//ask user y/n to copy credentials
	var copyPass string = "y"
	fmt.Print("Copy pass? (y/n): ")
	fmt.Scanln(&copyPass)
	if copyPass == "y" {
		//initilize clipboard package
		err := clipboard.Init()
		if err != nil {
			panic(err)
		}
		//write credentials to clipboard
		clipboard.Write(clipboard.FmtText, []byte(creds))
		//verify clipboard has credentials
		board := string(clipboard.Read(clipboard.FmtText))
		if board == creds {
			fmt.Print("Successfully copied to clipboard!\n")

		}
		if board != creds {
			fmt.Print("Error copying to clipboard.")
		}
		input := ""
		fmt.Print("Press any key to exit...")
		fmt.Scanln(&input) //exit on any key
		return
	}
	if copyPass == "n" {
		//threaten user and make them feel bad for not copying the password
		fmt.Print("Your funeral.\nExiting...")
		return
	}
}