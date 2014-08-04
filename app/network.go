package main

import (
    "fmt"
    "github.com/highway900/goNetwork"
)

func main() {

    fmt.Println("Generating Network")

    nNodes := 4
    nEdges := 5
    nodes := make([]*goNetwork.Node, nNodes)
    edges := make([]*goNetwork.Edge, nEdges)

    for _, edge := range edges {
        newEdge := new(goNetwork.Edge)
        newEdge.Value = i
    }

    c := 0
    for _, node := range nodes {
        node = new(goNetwork.Node)
        node.Value = c
        node.AddEdge(newEdge)
        c++
    }


}
