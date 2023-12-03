package main

import (
	"testing"

	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

// Global variables for benchmarking
var g *gorgonia.ExprGraph
var x, y *gorgonia.Node

func newNN(g *gorgonia.ExprGraph) *gorgonia.Node {
	// Define your neural network architecture here
	// Example implementation:
	w := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(784, 10)) // Example shape
	b := gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(1, 10))   // Example shape

	return gorgonia.Must(gorgonia.Add(gorgonia.Must(gorgonia.Mul(x, w)), b))
}

func fwd(x *gorgonia.Node) error {

	return nil
}

func BenchmarkNeuralNetwork(b *testing.B) {
	g = gorgonia.NewGraph()
	x = gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(1, 784)) // Example shape
	y = gorgonia.NewMatrix(g, tensor.Float64, gorgonia.WithShape(1, 10))  // Example shape

	net := newNN(g)
	if err := net.fwd(x); err != nil {
		b.Fatal(err)
	}

	machine := gorgonia.NewTapeMachine(g)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := machine.RunAll(); err != nil {
			b.Fatal(err)
		}
		machine.Reset() // Reset the VM after each run
	}
}
