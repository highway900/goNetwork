package goNetwork


import (
    "testing"
    "fmt"
)


var seed int64


func SetupTest(s string, t ...int) {

    fmt.Printf("\n================\nTesting: %s\n................\n\n", s)

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

    fmt.Println(len(node1.Edges))
    node1.Print(true, true)

}

func printKV(r VisitMap) {
    for k, v := range r {
        fmt.Println(k.Data, v)
    }
}
func TestDfsMultiGraph(t *testing.T) {

    SetupTest("DFS")

    // create a node
    node1 := &Node{Data: 1}
    node2 := &Node{Data: 2}
    node3 := &Node{Data: 3}
    node4 := &Node{Data: 4}
    node5 := &Node{Data: 5}
    node6 := &Node{Data: 6}

    graph := &Graph{
        Visited: make(map[*Node]bool),
    }

    graph.Connect(node1, node2, 1.0, "k")
    graph.Connect(node1, node3, 1.0, "k")

    graph.Connect(node2, node4, 1.0, "k")
    graph.Connect(node3, node4, 1.0, "k")
    graph.Connect(node5, node4, 1.0, "k")
    graph.Connect(node1, node6, 1.0, "k")
    graph.Connect(node6, node5, 1.0, "k")


    n := graph.Dfs(node6)
    m := graph.Dfs_(node5)
    if n != nil {
        fmt.Println("---Success---")
        fmt.Println("---Looped DFS---")
        printKV(n)
        fmt.Println("---Recursive DFS---")
        printKV(m)
    } else {
        fmt.Println("Failed")
    }

}
