package main

import "sort"


func BuildHashMap(arr []float64) map[float64]bool {
	hMap := make(map[float64]bool)

	for _, item := range arr {
		hMap[item] = true
	}
	return hMap
}


func OrderList(list [][]float64) [][]float64{
	sort.Slice(list, func(i, j int) bool {
		return list[i][0] < list[j][0]
	})
	return list
}