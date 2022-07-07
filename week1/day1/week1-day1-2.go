package main

import "fmt"

type Node struct{
	val string
	left *Node
	right *Node
}
func (Tree *Node)Print_Inorder() {
	if Tree == nil{
		return
	}
	Tree.left.Print_Inorder()
	fmt.Println(Tree.val)
	Tree.right.Print_Inorder()
}
func (Tree *Node)Print_Preorder() {
	if Tree == nil{
		return
	}
	fmt.Println(Tree.val)
	Tree.left.Print_Preorder()
	Tree.right.Print_Preorder()
}
func (Tree *Node)Print_Postorder() {
	if Tree == nil{
		return
	}
	Tree.left.Print_Postorder()
	Tree.right.Print_Postorder()
	fmt.Println(Tree.val)
}
func main() {
	var Tree Node = Node{"+",&Node{"a",nil,nil},&Node{"-",&Node{"b",nil,nil},&Node{"c",nil,nil}}}
	Tree.Print_Postorder()
}
