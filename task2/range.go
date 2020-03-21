package main

import "math"


func UntangleList(list []float64) []float64 {
	var listPointer = &list[0]
	return _floatRange(listPointer, list[_getLastItem(list)])
}

func _getLastItem(inter []float64) int{
	return len(inter) -1
}

func _floatRange(start *float64, stop float64) [] float64{
	var listOfRange []float64
	startValue := *start
	size := 100.0

	for math.Round(startValue * size) <= math.Round(stop * size) {
		listOfRange = append(listOfRange, math.Round(startValue * size) / size)
		startValue += 0.1
	}

	return listOfRange
}

