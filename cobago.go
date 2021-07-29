package main
import "fmt"

func main (){
	fmt.Print("Masukkan nomor: ")
	var a int
	fmt.Scanln(&a)

	if a%5==0 && a%3==0 {
		fmt.Println("Hello World")
	}else if a%5==0 {
		fmt.Println("World")
	}else if a%3==0 {
		fmt.Println("Hello")
	}else{
		fmt.Println("Salah Input")
	}
}