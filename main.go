/*
 * Title: Q5 Practise Sheet 7 Overkill
 * Author: Oluwatofunmi Yinka-Adebimpe
 * Date: 26 February 2023 21:35
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

const rows = 3
const cols = 3

// TODO: Find inverse for 4x4 and 5x5 matrices

func main() {
	matA := [3][3]int{}
	inputMatrix(&matA)
	displayMatrix(&matA, "matA")
	matrixInverse(&matA)
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

// 1 2 3 4; 5 6 7 8
func inputMatrix(mat *[3][3]int) {
	fmt.Printf("Input a 3 x 3 Matrix\n")
	fmt.Printf("Please seperate numbers with spaces and use ';' for columns\n")
	fmt.Printf("E.g. 1 2 3; 4 5 6; 7 8 9\n")

	input, err := getInput(">> matA = ", reader)
	if err != nil {
		log.Println(err)
	}

	// input := "10 11 14; 56 57 58; -20 5 6"

	rows := strings.Split(input, ";")
	for i, v := range rows {
		cols := strings.Split(strings.TrimSpace(v), " ")
		for j, el := range cols {
			var err error
			mat[i][j], err = strconv.Atoi(el)
			if err != nil {
				panic(fmt.Sprintf("%s is not a number", el))
			}
		}
	}
}

/*

A =
1 2 3
4 5 6
7 8 9

a11 a12 a13
a21 a22 a23
a31 a32 a33

A11
a22 a23
a32 a33

*/

// STEPS
// find minors of first row
// check if determinant is zero (using minors of row 1) if it is exit and show else continue
// find minors of 2nd and 3rd rows
// find matrix of cofactors
// take transpose of matrix of matrix of cofactors
// multiply each element of matrix of transpose with 1/determinant
func matrixInverse(mat *[3][3]int) {
	minors := [3][3]float64{}

	// Minors of row 1
	for i := 0; i < cols; i++ {
		minors[0][i] = determinant(matrixMinors(mat, 0, i))
	}

	// Determinant
	// det = a11*A11 - a12*A12 + a13*A13
	det := (float64(mat[0][0]) * minors[0][0]) - (float64(mat[0][1]) * minors[0][1]) + (float64(mat[0][2]) * minors[0][2])

	// fmt.Println("det =", det)

	// Check if there is an inverse for the matrix
	if det == 0 {
		fmt.Printf("\n")
		fmt.Println("There is no inverse for the matrix")
		fmt.Println("because the determinant is equal to zero")
		return
	}

	// Find minors for 2nd and 3rd rows
	for i := 1; i < rows; i++ {
		for j := 0; j < cols; j++ {
			minors[i][j] = determinant(matrixMinors(mat, i, j))
		}
	}
	// displayMatrixFloat64(&minors, "minors")

	// TODO: Reduce the amount of times the matrix is iterated over

	count := 1
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// Convert matrix of minors to matrix of cofactors
			if count%2 == 0 {
				// Even
				minors[i][j] *= -1
			}
			count++
		}
	}
	// displayMatrixFloat64(&minors, "cofactors")

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// Transpose
			if i < j {
				minors[i][j], minors[j][i] = minors[j][i], minors[i][j]
			}
		}
	}
	// displayMatrixFloat64(&minors, "cofactors transpose")

	// fmt.Println(det)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// Divide all elements by det
			minors[i][j] /= det
		}
	}

	displayMatrixFloat64(&minors, "inv(matA)")
}

func matrixMinors(mat *[3][3]int, i, j int) [4]float64 {
	minors := [4]float64{}
	a := 0
	for x := 0; x < rows; x++ {
		if x == i {
			continue
		}
		for y := 0; y < cols; y++ {
			if y == j {
				continue
			}

			// fmt.Println(minors, i, j)
			minors[a] = float64(mat[x][y])
			a++
		}
	}

	return minors
}

/*
A11 =
a22 a23
a32 a33

minors = [a22, a23, a32, a33]
*/
func determinant(minors [4]float64) float64 {
	return (minors[0] * minors[3]) - (minors[2] * minors[1])
}

// TODO: combine display matrix and display matrix float 64
func displayMatrix(mat *[3][3]int, name string) {
	fmt.Printf("\n%s =\n\n", name)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("\t%d", mat[i][j])
		}
		fmt.Printf("\n")
	}
}

func displayMatrixFloat64(mat *[3][3]float64, name string) {
	fmt.Printf("\n%s =\n\n", name)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("\t%10.4f", mat[i][j])
		}
		fmt.Printf("\n")
	}
}
