package goNetwork

import (
    "fmt"
)

type SizedArray struct {
    items []interface{}
    current int // Array is filled to here
}

func MakeSizedArray() SizedArray {
    return SizedArray {
        make([]interface{}, 1),
        0,
    }
}

// returns the last item in the array
func (sa *SizedArray) Last() interface{} {
    return sa.items[sa.current]
}

// pushes the item onto the array at the back
func (sa *SizedArray) PushBack(item interface{}) {
    sa.items[sa.current] = item
    sa.current += 1
}


type VisitMap map[*Node]bool


type Graph struct {
    Edges []*Edge
    Nodes []*Node
    Visited VisitMap
    VisitOrder SizedArray
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


func (g *Graph) AddNode(data int) *Node {
    node := &Node{Data: data}
    g.Nodes = append(g.Nodes, node)
    return node
}

// Recursive version
func (g *Graph) Dfs_(start *Node) SizedArray {
    g.Visited[start] = true
    g.VisitOrder.PushBack(start)
    for _, node := range start.OutNodes() {
        if !g.Visited[node] {
            g.Dfs_(node)
        }
    }
    return g.VisitOrder
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
