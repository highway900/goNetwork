package goNetwork


import (
    "testing"
    "fmt"
)


var seed int64


func SetupTest(s string) *Graph {

    fmt.Printf("\n================\nTesting: %s\n................\n\n", s)

    graph := new(Graph)

    // create a node
    node1 := graph.AddNode(1)
    node2 := graph.AddNode(2)
    node3 := graph.AddNode(3)
    node4 := graph.AddNode(4)
    node5 := graph.AddNode(5)
    node6 := graph.AddNode(6)
    node7 := graph.AddNode(7)

    graph.Connect(node1, node2, 1.0, "k")
    graph.Connect(node1, node3, 1.0, "k")

    graph.Connect(node2, node4, 1.0, "k")
    graph.Connect(node3, node4, 1.0, "k")
    graph.Connect(node5, node4, 1.0, "k")
    graph.Connect(node1, node6, 1.0, "k")
    graph.Connect(node6, node5, 1.0, "k")
    graph.Connect(node6, node7, 1.0, "k")

    return graph

}

func TestMultiGraphCreate(t *testing.T) {

    SetupTest("Connect MultiGraph")

    // create a node
    node1 := &Node{Data: 1}
    node2 := &Node{Data: 2}
    node3 := &Node{Data: 3}

    graph := new(Graph)
    graph.Connect(node1, node2, 1.0, "k")
    graph.Connect(node1, node3, 1.0, "k")

    node1.Print(true, true)

}


func printKV(r []interface{}) {
    for _, v := range r {
        fmt.Println(v.(*Node).Data)
    }
}


func TestGraphOutNodes(t *testing.T) {
    graph := SetupTest("OutNodes")
    for _, node := range graph.Nodes {
        for _, n := range node.OutNodes() {
            fmt.Printf("%d --> %d\n", node.Data, n.Data)
        }
    }
}


func TestDfsMultiGraph(t *testing.T) {
    graph := SetupTest("DFS Recursive")

    m := graph.Dfs(graph.Nodes[5])
    fmt.Println("visit order items", m.items)
    // Compare the m.items to graph.Nodes both should contain the same items.
    should_be_empty := m.items
    fmt.Println(should_be_empty)
    for x, i := range m.items {
        for _, j := range graph.Nodes {
            if i == j {
                should_be_empty = should_be_empty[1:]
            }
        }
    }
    fmt.Println(should_be_empty)
}
