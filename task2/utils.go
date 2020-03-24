package main

import "math"


func Intersect(firstInterval, secondInterval []float64) bool {
	hashMap := BuildHashMap(secondInterval)
	equals := PickElements(hashMap, firstInterval, true)
	return equals
}

func UntangleList(list []float64) []float64 {
	var listPointer = &list[0]
	return _floatRange(listPointer, list[GetLastItem(list)])
}

func GetLastItem(inter []float64) int{
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

func PickElements(hashMap map[float64]bool, iterator []float64, conditional bool) bool{
	for _, item := range iterator {
		if _, ok := hashMap[item]; ok == conditional {
			return true
		}
	}
	return false
}


func RemoveInOrder(slice [][]float64, position int) [][]float64 {
	return append(slice[:position], slice[position+1:]...)
}

func GetExtremitiesOfIntervals(interval []float64) []float64{
	return []float64{interval[0], interval[len(interval)-1]}
}