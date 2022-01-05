package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
		return
	}
	return
}

//// Same determines whether the trees
//// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	// Hacer los walk
	ch1 := walkTreeIntoChannel(t1)
	ch2 := walkTreeIntoChannel(t2)

	// Comparar las señales en ambos channels
	return compareUsingSlices(ch1, ch2)
	//return compareUsingSelect(ch1, ch2)
	//return compareUsingSelect2(ch1, ch2)

}

func walkTreeIntoChannel(t *tree.Tree) chan int {
	ch := make(chan int)
	go func() {
		Walk(t, ch)
		close(ch)
	}()
	return ch
}

func compareUsingSelect2(ch1 chan int, ch2 chan int) bool {
	for v1 := range ch1 {
		v2 := <-ch2

		fmt.Println(fmt.Sprintf("comapring %v with %v", v1, v2))

		if v1 != v2 {
			return false
		}
	}
	_, ok2 := <-ch2
	if ok2 {
		return false
	}

	return true
}

func compareUsingSelect(ch1 chan int, ch2 chan int) bool {
	for v1 := range ch1 {
		v2 := <-ch2

		fmt.Println(fmt.Sprintf("comapring %v with %v", v1, v2))

		if v1 != v2 {
			return false
		}
	}
	_, ok2 := <-ch2
	if ok2 {
		return false
	}

	return true
}

func compareUsingSlices(ch1 chan int, ch2 chan int) bool {
	// Recibir las señales de los channels y guardarlas en slices
	g := new(errgroup.Group)
	signals1 := []int{}
	g.Go(func() error {
		for v := range ch1 {
			signals1 = append(signals1, v)
		}
		return nil
	})

	signals2 := []int{}
	g.Go(func() error {
		for v := range ch2 {
			signals2 = append(signals2, v)
		}
		return nil
	})

	_ = g.Wait()

	// Comparar slices
	return areEqual(signals1, signals2)
}

func areEqual(signals1 []int, signals2 []int) bool {
	if len(signals1) != len(signals2) {
		return false
	}

	for i := 0; i < len(signals1); i++ {
		if signals1[i] != signals2[i] {
			return false
		}
	}

	return true
}

// MAIN DE ACA PARA ABAJO

func leer10() {
	// Create a new channel ch and kick off the walker:
	// go Walk(tree.New(1), ch)
	// Then read and print 10 values from the channel.
	//It should be the numbers 1, 2, 3, ..., 10.
	ch := make(chan int)
	go func() {
		Walk(tree.New(10), ch)
		close(ch)
	}()

	for signal := range ch {
		fmt.Print(signal, " ")
	}
}

func main() {
	leer10()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(tree.New(10))
	//}
	return

}
