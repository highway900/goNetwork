package goNetwork

import "fmt"

// Node interface
// User needs to return a string which represents an ID
//
type INode interface {
    ID() string
    Edges() []IEdge
    Connect(n IEdge)
}


type Node struct {
    Data interface{}
    id string
    edges []IEdge
}


func (n *Node) ID() string {
    return n.id
}


func (n *Node) Edges() []IEdge {
    return n.edges
}


func (n *Node) Connect(e IEdge) {
    n.edges = append(n.edges, e)
}


func (node *Node) OutNodes() []*Node {
    outNodes := make([]*Node, len(node.Edges()))
    for i, edge := range node.Edges() {
        outNodes[i] = node.GetDest(edge.(*Edge))
    }
    return outNodes
}


// This is not directed
// Would need to use parent Graph object to handle a graph type system
func (node *Node) GetDest(edge *Edge) *Node {
    if node == edge.N1() {
        return edge.N0().(*Node)
    } else {
        return edge.N1().(*Node)
    }
}


func (node *Node) Print() {
    fmt.Println(node)
}
