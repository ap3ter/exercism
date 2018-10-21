package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	rowSep = "\n"
	colSep = " "
)

// Matrix holds a matrix
type Matrix [][]int

// New creates a new matrix out of a string with rows seperated by new line charcter
func New(str string) (Matrix, error) {
	// Get rows
	strR := (strings.Split(strings.Trim(str, " "), rowSep))
	// Get number of rows
	rowC := len(strR)
	// Create a matirx with number of rows
	m := make(Matrix, rowC)
	// Get number of columns
	colC := len(strings.Split(strR[0], colSep))
	// Initialize the matrix with zero values
	for i := 0; i < rowC; i++ {
		r := make([]int, colC)
		m[i] = r
	}
	// Fill up the matrix
	for i := range strR {
		strC := strings.Split(strings.Trim(strR[i], " "), colSep)
		// All column must have same length
		if len(strC) != colC {
			return nil, fmt.Errorf("All rows must have same column %v", str)
		}
		for j := range strC {
			v, err := strconv.Atoi(strC[j])
			if err != nil {
				return nil,
					fmt.Errorf("Non integer value detected while setting value for matrix at row %v and column %v", i, j)
			}
			m.Set(i, j, v)
		}
	}
	return m, nil
}

// Rows gets the rows of the matrix in a matrix form
func (m Matrix) Rows() Matrix {
	n := make(Matrix, len(m))
	for i := range n {
		c := make([]int, len(m[0]))
		for j := range c {
			c[j] = m[i][j]
		}
		n[i] = c
	}
	return n
}

// Cols gets the columns of the matrix in a matrix form
func (m Matrix) Cols() Matrix {
	// New matrix with no of rows as no of col of originall matrix
	n := make(Matrix, len(m[0]))
	for i := range n {
		c := make([]int, len(m))
		for j := range c {
			c[j] = m[j][i]
		}
		n[i] = c
	}
	return n
}

// Set sets the value of element denoted by row and column in a matrix
func (m *Matrix) Set(r int, c int, v int) bool {
	(*m)[r][c] = v
	return true
}
