package main

func main() {
	zeroOnePack()
}

func zeroOnePack() {
	N := 10
	V := 100
	F := make([][]int, N+1)
	data := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		// Ci Wi
		data[i] = make([]int, 2)
		F[i] = make([]int, V+1)
	}
	// according to golang initialize the data after it allocate.
	for i := 1; i <= N; i++ {
		for v := data[i][0]; v <= V; v++ {
			F[i][v] = max(F[i-1][v], F[i-1][v-data[i][0]]+data[i][1])
		}
	}
}

func ZeroOnePackArray() {
	N := 10
	V := 100
	F := make([]int, V+1)
	data := make([][]int, N+1)
	for i := 1; i <= N; i++ {
		for v := V; v <= data[i][0]; v-- {
			F[v] = max(F[v], F[v-data[i][0]]+data[i][1])
		}
	}
}

func max(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}
