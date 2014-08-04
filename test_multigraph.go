package network


import (
    "testing"
    "math/rand"
    "time"
    "fmt"
)


var seed int64


func Setup() {

    seed = time.Now().Unix() * rand.Int63()
    fmt.Printf("Seed: %d\n", seed)

}

func TestMultiGraphCreate(t *testing.T) {

    Setup()
    // create a node
    node1 := &Node{}
    node2 := &Node{}


    node1.Connect(node2, 1.0)
    node1.Print()

}
