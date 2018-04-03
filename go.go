package main

import "fmt"

type Position struct {
	row int
	col int
	name string
}

func (posstn Position) isvalid() bool {
	if posstn.row > 8 || posstn.row < 0 || posstn.col > 8 || posstn.col < 0 {
		return false
	}
	return true
}

func Possmov(pos Position) {
	isval := pos.isvalid()
	if isval == true {
		//name := Position{}
		fmt.Println("Something")
		fmt.Println(pos.name, pos.col, pos.row)
	}
}
func main() {
	Possmov(Position{1, 7,"shik"}) //Something
}
