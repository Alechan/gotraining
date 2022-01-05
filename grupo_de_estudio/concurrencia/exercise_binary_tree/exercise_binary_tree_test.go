package main

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/tour/tree"
	"testing"
)

func TestWalk(t *testing.T) {
	expectedSignalsGoTour := []int{1, 1, 2, 3, 5, 8, 13}
	expectedSignalsLeafNode := []int{0}

	tests := []struct {
		name            string
		tree            *tree.Tree
		expectedSignals []int
	}{
		{
			name:            "tree 1 from go tour",
			tree:            getFirstTreeExample(),
			expectedSignals: expectedSignalsGoTour,
		},
		{
			name:            "tree 2 from go tour",
			tree:            getSecondTreeExample(),
			expectedSignals: expectedSignalsGoTour,
		},
		{
			name:            "leaf tree",
			tree:            getLeafTree(),
			expectedSignals: expectedSignalsLeafNode,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// WHEN
			ch := make(chan int, 10)
			Walk(tt.tree, ch)
			close(ch)

			// THEN
			actualAllSignals := []int{}
			for actualSignal := range ch {
				actualAllSignals = append(actualAllSignals, actualSignal)
			}

			for i := range tt.expectedSignals {
				require.Equal(t, tt.expectedSignals[i], actualAllSignals[i])
			}
		})
	}
}

func TestSame(t *testing.T) {
	tests := []struct {
		name string
		t1   *tree.Tree
		t2   *tree.Tree
		want bool
	}{
		{
			name: "Nil trees should be true",
			t1:   nil,
			t2:   nil,
			want: true,
		},
		{
			name: "Nil tree and leaf-tree should be false",
			t1:   nil,
			t2:   getLeafTree(),
			want: false,
		},
		{
			name: "Go tour examples comparison should be true",
			t1:   getFirstTreeExample(),
			t2:   getSecondTreeExample(),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Same(tt.t1, tt.t2); got != tt.want {
				t.Errorf("Same() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compareUsingSelect(t *testing.T) {
	tests := []struct {
		name     string
		signals1 []int
		signals2 []int
		want     bool
	}{
		{
			name:     "Empty channels should return true",
			signals1: []int{},
			signals2: []int{},
			want:     true,
		},
		{
			name:     "First channel with one signal and second without should return false",
			signals1: []int{1},
			signals2: []int{},
			want:     false,
		},
		{
			name:     "Second channel with one signal and first without should return false",
			signals1: []int{},
			signals2: []int{1},
			want:     false,
		},
		{
			name:     "Both with one different element should return false ",
			signals1: []int{0},
			signals2: []int{1},
			want:     false,
		},
		{
			name:     "Both with 3 same elements should return true ",
			signals1: []int{1, 2, 3},
			signals2: []int{1, 2, 3},
			want:     true,
		},
		{
			name:     "Both with 3 different elements should return false ",
			signals1: []int{3, 2, 1},
			signals2: []int{1, 2, 3},
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch1 := newChannelWithElementsAsSignals(tt.signals1)
			ch2 := newChannelWithElementsAsSignals(tt.signals2)

			if got := compareUsingSelect(ch1, ch2); got != tt.want {
				t.Errorf("compareUsingSelect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getLeafTree() *tree.Tree {
	return &tree.Tree{
		Left:  nil,
		Value: 0,
		Right: nil,
	}
}

func getSecondTreeExample() *tree.Tree {
	hoja1 := &tree.Tree{
		Left:  nil,
		Value: 1,
		Right: nil,
	}
	hoja2 := &tree.Tree{
		Left:  nil,
		Value: 2,
		Right: nil,
	}
	rama1 := &tree.Tree{
		Left:  hoja1,
		Value: 1,
		Right: hoja2,
	}
	hoja5 := &tree.Tree{
		Left:  nil,
		Value: 5,
		Right: nil,
	}
	rama3 := &tree.Tree{
		Left:  rama1,
		Value: 3,
		Right: hoja5,
	}
	hoja13 := &tree.Tree{
		Left:  nil,
		Value: 13,
		Right: nil,
	}
	rama8 := &tree.Tree{
		Left:  rama3,
		Value: 8,
		Right: hoja13,
	}
	return rama8
}

func getFirstTreeExample() *tree.Tree {
	hoja1 := &tree.Tree{
		Left:  nil,
		Value: 1,
		Right: nil,
	}
	hoja2 := &tree.Tree{
		Left:  nil,
		Value: 2,
		Right: nil,
	}
	rama1 := &tree.Tree{
		Left:  hoja1,
		Value: 1,
		Right: hoja2,
	}
	hoja5 := &tree.Tree{
		Left:  nil,
		Value: 5,
		Right: nil,
	}
	hoja13 := &tree.Tree{
		Left:  nil,
		Value: 13,
		Right: nil,
	}
	rama8 := &tree.Tree{
		Left:  hoja5,
		Value: 8,
		Right: hoja13,
	}
	rama3 := &tree.Tree{
		Left:  rama1,
		Value: 3,
		Right: rama8,
	}
	return rama3
}

func newChannelWithElementsAsSignals(inputSlice []int) chan int {
	ch := make(chan int)
	go func() {
		for _, v := range inputSlice {
			ch <- v
		}
		// Cerramos el channel porque:
		// 1) Fuimos quienes lo creamos
		// 2) Somos los únicos que mandamos señales
		close(ch)
	}()
	return ch
}
