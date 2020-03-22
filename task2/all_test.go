package main

import (
	"reflect"
	"testing"
)


func TestIntersect(t *testing.T) {
	firstInterval := []float64{1, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8}
	secondInterval := []float64{1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 2.0, 2.1, 2.2, 2.3, 2.4, 2.5}
	sameValues := intersect(firstInterval, secondInterval)
	if !reflect.DeepEqual(sameValues, []float64{1.4, 1.5, 1.6, 1.7, 1.8}) {
		t.Error("Result intersect not matched!")
	}
}


func TestDifference(t *testing.T) {
	firstInterval := []float64{1, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8}
	secondInterval := []float64{1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 2.0, 2.1, 2.2, 2.3, 2.4, 2.5}

	sameValues := intersect(firstInterval, secondInterval)
	mergedIntervals := append(firstInterval, secondInterval...)

	diffValues := difference(sameValues, mergedIntervals)
	if !reflect.DeepEqual(diffValues, []float64{1, 1.1, 1.2, 1.3, 1.9, 2, 2.1, 2.2, 2.3, 2.4, 2.5}) {
		t.Error("Result difference not matched!")
	}
}

