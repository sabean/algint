package main 

import "fmt"

func result(contain []string, length int){
	count := 0
	for i:=0; i < length; i++{
		if contain[i] == " "{
			count++
		}
	}
	newLength := length + count*2
	j := length-1
	for k := newLength-1; k >= 0; k-- {
		if contain[j] == " "{
			contain[k] = "0"
			contain[k-1] = "2"
			contain[k-2] = "%"
			k = k -2
		} else{
			contain[k] = contain[j]
		}
		j--
	}
	fmt.Println(contain)
}
func main() {
	contain := []string{"M", "y", " ", "n", "a", "m", "e", " ", "J", "o", "n", "e", "s", " ", " ", " ", " "}
	var length int = 13
	result(contain, length)
}