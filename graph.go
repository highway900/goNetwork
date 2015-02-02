package goNetwork

import (
    "fmt"
)

type SizedArray struct {
    items []interface{}
    current int // Array is filled to here
}

func MakeSizedArray(size int32) SizedArray {
    return SizedArray {
        make([]interface{}, size),
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


type visitMap map[*Node]bool


// The goal of this object is to keep track of things
// while searching etc happens
type Visitor struct {
    Visited visitMap
    VisitOrder SizedArray
}


func MakeVisitor(size int) *Visitor {
     return &Visitor{
        Visited: make(visitMap),
        VisitOrder: MakeSizedArray(int32(size)),
     }
}


type Graph struct {
    Edges []*Edge
    Nodes []*Node
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


func (g *Graph) Dfs(start *Node) SizedArray {
    visitor := MakeVisitor(len(g.Nodes))
    return g.dfs(start, visitor)
}


// Recursive version
func (g *Graph) dfs(start *Node, visitor *Visitor) SizedArray {
    visitor.Visited[start] = true
    visitor.VisitOrder.PushBack(start)
    for _, node := range start.OutNodes() {
        if !visitor.Visited[node] {
            g.dfs(node, visitor)
        }
    }
    return visitor.VisitOrder
}


func (g *Graph) DfsPath(start *Node, goal *Node) SizedArray {
    visitor := MakeVisitor(len(g.Nodes)) // TODO: This is likely vBad for large networks. Consider dynamic resizing...
    return g.dfsPath(start, goal, visitor)
}


func (g *Graph) dfsPath(start *Node, goal *Node, visitor *Visitor) SizedArray {
    visitor.VisitOrder.PushBack(start)
    if start == goal {
        return visitor.VisitOrder
    }
    for _, node := range start.OutNodes() {
        if !visitor.Visited[node] {
            g.dfsPath(start, goal, visitor)
        }
    }
    return visitor.VisitOrder
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
