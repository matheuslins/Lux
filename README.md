# Lux

## Computational complexity
 
 Basing on concepts of algorithmic complexity, it's necessary consider just the dominants terms.
 That is, doesn't matter how much elementary operations your code have, because is irrelevant when
 we have a larger input. So, in view of this, the tasks has following complexity:

 ### Task 1
 
 
 ### Task 2
 
```bash
Quadratic time: O(n^2)
```
Means that as the input grows the algorithms take n^2 time longer to complete.
In this case, I needed to read a list of list:

 ```go

func compareIntervals(intervals [][]float64) [][]float64{
    ...
	for len(*intervalsSize) > 0 {
		...
		firstInterval := UntangleList(intervals[0])
...

UntangleList > _floatRange

func _floatRange(start *float64, stop float64) [] float64{
	...
	for math.Round(startValue * size) <= math.Round(stop * size) {
		...
	}

```
## Run

Put the project inside your $GOPATH, and run: 
```go

cd [task1/task2]
go run main.go

```

## Outputs

### Task 1

### Task 2

```
If the input is: [[1.0, 3.5], [4.0, 8.2], [3.5, 3.8], [6.1, 10.3], [0.1, 0.4]]

###

The output is:
[[0.1, 0.4], [1.0, 3.8], [4.0, 10.3]]

```


## Run Tests

```go

cd [task1/task2]
go test

```
