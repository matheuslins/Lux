package main

import (
	"fmt"
)


func difference(containsArray, fullArray []float64) []float64 {
	hashMap := BuildHashMap(containsArray)
	diff := PickElements(hashMap, fullArray, false)
	return diff
}

func intersect(firstInterval, secondInterval []float64) []float64 {
	hashMap := BuildHashMap(secondInterval)
	equals := PickElements(hashMap, firstInterval, true)
	return equals
}

func main()  {

	firstInterval := UntangleList([]float64{4.0, 8.2})
	secondInterval := UntangleList([]float64{6.1, 10.3})

	sameValues := intersect(firstInterval, secondInterval)
	fmt.Println(sameValues)

	mergedIntervals := append(firstInterval, secondInterval...)

	differentValues := difference(sameValues, mergedIntervals)
	fmt.Println(differentValues)


}
