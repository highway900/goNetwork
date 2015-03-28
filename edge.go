package goNetwork

import (
	"fmt"
)


type IEdge interface {
    ID() string
    Weight() float64
    N0() INode
    N1() INode
}


type Edge struct {
    n0 INode
    n1 INode
    weight float64
    id string
    Key string
}


func (e *Edge) ID() string {
	return e.id
}


func (e *Edge) Weight() float64 {
    return e.weight
}


func (e *Edge) N0() INode {
    return e.n0
}


func (e *Edge) N1() INode {
    return e.n1
}


func (edge *Edge) Print() {
    fmt.Println(edge.Weight)
}
