package main

import (
	"fmt"
	"log"

	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

// Define the neural network structure
type nn struct {
	g      *gorgonia.ExprGraph
	w0, w1 *gorgonia.Node
	b0, b1 *gorgonia.Node
	pred   *gorgonia.Node
}

// Create a new neural network
func newNN(g *gorgonia.ExprGraph) *nn {
	// Initialize weights and biases
	w0 := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(784, 256), gorgonia.WithName("w0"), gorgonia.WithInit(gorgonia.GlorotU(1)))
	b0 := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(256), gorgonia.WithName("b0"), gorgonia.WithInit(gorgonia.Zeroes()))

	w1 := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(256, 10), gorgonia.WithName("w1"), gorgonia.WithInit(gorgonia.GlorotU(1)))
	b1 := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(10), gorgonia.WithName("b1"), gorgonia.WithInit(gorgonia.Zeroes()))

	return &nn{
		g:  g,
		w0: w0,
		b0: b0,
		w1: w1,
		b1: b1,
	}
}

// Define the forward pass of the network
func (net *nn) fwd(x *gorgonia.Node) (err error) {
	// First layer
	layer1 := gorgonia.Must(gorgonia.Rectify(gorgonia.Must(gorgonia.Add(gorgonia.Must(gorgonia.Mul(x, net.w0)), net.b0))))

	// Output layer
	net.pred = gorgonia.Must(gorgonia.SoftMax(gorgonia.Must(gorgonia.Add(gorgonia.Must(gorgonia.Mul(layer1, net.w1)), net.b1))))
	return
}

func main() {
	batchSize := 32 // Set the batchSize value here

	g := gorgonia.NewGraph()

	// Define input and output nodes
	x := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(batchSize, 784), gorgonia.WithName("x"))
	y := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(batchSize, 10), gorgonia.WithName("y"))

	// Create the neural network
	net := newNN(g)

	// Define the forward pass
	if err := net.fwd(x); err != nil {
		log.Fatal(err)
	}

	// Define the cost function (e.g., cross-entropy loss)
	loss := gorgonia.Must(gorgonia.Neg(gorgonia.Must(gorgonia.Sum(gorgonia.Must(gorgonia.HadamardProd(y, gorgonia.Must(gorgonia.Log(net.pred))))))))

	// Create a VM to execute the graph
	machine := gorgonia.NewTapeMachine(g)

	// Run the forward pass
	if err := machine.RunAll(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Completed forward pass")
	fmt.Println(loss.Value())
}
