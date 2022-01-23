package music

import (

	"layeh.com/gopus"
)

type Music struct {
	link	string
	art		string
}

type Queue struct {
	data	int
	next	*Node
}

func (n *Node) appendToTail(d *Music) {
	addnode := &Node{data: d, next: nil}

	for {
		if n.next != nil {
			n = n.next
		} else {
			break
		}
	}

	n.next = addnode
}

func (n *Node) printAll() {
	iter := n
	for iter != nil {
	   fmt.Println(iter.data)
	   iter = iter.next
	}
}

func (n *Node) deleteNode(d int) *Node {
	if n.data == d {
		return n.next
	}

	for n.next != nil {
		if n.next.data == d {
			n.next = n.next.next
			return n
		}
	n = n.next 
	}

	return n
}
