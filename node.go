package goNetwork

import (
    "fmt"
)


type Edge struct {
    Node_src *Node
    Node_dest *Node
    Weight float64
}


type Node struct {
    Data int
    Edges []*Edge
}


func Dfs(node *Node, data int) *Node {
    visited := make(map[*Node]bool)
    for _, edge := range node.Edges {
        if data == node.Data {
            return node
        }
        _, found := visited[node]
        if found {
            continue
        }
        visited[node] = true
        node_dest := node.GetDest(edge)
        return Dfs(node_dest, data)
    }
    return node
}


func (node_src *Node) Connect(node_dest *Node, weight float64) {

    edge := &Edge{
        node_src,
        node_dest,
        weight,
    }

    // Allocating before might be a better way
    // of handling this. A good test for profiling.
    node_src.Edges = append(node_src.Edges, edge)
    node_dest.Edges = append(node_dest.Edges, edge)

}


func (node *Node) Print() {
    fmt.Println("Data:", node.Data)
    for i, edge := range node.Edges {
        fmt.Printf("Edge Weight %d: %f\n", i, edge.Weight)
    }
}


func MakeEdge(src *Node, dest *Node, weight float64) {

    edge := &Edge{}
    edge.Node_src = src
    edge.Node_dest = dest
    edge.Weight = weight

}


// This is not directed
// Would need to use parent Graph object to handle a graph type system
func (node *Node) GetDest(edge *Edge) *Node {
    if node == edge.Node_dest {
        return edge.Node_src
    } else {
        return edge.Node_dest
    }
}


func (edge *Edge) Print() {

    fmt.Println(edge.Weight)

}
