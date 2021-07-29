package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Print("Masukkan Kata : ")
	var str string
	fmt.Scanln(&str)
	
	var revstr string 
	var length=len(str)

	for a:=length - 1; a>=0; a-- {
		revstr = revstr + string(str[a])
	}
	fmt.Printf("Kata yang dimasukkan adalah %s", str)
	fmt.Printf("\nSetelah dibalik menjadi %s", revstr)

	if strings.ToLower(str) == strings.ToLower(revstr) {
		fmt.Println("\nKata ini adalah PALINDROME")
	} else {
		fmt.Println("\nKata ini BUKAN PALINDROME")
	}
}
