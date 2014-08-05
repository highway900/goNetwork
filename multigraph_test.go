package goNetwork


import (
    "testing"
    "fmt"
)


var seed int64


func SetupTest(s string) {

    fmt.Printf("\n==================\nTesting: %s\n................\n\n", s)

}

func TestMultiGraphCreate(t *testing.T) {

    SetupTest("Connect MultiGraph")

    // create a node
    node1 := &Node{Data: 1}
    node2 := &Node{Data: 2}


    node1.Connect(node2, 1.0)
    node2.Connect(node1, 2.0)
    node1.print()

}


func TestDfsMultiGraph(t *testing.T) {

    SetupTest("DFS")

    // create a node
    node1 := &Node{Data: 1}
    node2 := &Node{Data: 2}
    node3 := &Node{Data: 3}
    //node4 := &Node{Data: 4}
    //node5 := &Node{Data: 5}

    node1.Connect(node2, 1.0)
    node1.Connect(node3, 2.0)
    //node3.Connect(node2, 3.0)
    //node4.Connect(node3, 4.0)
    //node5.Connect(node1, 5.0)
    //node5.Connect(node3, 6.0)


    node := Dfs(node1, 3)
    node.print()

}
