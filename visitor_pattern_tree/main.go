package main

import (
	"fmt"
)

// --- 1. The Visitor Interface ---
// This defines an action for every type in your "union".
type NodeVisitor interface {
    VisitNodeA(*NodeA) 
    VisitNodeB(*NodeB) 
    VisitNodeC(*NodeC) 
}

// --- 2. The Updated Union Interface 
type Node interface {
    // Accept allows a visitor to interact with the node 
    Accept(NodeVisitor)
}

// --- 3. Concrete Nodes --- 

type NodeA struct {
    Name string 
    Children []Node 
}
func (n *NodeA) Accept(v NodeVisitor) { v.VisitNodeA(n) }

type NodeB struct {
    Value int
}
func (n *NodeB) Accept(v NodeVisitor) { v.VisitNodeB(n) }

type NodeC struct {
    Data string 
}
func (n *NodeC) Accept(v NodeVisitor) { v.VisitNodeC(n) }

// --- 4. Concrete Visitor: Tree Printer --- 
type TreePrinter struct {
    depth int 
}

func (p *TreePrinter) printIndent() {
    for range p.depth {
        fmt.Print("    ")
    }
}

func (p *TreePrinter) VisitNodeA(n *NodeA) {
    p.printIndent()
    fmt.Printf("Node A: %s\n", n.Name)
    p.depth++ 
    for _, child := range n.Children {
        child.Accept(p) // The recursion happens here
    }
    p.depth--
}

func (p *TreePrinter) VisitNodeB(n *NodeB) {
    p.printIndent()
    fmt.Printf("NodeB (Value: %d)\n", n.Value)
}

func (p *TreePrinter) VisitNodeC(n *NodeC) {
    p.printIndent()
    fmt.Printf("NodeC (Data: %s)\n", n.Data)
}

// Usage
func main() {
    tree := &NodeA{
        Name: "Root",
        Children: []Node{
            &NodeB{Value: 42},
            &NodeC{Data: "Visitor Pattern"},
            &NodeA{
                Name: "Subtree",
                Children: []Node{&NodeB{Value: 99}},
            },
        },
    }
    // We just pass the visitor to the root 
    printer := &TreePrinter{}
    tree.Accept(printer)
}

