package goNetwork

import (
    "fmt"
)



type Graph struct {
    Edges []*Edge
    Nodes []*Node
    Visited VisitMap
}


type Edge struct {
    Node_src *Node
    Node_dest *Node
    Weight float64
    Key string
}


type Node struct {
    Data int
    Edges []*Edge
}


// Loop Version
func (g *Graph) Dfs(start *Node) VisitMap {
    stack := new(Stack)
    stack.Push(start)
    for stack.Len() > 0 {
        node := stack.Pop().(*Node)
        if !g.Visited[node] {
            g.Visited[node] = true
            for _, outNode := range start.OutNodes() {
                if !g.Visited[outNode] {
                    stack.Push(outNode)
                }
            }
        }
    }
    return g.Visited
}


// Recursive version
func (g *Graph) Dfs_(start *Node) VisitMap {
    g.Visited[start] = true
    for _, node := range start.OutNodes() {
        if !g.Visited[node] {
            g.Dfs_(node)
        }
    }
    return g.Visited
}


func _dfs_path(graph *Graph, start *Node, goal *Node, path VisitMap) VisitMap {
    return nil
}


func (graph *Graph) Connect(node_src *Node, node_dest *Node, weight float64, key string) {

    edge := &Edge{
        node_src,
        node_dest,
        weight,
        key,
    }

    node_src.Edges = append(node_src.Edges, edge)
    node_dest.Edges = append(node_dest.Edges, edge)

    graph.Edges = append(graph.Edges, edge)
    graph.Nodes = append(graph.Nodes, node_src, node_dest)

}


func (node *Node) Print(edges bool, adjacent bool) {
    fmt.Println("Data:", node.Data)
    for i, edge := range node.Edges {
        if adjacent {
            fmt.Printf("Adjacent ---> %d\n", node.GetDest(edge).Data)
        }
        if edges {
            fmt.Printf("Edge %d weight: %f\n", i, edge.Weight)
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
    if node == edge.Node_dest {
        return edge.Node_src
    } else {
        return edge.Node_dest
    }
}


func (edge *Edge) Print() {

    fmt.Println(edge.Weight)

}
