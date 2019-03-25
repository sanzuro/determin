package main

import "fmt"

var mat = make([]int, 0)

func det(a []int, r int) int {
	var LocalDet = 0
	if r == 2 {
		LocalDet = a[0]*a[3] - a[2]*a[1]
	} else {
		for i := 0; i < r; i++ {
			newMat := make([]int, 0)
			for j := range a {
				if j/r != i && j%r != i {
					newMat = append(newMat, a[j])
				}
			}
			if i%2 == 0 {
				LocalDet += a[i] * det(newMat, r-1)
			} else {
				LocalDet -= a[i] * det(newMat, r-1)
			}
		}
	}
	return LocalDet
}
func main() {
	mat = append(mat, 1, 1, 1, 1, 1, 1, 1, 1, 1)
	fmt.Printf("%v", det(mat, 3))
}
