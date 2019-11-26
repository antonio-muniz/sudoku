package sudoku_test

import (
	"testing"

	"github.com/antonio-muniz/sudoku-solver/sudoku"
	"github.com/stretchr/testify/require"
)

type scenarioExcept struct {
	description    string
	set            sudoku.ValueSet
	otherSet       sudoku.ValueSet
	expectedResult sudoku.ValueSet
}

func TestExcept(t *testing.T) {
	scenarios := []scenarioExcept{
		scenarioExceptSetAndSubset(),
		scenarioExceptSetAndEmptySet(),
		scenarioExceptSetAndEqualSet(),
		scenarioExceptSetAndContainingSet(),
		scenarioExceptEmptySetAndNonEmptySet(),
		scenarioExceptSetAndNonSubset(),
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			result := scenario.set.Except(scenario.otherSet)
			require.Equal(t, scenario.expectedResult, result)
		})
	}
}

func scenarioExceptSetAndSubset() scenarioExcept {
	return scenarioExcept{
		description:    "set_except_subset",
		set:            sudoku.NewValueSet(1, 2, 3, 4),
		otherSet:       sudoku.NewValueSet(1, 3),
		expectedResult: sudoku.NewValueSet(2, 4),
	}
}

func scenarioExceptSetAndEmptySet() scenarioExcept {
	return scenarioExcept{
		description:    "set_except_subset",
		set:            sudoku.NewValueSet(1, 2, 3, 4),
		otherSet:       sudoku.NewValueSet(),
		expectedResult: sudoku.NewValueSet(1, 2, 3, 4),
	}
}

func scenarioExceptSetAndEqualSet() scenarioExcept {
	return scenarioExcept{
		description:    "set_except_equal_set",
		set:            sudoku.NewValueSet(1, 2, 3, 4),
		otherSet:       sudoku.NewValueSet(1, 2, 3, 4),
		expectedResult: sudoku.NewValueSet(),
	}
}

func scenarioExceptSetAndContainingSet() scenarioExcept {
	return scenarioExcept{
		description:    "set_except_containing_set",
		set:            sudoku.NewValueSet(1, 2, 3, 4),
		otherSet:       sudoku.NewValueSet(1, 2, 3, 4, 5),
		expectedResult: sudoku.NewValueSet(),
	}
}

func scenarioExceptEmptySetAndNonEmptySet() scenarioExcept {
	return scenarioExcept{
		description:    "empty_set_except_non_empty_set",
		set:            sudoku.NewValueSet(),
		otherSet:       sudoku.NewValueSet(1, 2, 3, 4),
		expectedResult: sudoku.NewValueSet(),
	}
}

func scenarioExceptSetAndNonSubset() scenarioExcept {
	return scenarioExcept{
		description:    "set_except_non_subset",
		set:            sudoku.NewValueSet(1, 2, 3, 4),
		otherSet:       sudoku.NewValueSet(4, 5, 6),
		expectedResult: sudoku.NewValueSet(1, 2, 3),
	}
}
