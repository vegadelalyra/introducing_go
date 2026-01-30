package main

import "fmt"

// Node is our union type interface.
// The presence of the unexported 'isNode()' method ensures
// that only types in this package can satisfy this interface.
type Node interface {
	isNode()
}

// nodeBase is the "empty struct" you mentioned.
// We embed this into our concrete types to satisfy the Node interface.
type nodeBase struct{}

func (nodeBase) isNode() {}

// --- Concrete Node Types ---

// NodeA can contain a slice of other Nodes (A, B, or C)
type NodeA struct {
	nodeBase
	Name     string
	Children []Node
}

// NodeB is a leaf-ish node that holds an integer
type NodeB struct {
    nodeBase 
    Value int
}

// NodeC holds a string value 
type NodeC struct {
    nodeBase 
    Data string 
}

// --- Usage --- 

func main() {
    // Creating a heterogeneous tree 
    tree := &NodeA{
        Name: "Root",
        Children: []Node{
            NodeB{Value: 42}, 
            NodeC{Data: "Hello Go"},
            &NodeA{
                Name: "Subtree",
                Children: []Node{
                    NodeB{Value: 100},
                },
            },
        },
    }

    printTree(tree, 0)
}

// printTree demonstrates how to handle the union type using a type switch.
func printTree(n Node, depth int) {
    indent := ""
    for range depth {
        indent += "    "
    }

    switch v := n.(type) {
    case *NodeA:
        fmt.Printf("%sNodeA: %s\n", indent, v.Name)
        for _, child := range v.Children {
            printTree(child, depth+1)
        }
    case NodeB: 
        fmt.Printf("%sNodeB (Value: %d)\n", indent, v.Value)
    case NodeC: 
        fmt.Printf("%sNodeC (Data: %s)\n", indent, v.Data)
    default:
        fmt.Printf("%sUnknown Node Type\n", indent)
    }
}
