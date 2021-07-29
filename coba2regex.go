package main

import (
    "fmt"
    "regexp"
)

func main() {
    fmt.Println(isEmailValid("test@.gmail.co.id")) 
    fmt.Println(isEmailValid("te.st@.gmail.com"))       
	fmt.Println(isEmailValid("bad-email"))                    
	fmt.Println(isEmailValid("test-email.com"))         
	fmt.Println(isEmailValid("testingmaksimal20kar@.test.id"))
}

func isEmailValid(a string) bool {
    emailRegex := regexp.MustCompile(`^[a-z0-9._-]{1,20}@\.[a-z0-9.-]+(\.id)|(\.co.id)$`)
    return emailRegex.MatchString(a)
	
}
