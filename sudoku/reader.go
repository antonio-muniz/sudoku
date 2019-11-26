package sudoku

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Read(reader io.Reader) ([][]int, error) {
	matrix, err := parseMatrix(reader)
	if err != nil {
		return nil, err
	}

	err = ensureMatrixSize(matrix)
	if err != nil {
		return nil, err
	}

	return matrix, nil
}

func parseMatrix(reader io.Reader) ([][]int, error) {
	var matrix [][]int

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		rowLine := scanner.Text()
		row, err := parseRow(rowLine)
		if err != nil {
			return nil, fmt.Errorf("error reading row %d: %s", len(matrix), err)
		}
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading matrix: '%s'", err)
	}

	return matrix, nil
}

func parseRow(rowLine string) ([]int, error) {
	var row []int
	elementTokens := strings.Split(rowLine, " ")
	for columnNumber, elementToken := range elementTokens {
		element, err := parseElement(elementToken)
		if err != nil {
			return nil, fmt.Errorf("element '%s' at column %d is invalid: %s", elementToken, columnNumber, err)
		}
		row = append(row, element)
	}
	return row, nil
}

func parseElement(elementToken string) (int, error) {
	if elementToken == "_" {
		return 0, nil
	}
	element, err := strconv.ParseInt(elementToken, 10, 64)
	if err != nil {
		return 0, errors.New("cannot be parsed into integer")
	}
	if element <= 0 || element > 9 {
		return 0, errors.New("out of range")
	}
	return int(element), nil
}

func ensureMatrixSize(matrix [][]int) error {
	if rowCount := len(matrix); rowCount != 9 {
		return fmt.Errorf("matrix does not have 9 rows (found %d)", rowCount)
	}
	for rowIndex, row := range matrix {
		if elementCount := len(row); elementCount != 9 {
			return fmt.Errorf("row %d does not have 9 elements (found %d)", rowIndex, elementCount)
		}
	}
	return nil
}
