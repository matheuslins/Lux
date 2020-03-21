package main

import (
	"fmt"
)


func Difference(containsArray, fullArray []float64) (diff []float64) {
	hashMap := BuildHashMap(containsArray)

	for _, item := range fullArray {
		if _, ok := hashMap[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
}

func Intersect(firstInterval, secondInterval []float64) (equals []float64) {
	hashMap := BuildHashMap(secondInterval)

	for _, item := range firstInterval {
		if _, ok := hashMap[item]; ok {
			equals = append(equals, item)
		}
	}

	return equals
}

func main()  {

	firstInterval := UntangleList([]float64{4.0, 8.2})
	secondInterval := UntangleList([]float64{6.1, 10.3})

	sameValues := Intersect(firstInterval, secondInterval)
	fmt.Println(sameValues)

	mergedIntervals := append(firstInterval, secondInterval...)

	differentValues := Difference(sameValues, mergedIntervals)
	fmt.Println(differentValues)


}
