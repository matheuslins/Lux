package main

func BuildHashMap(arr []float64) map[float64]bool {
	hMap := make(map[float64]bool)

	for _, item := range arr {
		hMap[item] = true
	}
	return hMap
}
