package main

import (
	"github.com/vectaport/flowgraph"
	"time"
)

func tbi(x flowgraph.Edge) {

	node := flowgraph.NewNode("tbi", nil, []*flowgraph.Edge{&x}, nil)

	var i int = 0
	for {
		if (i>10) { break }

		if node.Rdy(){
			node.Tracef("writing x.Data: %d\n", x.Val.(int))
			node.TraceVals()
			x.Data <- x.Val
			x.Rdy = false
			x.Val = x.Val.(int) + 1
			i = i + 1
		}

		node.Select()

	}
}

func tbo(a flowgraph.Edge) {

	node := flowgraph.NewNode("tbo", []*flowgraph.Edge{&a}, nil, nil)

	for {
		if node.Rdy() {
			node.Tracef("writing a.Ack\n")
			node.TraceVals()
			a.Ack <- true
			a.Rdy = false
		}

		node.Select()

	}

}

func main() {

	flowgraph.Debug = false
	flowgraph.Indent = false

	a := flowgraph.NewEdge("ae",nil)
	b := flowgraph.NewConstEdge("be",1000)
	x := flowgraph.NewEdge("xe",nil)

	a.Val = 0

	go tbi(a)
	// go flowgraph.FuncConst(b)
	go flowgraph.FuncAdd(a, b, x)
	go tbo(x)

	time.Sleep(1000000000)

}
