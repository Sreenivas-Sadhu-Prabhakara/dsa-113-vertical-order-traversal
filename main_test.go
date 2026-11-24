package main

import (
	"reflect"
	"testing"
)

func TestVerticalOrder(t *testing.T) {
	root := tn(3, tn(9, nil, nil), tn(20, tn(15, nil, nil), tn(7, nil, nil)))
	want := [][]int{{9}, {3, 15}, {20}, {7}}
	if got := VerticalOrder(root); !reflect.DeepEqual(got, want) {
		t.Errorf("VerticalOrder = %v, want %v", got, want)
	}

	// Same row and column: values sort ascending (LC 987 rule).
	r2 := tn(1, tn(2, tn(4, nil, nil), tn(6, nil, nil)), tn(3, tn(5, nil, nil), tn(7, nil, nil)))
	// Nodes 6 and 5 share column 0, row 2 -> sorted as 5,6.
	want2 := [][]int{{4}, {2}, {1, 5, 6}, {3}, {7}}
	if got := VerticalOrder(r2); !reflect.DeepEqual(got, want2) {
		t.Errorf("VerticalOrder = %v, want %v", got, want2)
	}

	if got := VerticalOrder(nil); !reflect.DeepEqual(got, [][]int{}) {
		t.Errorf("nil: %v", got)
	}
}
