# Lux

## Computational complexity
 
 Basing on concepts of algorithmic complexity, it's necessary consider just the dominants terms.
 That is, doesn't matter how much elementary operations your code have, because is irrelevant when
 we have a larger input. So, in view of this, the tasks has following complexity:

 ### Task 1
 
 
 ### Task 2
 
```bash
Linear time: O(n)
```
Means that as the input grows the algorithms take proportionally longer to complete.
In this case, it was used a HashMap to make the intersect and difference functions:

 ```go

func BuildHashMap(arr []float64) map[float64]bool {
	hMap := make(map[float64]bool)

	for _, item := range arr {
		hMap[item] = true
	}
	return hMap
}

```
## Run

```go

cd [task1/task2]
go run *.go

```