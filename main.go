package main

import (
	"fmt"
	"testing/houserobber"
	"testing/treeStracture"
)

func main() {

	var t treeStracture.Tree
	t.InsertData(10)
	t.InsertData(20)
	t.InsertData(30)
	t.InsertData(40)
	t.InsertData(50)
	t.InsertData(50)
	fmt.Printf("PreOrderTraversal: ")
	treeStracture.PreOrderTraversal(t.Root)
	fmt.Printf("\nInOrderTraversal: ")
	treeStracture.InOrderTraversal(t.Root)
	fmt.Printf("\nPostOrderTraversal: ")
	treeStracture.PostOrderTraversal(t.Root)
	fmt.Println()
	input1 := []int{2,3,2}
	fmt.Printf("Input Parameter : %v \nOutput : %v \n",input1,houserobber.MaxedSumAcquired(input1))
	input2 := []int{1,2,3,1}
	fmt.Printf("Input Parameter : %v \nOutput : %v \n",input2,houserobber.MaxedSumAcquired(input2))
	input3 := []int{0}
	fmt.Printf("Input Parameter : %v \nOutput : %v \n",input3,houserobber.MaxedSumAcquired(input3))
}
