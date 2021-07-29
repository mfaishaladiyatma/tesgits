package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(validasiemail("test@.gmail.co.id"))
	fmt.Println(validasiemail("te.st@.gmail.com"))
	fmt.Println(validasiemail("test-email.com"))
	fmt.Println(validasiemail("testingmaksimal20karakterrr@.test.id"))
}

func validasiemail(a string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._-]{1,20}@\.[a-z0-9.-]+(\.id)|(\.co.id)$`)
	return emailRegex.MatchString(a)

}
