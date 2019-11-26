package sudoku_test

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/antonio-muniz/sudoku-solver/sudoku"
)

type scenarioRead struct {
	description    string
	input          io.Reader
	expectedMatrix [][]int
	expectedError  string
}

func TestRead(t *testing.T) {
	scenarios := []scenarioRead{
		scenarioReadValidCompleteMatrix(),
		scenarioReadValidIncompleteMatrix(),
		scenarioReadMatrixWithMoreThan9Rows(),
		scenarioReadMatrixWithLessThan9Rows(),
		scenarioReadMatrixWithMoreThan9ElementsInRow(),
		scenarioReadMatrixWithLessThan9ElementsInRow(),
		scenarioReadMatrixWithUnparseableElement(),
		scenarioReadMatrixWithElementOutOfRange(),
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			matrix, err := sudoku.Read(scenario.input)

			if scenario.expectedError == "" {
				require.NoError(t, err)
				require.Equal(t, scenario.expectedMatrix, matrix)
			} else {
				require.EqualError(t, err, scenario.expectedError)
			}
		})
	}
}

func scenarioReadValidCompleteMatrix() scenarioRead {
	return scenarioRead{
		description: "read_valid_complete_matrix",
		input: readerWithLines([]string{
			"1 2 3 4 5 6 7 8 9",
			"4 5 6 7 8 9 1 2 3",
			"7 8 9 1 2 3 4 5 6",
			"2 3 4 5 6 7 8 9 1",
			"5 6 7 8 9 1 2 3 4",
			"8 9 1 2 3 4 5 6 7",
			"3 4 5 6 7 8 9 1 2",
			"6 7 8 9 1 2 3 4 5",
			"9 1 2 3 4 5 6 7 8",
		}),
		expectedMatrix: [][]int{
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{4, 5, 6, 7, 8, 9, 1, 2, 3},
			{7, 8, 9, 1, 2, 3, 4, 5, 6},
			{2, 3, 4, 5, 6, 7, 8, 9, 1},
			{5, 6, 7, 8, 9, 1, 2, 3, 4},
			{8, 9, 1, 2, 3, 4, 5, 6, 7},
			{3, 4, 5, 6, 7, 8, 9, 1, 2},
			{6, 7, 8, 9, 1, 2, 3, 4, 5},
			{9, 1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
}

func scenarioReadValidIncompleteMatrix() scenarioRead {
	return scenarioRead{
		description: "read_valid_incomplete_matrix",
		input: readerWithLines([]string{
			"1 2 3 4 5 6 7 8 9",
			"4 5 6 _ 8 9 1 2 3",
			"7 8 9 1 2 3 4 5 6",
			"2 3 4 5 6 7 8 9 1",
			"5 6 7 8 9 1 2 _ 4",
			"8 9 1 2 3 4 5 6 7",
			"3 4 5 6 7 8 9 1 2",
			"6 7 _ 9 1 2 3 4 5",
			"9 _ 2 3 4 5 6 7 8",
		}),
		expectedMatrix: [][]int{
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{4, 5, 6, 0, 8, 9, 1, 2, 3},
			{7, 8, 9, 1, 2, 3, 4, 5, 6},
			{2, 3, 4, 5, 6, 7, 8, 9, 1},
			{5, 6, 7, 8, 9, 1, 2, 0, 4},
			{8, 9, 1, 2, 3, 4, 5, 6, 7},
			{3, 4, 5, 6, 7, 8, 9, 1, 2},
			{6, 7, 0, 9, 1, 2, 3, 4, 5},
			{9, 0, 2, 3, 4, 5, 6, 7, 8},
		},
	}
}

func scenarioReadMatrixWithMoreThan9Rows() scenarioRead {
	return scenarioRead{
		description: "read_matrix_with_more_than_9_rows",
		input: readerWithLines([]string{
			"1 2 3 4 5 6 7 8 9",
			"4 5 6 7 8 9 1 2 3",
			"7 8 9 1 2 3 4 5 6",
			"2 3 4 5 6 7 8 9 1",
			"5 6 7 8 9 1 2 3 4",
			"8 9 1 2 3 4 5 6 7",
			"3 4 5 6 7 8 9 1 2",
			"6 7 8 9 1 2 3 4 5",
			"9 1 2 3 4 5 6 7 8",
			"9 1 2 3 4 5 6 7 8",
		}),
		expectedError: "matrix does not have 9 rows (found 10)",
	}
}

func scenarioReadMatrixWithLessThan9Rows() scenarioRead {
	return scenarioRead{
		description: "read_matrix_with_less_than_9_rows",
		input: readerWithLines([]string{
			"1 2 3 4 5 6 7 8 9",
			"4 5 6 7 8 9 1 2 3",
			"7 8 9 1 2 3 4 5 6",
			"2 3 4 5 6 7 8 9 1",
			"5 6 7 8 9 1 2 3 4",
			"8 9 1 2 3 4 5 6 7",
			"3 4 5 6 7 8 9 1 2",
			"6 7 8 9 1 2 3 4 5",
		}),
		expectedError: "matrix does not have 9 rows (found 8)",
	}
}

func scenarioReadMatrixWithMoreThan9ElementsInRow() scenarioRead {
	return scenarioRead{
		description: "read_matrix_with_more_than_9_elements_in_row",
		input: readerWithLines([]string{
			"1 2 3 4 5 6 7 8 9",
			"4 5 6 7 8 9 1 2 3",
			"7 8 9 1 2 3 4 5 6",
			"2 3 4 5 6 7 8 9 1",
			"5 6 7 8 9 1 2 3 4",
			"8 9 1 2 3 4 5 6 7 8",
			"3 4 5 6 7 8 9 1 2",
			"6 7 8 9 1 2 3 4 5",
			"9 1 2 3 4 5 6 7 8",
		}),
		expectedError: "row 5 does not have 9 elements (found 10)",
	}
}

func scenarioReadMatrixWithLessThan9ElementsInRow() scenarioRead {
	return scenarioRead{
		description: "read_matrix_with_less_than_9_elements_in_row",
		input: readerWithLines([]string{
			"1 2 3 4 5 6 7 8 9",
			"4 5 6 7 8 9 1 2 3",
			"7 8 9 1 2 3 4 5",
			"2 3 4 5 6 7 8 9 1",
			"5 6 7 8 9 1 2 3 4",
			"8 9 1 2 3 4 5 6 7",
			"3 4 5 6 7 8 9 1 2",
			"6 7 8 9 1 2 3 4 5",
			"9 1 2 3 4 5 6 7 8",
		}),
		expectedError: "row 2 does not have 9 elements (found 8)",
	}
}

func scenarioReadMatrixWithUnparseableElement() scenarioRead {
	return scenarioRead{
		description: "read_matrix_with_unparseable_element",
		input: readerWithLines([]string{
			"1 2 3 4 5 6 7 8 9",
			"4 5 6 X 8 9 1 2 3",
			"7 8 9 1 2 3 4 5 6",
			"2 3 4 5 6 7 8 9 1",
			"5 6 7 8 9 1 2 3 4",
			"8 9 1 2 3 4 5 6 7",
			"3 4 5 6 7 8 9 1 2",
			"6 7 8 9 1 2 3 4 5",
			"9 1 2 3 4 5 6 7 8",
		}),
		expectedError: "error reading row 1: element 'X' at column 3 is invalid: cannot be parsed into integer",
	}
}

func scenarioReadMatrixWithElementOutOfRange() scenarioRead {
	return scenarioRead{
		description: "read_matrix_with_element_out_of_range",
		input: readerWithLines([]string{
			"1 2 3 4 5 6 7 8 9",
			"4 5 6 7 8 9 1 2 3",
			"7 8 9 1 2 3 4 5 6",
			"2 3 4 5 6 7 8 9 1",
			"5 6 7 8 9 1 2 3 4",
			"8 9 1 2 3 4 5 6 7",
			"3 4 5 6 7 8 9 1 2",
			"6 7 8 9 1 2 0 4 5",
			"9 1 2 3 4 5 6 7 8",
		}),
		expectedError: "error reading row 7: element '0' at column 6 is invalid: out of range",
	}
}

func readerWithLines(lines []string) io.Reader {
	text := strings.Join(lines, "\n")
	reader := strings.NewReader(text)
	return reader
}
