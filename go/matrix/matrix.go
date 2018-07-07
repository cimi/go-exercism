package matrix

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Matrix struct {
	e [][]int
}

func (m Matrix) Rows() [][]int {
	rows := make([][]int, 0)
	for _, vals := range m.e {
		tmp := make([]int, len(vals))
		copy(tmp, vals)
		rows = append(rows, tmp)
	}
	return rows
}

func (m Matrix) Cols() [][]int {
	numRows := len(m.e)
	numCols := len(m.e[0])
	cols := make([][]int, numCols)
	for i := 0; i < numCols; i++ {
		cols[i] = make([]int, numRows)
	}
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			cols[col][row] = m.e[row][col]
		}
	}
	return cols
}

func (m *Matrix) Set(r, c, val int) bool {
	if r < 0 || c < 0 || r >= len(m.e) || c >= len(m.e[r]) {
		return false
	}
	m.e[r][c] = val
	return true
}

func parseLine(line string) ([]int, error) {
	result := make([]int, 0)
	for _, f := range strings.Fields(line) {
		num, err := strconv.Atoi(f)
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}
	return result, nil
}

func New(input string) (*Matrix, error) {
	elements := make([][]int, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		row, err := parseLine(strings.TrimSpace(scanner.Text()))
		if err != nil {
			return nil, err
		}
		elements = append(elements, row)
	}
	l := len(elements[0])
	for _, row := range elements {
		if len(row) != l {
			return nil, fmt.Errorf("Invalid row length %d, expected %d", len(row), l)
		}
	}
	return &Matrix{elements}, nil
}
