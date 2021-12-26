package main

import "fmt"

type Node struct{
	value int
	next  *Node
}

func create_node(value int) Node {
	var new_node Node
	new_node.value = value
	new_node.next = nil
	return new_node
}

func add_node_at_head(head *Node, value int) *Node{
	var node Node = create_node(value)
	if head == nil{
		head = &node
	}else{
		var temp *Node = head
		head = &node
		node.next = temp
	}
	return head
}

func add_node_at_end(head *Node, value int) *Node{
	var node Node = create_node(value)
	if head == nil{
		head = &node
	}else{
		temp := head

		for temp.next != nil{
			temp = temp.next		
		}
		temp.next = &node
	}
	return head

}

func traverse_linked_list(head *Node){
	var temp *Node = head

	for temp != nil{
		fmt.Println(temp.value)
		temp = temp.next
	}
}


func main(){

	var head *Node = nil

	fmt.Println("-----------------------------Insert node at the start------------------------------")
	head = add_node_at_head(head, 5)
	head = add_node_at_head(head, 10)
	traverse_linked_list(head)

	fmt.Println("-----------------------------Insert node at the end------------------------------")

	head = add_node_at_end(head, 20)
	head = add_node_at_end(head, 40)
	traverse_linked_list(head)

}
