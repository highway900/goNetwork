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
    for _, edge := range node.Edges {
        _, found := visited[node]
        if found {
            fmt.Println("Node visited")
            continue
        }
        if data == node.Data {
            fmt.Println("Node found!")
            return node
        }
        visited[node] = true
        node_dest := node.GetDest(edge)
        node_dest.print()
        return dfs(node_dest, data, visited)
    }
    fmt.Println("Node Not found!")
    return node
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


func (node *Node) print() {
    fmt.Println("Data:", node.Data)
    for i, edge := range node.Edges {
        fmt.Printf("Edge %d weight: %f\n", i, edge.Weight)
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
