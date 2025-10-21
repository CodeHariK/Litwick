package main

import (
	"fmt"

	nnfs "github.com/codeharik/nnfs/src"
)

func main() {
	weights := [][]float32{
		{0.2, 0.8, -0.5, 1},
		{0.5, -0.91, 0.26, -0.5},
		{-0.26, -0.27, 0.17, 0.87},
	}
	bias := []float32{2, 3, 0.5}

	l, err := nnfs.NewLayer(weights, bias)
	if err != nil {
		panic(err)
	}

	out, err := l.Calc([]float32{1, 2, 3, 2.5})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s", fmt.Sprint(out))
}
