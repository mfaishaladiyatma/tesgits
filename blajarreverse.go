package main
import "fmt"

func main(){
	fmt.Print("Masukkan Kata : ")
	var str string
	fmt.Scanln(&str)

	var revstr string
	var length=len(str)

	for a:=length - 1; a>=0; a-- {
		revstr = revstr + string(str[a])
	}
	fmt.Printf("\nKata yang dimasukkan adalah %s", str)
	fmt.Printf("\nSetelah dibalik menjadi %s", revstr)
}