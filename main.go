package main
import "fmt"

func main(){


var alphabet[3]rune

/*
alphabet[0]='a'
alphabet[1]='b'
alphabet[2]='c'

x := alphabet[0:2]

fmt.Println(x) 
*/


/*
x := make([]rune, 0)
fmt.Println(x)
*/




var chr []rune
chr=append(chr, 'A', 'B', 'Y', 'Z')
fmt.Println(chr)
fmt.Println (chr, len(chr))
fmt.Printf("%T\n", chr)
fmt.Printf("%T",alphabet )


}