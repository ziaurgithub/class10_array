package main
import "fmt"

func main(){


x := make(map[rune]float32)
x['a'] = 5.7
x['b'] = 5.6
x['c'] = 5.5
fmt.Println(x)
}