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

## Outputs

### Task 1

### Task 2

```
If the inputs are [4.0, 8.2] and [6.1, 10.3]

###

The same values are:
[6.1 6.2 6.3 6.4 6.5 6.6 6.7 6.8 6.9 7 7.1 7.2 7.3 7.4 7.5 7.6 7.7 7.8 7.9 8 8.1 8.2]

The different values are:
[4 4.1 4.2 4.3 4.4 4.5 4.6 4.7 4.8 4.9 5 5.1 5.2 5.3 5.4 5.5 5.6 5.7 5.8 5.9 6 8.3 8.4 8.5 8.6 8.7 8.8 8.9 9 9.1 9.2 9.3 9.4 9.5 9.6 9.7 9.8 9.9 10 10.1 10.2 10.3]
```


## Run Tests

```go

cd [task1/task2]
go test

```
