package goNetwork


import (
    "testing"
    "fmt"
)


var seed int64


func SetupTest(s string) *Graph {

    fmt.Printf("\n================\nTesting: %s\n................\n\n", s)

    graph := &Graph{
        Visited: make(map[*Node]bool),
        VisitOrder: MakeSizedArray(),
    }

    // create a node
    node1 := graph.AddNode(1)
    node2 := graph.AddNode(2)
    node3 := graph.AddNode(3)
    node4 := graph.AddNode(4)
    node5 := graph.AddNode(5)
    node6 := graph.AddNode(6)

    graph.Connect(node1, node2, 1.0, "k")
    graph.Connect(node1, node3, 1.0, "k")

    graph.Connect(node2, node4, 1.0, "k")
    graph.Connect(node3, node4, 1.0, "k")
    graph.Connect(node5, node4, 1.0, "k")
    graph.Connect(node1, node6, 1.0, "k")
    graph.Connect(node6, node5, 1.0, "k")

    return graph

}

func TestMultiGraphCreate(t *testing.T) {

    SetupTest("Connect MultiGraph")

    // create a node
    node1 := &Node{Data: 1}
    node2 := &Node{Data: 2}
    node3 := &Node{Data: 3}

    graph := &Graph{
        Visited: make(map[*Node]bool),
    }

    graph.Connect(node1, node2, 1.0, "k")
    graph.Connect(node1, node3, 1.0, "k")

    node1.Print(true, true)

}


func printKV(r []interface{}) {
    for _, v := range r {
        fmt.Println(v)
    }
}


func TestGraphOutNodes(t *testing.T) {
    graph := SetupTest("OutNodes")
    for _, node := range graph.Nodes {
        node.Print(false, false)
        for _, n := range node.OutNodes() {
            fmt.Println("OutNodes --> ", n)
        }
    }
}


func TestDfsMultiGraph(t *testing.T) {

    graph := SetupTest("DFS")

    n := graph.Dfs(graph.Nodes[0])
    graph.VisitOrder.items = make([]interface{}, len(graph.Nodes))
    m := graph.Dfs_(graph.Nodes[5])
    fmt.Println(m.items[5])
    graph.Nodes[0].Print(false, false)
    if n != nil {
        fmt.Println("---Success---")
        fmt.Println("---Looped DFS---")
        fmt.Println("---Recursive DFS---")
        printKV(m.items)
    } else {
        fmt.Println("Failed")
    }

}
