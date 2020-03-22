package main

import "fmt"


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

	firstInterval := UntangleList([]float64{4.0, 8.2}) // O(n)
	secondInterval := UntangleList([]float64{6.1, 10.3}) // O(n)

	sameValues := intersect(firstInterval, secondInterval) // O(n)
	fmt.Println("The same values are:")
	fmt.Println(sameValues)

	mergedIntervals := append(firstInterval, secondInterval...)
	differentValues := difference(sameValues, mergedIntervals) // O(n)
	fmt.Println("The different values are:")
	fmt.Println(differentValues)


}
