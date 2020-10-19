package Minimum_Domino_Rotations_For_Equal_Row

func minDominoRotations(A []int, B []int) int {
	count := count(A, B)
	result := -1
	for i := range count {
		if count[i] == len(A) {
			cost := cost(A, B, i)
			if result == -1 || cost < result {
				result = cost
			}
		}
	}
	return result
}

func count(A []int, B []int) [7]int {
	var count [7]int
	for i, v := range A {
		count[v]++
		if A[i] != B[i] {
			count[B[i]]++
		}
	}
	return count
}

func cost(A []int, B []int, value int) int {
	costA := 0
	costB := 0
	for i, v := range A {
		if v != value {
			costA++
		}
		if value != B[i] {
			costB++
		}
	}
	if costA < costB {
		return costA
	}
	return costB
}
