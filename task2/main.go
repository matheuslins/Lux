package main

import "fmt"


func compareIntervals(intervals [][]float64) [][]float64{
	intervalsSize := &intervals
	var mergedIntervals [][]float64

	for len(*intervalsSize) > 0 {
		if len(*intervalsSize) < 2 {
			mergedIntervals = append(mergedIntervals, intervals[0])
			return mergedIntervals

		}
		firstInterval := UntangleList(intervals[0])
		secondInterval := UntangleList(intervals[1])
		sameValues := Intersect(firstInterval, secondInterval)

		if sameValues {
			fullIntervals := append(firstInterval, secondInterval...)
			extremitiesOfIntervals := GetExtremitiesOfIntervals(fullIntervals)
			mergedIntervals = append(mergedIntervals, extremitiesOfIntervals)
			intervals = RemoveInOrder(intervals, 0)
			intervals = RemoveInOrder(intervals, 0)
		}

		if !sameValues {
			mergedIntervals = append(mergedIntervals, intervals[0])
			intervals = RemoveInOrder(intervals, 0)
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
