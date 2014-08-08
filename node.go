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
    Visited bool
    Edges []*Edge
}


// Should be a function of a Graph struct
func dfs(node *Node, data int, visited map[*Node]bool) *Node {
    node.print(false)
    // Check the node has been visited
    if node.Data == data {
        return node
    } else {
        visited[node] = true
        for _, edge := range node.Edges {
            node_dest := node.GetDest(edge)
            if !visited[node_dest] {
                return dfs(node_dest, data, visited)
            }
        }
    }
    return nil
}


func Dfs(node *Node, data int) *Node {
    visited := make(map[*Node]bool)
    return dfs(node, data, visited)
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


func (node *Node) print(edge bool) {
    fmt.Println("Data:", node.Data)
    if edge {
        for i, edge := range node.Edges {
            fmt.Printf("Edge %d weight: %f\n", i, edge.Weight)
        }
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
