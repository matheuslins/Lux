package main

import "fmt"


func compareIntervals(intervals [][]float64) [][]float64{
	intervalsSize := len(intervals)
	var mergedIntervals [][]float64

	for i := 0; i < len(intervals); i++ {
		if intervalsSize < 2 {
			mergedIntervals = append(mergedIntervals, intervals[i])
			return mergedIntervals

		}
		nextElement := i + 1
		firstInterval := UntangleList(intervals[i])
		secondInterval := UntangleList(intervals[nextElement])
		sameValues := Intersect(firstInterval, secondInterval)

		if sameValues {
			fullIntervals := append(firstInterval, secondInterval...)
			extremitiesOfIntervals := GetExtremitiesOfIntervals(fullIntervals)
			mergedIntervals = append(mergedIntervals, extremitiesOfIntervals)
			intervals = RemoveInOrder(intervals, nextElement)
		}

		if !sameValues {
			mergedIntervals = append(mergedIntervals, intervals[i])
			intervals = RemoveInOrder(intervals, nextElement)
		}
	}

	return mergedIntervals
}

func main()  {

	intervals := [][]float64{{1.0, 3.5}, {4.0, 8.2}, {3.5, 3.8}, {6.1, 10.3}, {0.1, 0.4}}
	OrderList(intervals) // O(n^2)

	mergedIntervals := compareIntervals(intervals) // O(n^2)
	fmt.Println(mergedIntervals)

}
