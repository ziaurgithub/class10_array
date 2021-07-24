package main
import "fmt"
func main(){
/*
var chr[3]rune

chr[0]='a'
chr[1]='b'
chr[2]='c'

fmt.Println(chr)
*/
                          //short hand way

chr := [...]rune{'A', 'B', 'C', 'D'}
fmt.Println(chr)

}