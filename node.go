package goNetwork

import (
    "fmt"
)

// Simple edge that holds a int value
type Edge struct {
    Value int
    node1 *Node
    node2 *Node
}

// Simple Node
type Node struct {
   Value int
   edges []*Edge
}

func (this *Node) Print() {

    fmt.Println("Node Value:", this.Value)

    for _, edge := range this.edges {

        fmt.Println("Edge Value:", edge.Value)

    }

}

func (this *Node) AddEdge(edge *Edge) {

   this.edges = append(this.edges, edge)

}

func (this *Edge) AddNodes(node1 *Node, node2 *Node) {

    this.node1 = node1
    this.node2 = node2

}

func (this *Edge) PrintValue() {

    fmt.Println(this.Value)

}
