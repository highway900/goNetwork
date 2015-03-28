package goNetwork


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
    Edges []IEdge
    Nodes []INode
}


func (g *Graph) AddNode(data interface{}) INode {
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


func (graph *Graph) Connect(n0 INode, n1 INode, weight float64, id string, key string) {
    edge := &Edge{
        n0,
        n1,
        weight,
        id,
        key,
    }

    n0.Connect(edge)
    n1.Connect(edge)

    graph.Edges = append(graph.Edges, edge)
}
