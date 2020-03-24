package main

import (
	"reflect"
	"testing"
)


func TestMerged(t *testing.T) {
	intervals := [][]float64{{0.1, 0.4}, {1.0, 3.5}, {3.5, 3.8}, {4.0, 8.2}, {6.1, 10.3}}
	mergedIntervals := compareIntervals(intervals)
	if !reflect.DeepEqual(mergedIntervals, [][]float64{{0.1, 0.4}, {1.0, 3.8}, {4.0, 10.3}}) {
		t.Error("Result intersect not matched!")
	}
}
