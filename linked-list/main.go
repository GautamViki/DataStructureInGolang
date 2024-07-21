package main

import (
	"fmt"

	"main.go/helper"
)

func main() {
	fmt.Println("hello linked list")
	head := helper.Head{}
	head.InsertValueInList()
	head.GetAllNode()
	head.DeleteFirstNode()
	head.DeleteLastNode()
}
