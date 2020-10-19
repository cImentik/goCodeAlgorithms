package Minimum_Domino_Rotations_For_Equal_Row

import "testing"

func Test1(t *testing.T) {
	A := []int{2, 1, 2, 4, 2, 2}
	B := []int{5, 2, 6, 2, 3, 2}
	expected := 2
	actual := minDominoRotations(A, B)
	if actual != expected {
		t.Error("\nExpected:", expected, "\nActual:", actual)
	}
}
