package sudoku_test

import (
	"testing"

	"github.com/antonio-muniz/sudoku-solver/sudoku"
	"github.com/stretchr/testify/require"
)

type scenarioSolve struct {
	description      string
	matrix           [][]int
	expectedSolution [][]int
}

func TestSolve(t *testing.T) {
	scenarios := []scenarioSolve{
		scenarioSolveCompleteMatrix(),
		scenarioSolveBeginnerMatrix(),
		scenarioSolveEasyMatrix(),
		scenarioSolveMediumMatrix(),
		//scenarioSolveHardMatrix(),
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			solution := sudoku.Solve(scenario.matrix)
			require.Equal(t, scenario.expectedSolution, solution)
		})
	}
}

func scenarioSolveCompleteMatrix() scenarioSolve {
	return scenarioSolve{
		description: "solve_complete_matrix",
		matrix: [][]int{
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
		expectedSolution: [][]int{
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

func scenarioSolveBeginnerMatrix() scenarioSolve {
	return scenarioSolve{
		description: "solve_beginner_matrix",
		matrix: [][]int{
			{6, 0, 4, 1, 0, 0, 0, 3, 8},
			{9, 2, 0, 8, 0, 5, 0, 1, 6},
			{0, 3, 8, 0, 0, 6, 5, 9, 7},
			{0, 0, 1, 0, 6, 4, 9, 0, 2},
			{4, 9, 2, 5, 8, 3, 7, 6, 1},
			{7, 6, 0, 9, 0, 2, 8, 4, 3},
			{8, 4, 6, 3, 7, 0, 1, 2, 5},
			{0, 7, 0, 4, 2, 0, 6, 8, 0},
			{0, 0, 9, 6, 5, 8, 3, 0, 4},
		},
		expectedSolution: [][]int{
			{6, 5, 4, 1, 9, 7, 2, 3, 8},
			{9, 2, 7, 8, 3, 5, 4, 1, 6},
			{1, 3, 8, 2, 4, 6, 5, 9, 7},
			{3, 8, 1, 7, 6, 4, 9, 5, 2},
			{4, 9, 2, 5, 8, 3, 7, 6, 1},
			{7, 6, 5, 9, 1, 2, 8, 4, 3},
			{8, 4, 6, 3, 7, 9, 1, 2, 5},
			{5, 7, 3, 4, 2, 1, 6, 8, 9},
			{2, 1, 9, 6, 5, 8, 3, 7, 4},
		},
	}
}

func scenarioSolveEasyMatrix() scenarioSolve {
	return scenarioSolve{
		description: "solve_easy_matrix",
		matrix: [][]int{
			{4, 5, 9, 1, 7, 3, 2, 8, 6},
			{3, 0, 1, 6, 8, 0, 0, 0, 0},
			{0, 0, 6, 5, 0, 9, 0, 0, 1},
			{6, 3, 4, 2, 9, 0, 5, 7, 0},
			{5, 0, 0, 8, 0, 0, 3, 0, 0},
			{8, 1, 7, 0, 0, 4, 6, 0, 0},
			{1, 4, 0, 7, 3, 6, 0, 0, 9},
			{0, 0, 0, 9, 0, 5, 0, 0, 0},
			{0, 6, 3, 0, 0, 0, 1, 0, 7},
		},
		expectedSolution: [][]int{
			{4, 5, 9, 1, 7, 3, 2, 8, 6},
			{3, 7, 1, 6, 8, 2, 9, 4, 5},
			{2, 8, 6, 5, 4, 9, 7, 3, 1},
			{6, 3, 4, 2, 9, 1, 5, 7, 8},
			{5, 9, 2, 8, 6, 7, 3, 1, 4},
			{8, 1, 7, 3, 5, 4, 6, 9, 2},
			{1, 4, 5, 7, 3, 6, 8, 2, 9},
			{7, 2, 8, 9, 1, 5, 4, 6, 3},
			{9, 6, 3, 4, 2, 8, 1, 5, 7},
		},
	}
}

func scenarioSolveMediumMatrix() scenarioSolve {
	return scenarioSolve{
		description: "solve_medium_matrix",
		matrix: [][]int{
			{0, 4, 6, 0, 8, 5, 0, 0, 2},
			{0, 9, 0, 0, 0, 4, 8, 6, 3},
			{0, 0, 0, 2, 6, 0, 9, 4, 0},
			{0, 0, 9, 3, 0, 7, 0, 0, 4},
			{0, 0, 0, 0, 0, 1, 0, 0, 0},
			{0, 3, 0, 6, 0, 0, 0, 0, 7},
			{0, 6, 5, 8, 0, 0, 0, 2, 1},
			{4, 0, 3, 1, 5, 0, 7, 0, 9},
			{9, 0, 0, 4, 0, 0, 0, 0, 0},
		},
		expectedSolution: [][]int{
			{3, 4, 6, 9, 8, 5, 1, 7, 2},
			{5, 9, 2, 7, 1, 4, 8, 6, 3},
			{1, 8, 7, 2, 6, 3, 9, 4, 5},
			{8, 5, 9, 3, 2, 7, 6, 1, 4},
			{6, 7, 4, 5, 9, 1, 2, 3, 8},
			{2, 3, 1, 6, 4, 8, 5, 9, 7},
			{7, 6, 5, 8, 3, 9, 4, 2, 1},
			{4, 2, 3, 1, 5, 6, 7, 8, 9},
			{9, 1, 8, 4, 7, 2, 3, 5, 6},
		},
	}
}

func scenarioSolveHardMatrix() scenarioSolve {
	return scenarioSolve{
		description: "solve_hard_matrix",
		matrix: [][]int{
			{3, 5, 0, 0, 0, 0, 0, 8, 0},
			{0, 0, 0, 0, 5, 8, 9, 0, 2},
			{0, 0, 0, 0, 0, 7, 0, 0, 0},
			{0, 1, 0, 0, 0, 5, 0, 4, 0},
			{0, 0, 0, 7, 0, 0, 0, 9, 0},
			{0, 0, 9, 0, 1, 2, 0, 0, 7},
			{0, 0, 0, 0, 0, 0, 0, 0, 3},
			{6, 0, 0, 2, 0, 0, 0, 0, 0},
			{0, 0, 2, 3, 0, 0, 1, 0, 8},
		},
		expectedSolution: [][]int{
			{3, 5, 1, 9, 2, 6, 7, 8, 4},
			{4, 7, 6, 1, 5, 8, 9, 3, 2},
			{9, 2, 8, 4, 3, 7, 6, 1, 5},
			{7, 1, 3, 8, 9, 5, 2, 4, 6},
			{2, 6, 5, 7, 4, 3, 8, 9, 1},
			{8, 4, 9, 6, 1, 2, 3, 5, 7},
			{1, 8, 7, 5, 6, 9, 4, 2, 3},
			{6, 3, 4, 2, 8, 1, 5, 7, 9},
			{5, 9, 2, 3, 7, 4, 1, 6, 8},
		},
	}
}
