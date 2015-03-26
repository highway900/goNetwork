package main

import (
    "github.com/highway900/goNetwork"
)

func main() {

    node1 := &goNetwork.Node{Data:1}
    node2 := &goNetwork.Node{Data:2}


    for i := 0; i < 4; i++  {
        node1.Connect(node2, float64(i))
    }


    node1.print()
    node2.print()

    for _, edge := range node1.Edges {
        edge.N0().Print()
        edge.N1().Print()
    }

}
