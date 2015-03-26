package goNetwork

import (
	"fmt"
)

// Node interface
// User needs to return a string which represents an ID
//
type INode interface {
    ID() string
    Edges() []*IEdge
}


type Node struct {
    Data interface{}
    Edges []*IEdge
}


func (node *Node) Print(edges bool, adjacent bool) {
    fmt.Println("Data:", node.Data)
    for i, edge := range node.Edges {
        if adjacent {
            fmt.Printf("Adjacent ---> %d\n", node.GetDest(edge).Data)
        }
        if edges {
            fmt.Printf("Edge %d weight: %f\n", i, edge.Weight())
        }
    }
    fmt.Printf("\n")
}


func (node *Node) OutNodes() []*Node {
    outNodes := make([]*Node, len(node.Edges))
    for i, edge := range node.Edges {
        outNodes[i] = node.GetDest(edge)
    }
    return outNodes
}


// This is not directed
// Would need to use parent Graph object to handle a graph type system
func (node *Node) GetDest(edge *Edge) *Node {
    if node == edge.N1() {
        return edge.N0()
    } else {
        return edge.N1()
    }
}
