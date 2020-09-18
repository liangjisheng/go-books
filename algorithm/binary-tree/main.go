package main

import "fmt"

func main() {
	// buildTreeDemoByPreAndIn()
	// buildTreeDemoByInAndPost()

	preOrderData := []int{3, 9, 20, 15, 7}
	inOrderData := []int{9, 3, 15, 20, 7}
	// postOrderData := []int{9, 15, 7, 20, 3}

	root := buildTreeByPreAndIn(preOrderData, inOrderData)

	preOrder := preOrderIteration(root)
	fmt.Print("pre order: ")
	fmt.Println(preOrder)

	// preOrder = preOrderRecursion(root)
	// fmt.Print("pre order: ")
	// fmt.Println(preOrder)

	inOrder := inOrderIteration(root)
	fmt.Print("in order: ")
	fmt.Println(inOrder)

	// inOrder = inOrderRecursion(root)
	// fmt.Print("in order: ")
	// fmt.Println(inOrder)

	postOrderList := postOrderIteration(root)
	postOrder := listToArray(postOrderList)
	fmt.Print("post order: ")
	fmt.Println(postOrder)

	// postOrder = postOrderRecursion(root)
	// fmt.Print("post order: ")
	// fmt.Println(postOrder)

	// fmt.Println("max depth:", maxDepth(root))

	// resList := printFromTopToBottom(root)
	// resArray := listToArray(resList)
	// fmt.Print("FromTopToBottom: ")
	// fmt.Println(resArray)

	// fmt.Println(judge(postOrder, 0, len(postOrder)-1))

	// reverseTree(root)
	// preOrder = preOrderIteration(root)
	// fmt.Print("reverse pre order: ")
	// fmt.Println(preOrder)
}
