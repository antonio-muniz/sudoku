package sudoku_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/antonio-muniz/sudoku-solver/sudoku"
)

type scenarioSubsetOf struct {
	description    string
	set            sudoku.ValueSet
	otherSet       sudoku.ValueSet
	expectedResult bool
}

func TestSubsetOf(t *testing.T) {
	scenarios := []scenarioSubsetOf{
		scenarioSubsetOfSetOnContainingSet(),
		scenarioSubsetOfSetOnEqualSet(),
		scenarioSubsetOfEmptySetOnEmptySet(),
		scenarioSubsetOfEmptySetOnNonEmptySet(),
		scenarioSubsetOfNonEmptySetOnEmptySet(),
		scenarioSubsetOfSetOnNonContainingSet(),
		scenarioSubsetOfSetOnSubset(),
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			result := scenario.set.SubsetOf(scenario.otherSet)
			require.Equal(t, scenario.expectedResult, result)
		})
	}
}

func scenarioSubsetOfSetOnContainingSet() scenarioSubsetOf {
	return scenarioSubsetOf{
		description:    "set_subset_of_containing_set",
		set:            sudoku.NewValueSet(1, 3),
		otherSet:       sudoku.NewValueSet(1, 2, 3, 4),
		expectedResult: true,
	}
}

func scenarioSubsetOfSetOnEqualSet() scenarioSubsetOf {
	return scenarioSubsetOf{
		description:    "set_subset_of_equal_set",
		set:            sudoku.NewValueSet(1, 2, 3, 4),
		otherSet:       sudoku.NewValueSet(1, 2, 3, 4),
		expectedResult: true,
	}
}

func scenarioSubsetOfEmptySetOnEmptySet() scenarioSubsetOf {
	return scenarioSubsetOf{
		description:    "empty_set_subset_of_empty_set",
		set:            sudoku.NewValueSet(),
		otherSet:       sudoku.NewValueSet(),
		expectedResult: true,
	}
}

func scenarioSubsetOfEmptySetOnNonEmptySet() scenarioSubsetOf {
	return scenarioSubsetOf{
		description:    "empty_set_subset_of_non_empty_set",
		set:            sudoku.NewValueSet(),
		otherSet:       sudoku.NewValueSet(1, 2, 3, 4),
		expectedResult: true,
	}
}

func scenarioSubsetOfNonEmptySetOnEmptySet() scenarioSubsetOf {
	return scenarioSubsetOf{
		description:    "non_empty_set_subset_of_empty_set",
		set:            sudoku.NewValueSet(1, 2, 3, 4),
		otherSet:       sudoku.NewValueSet(),
		expectedResult: false,
	}
}

func scenarioSubsetOfSetOnNonContainingSet() scenarioSubsetOf {
	return scenarioSubsetOf{
		description:    "set_subset_of_non_containing_set",
		set:            sudoku.NewValueSet(1, 2, 3, 4),
		otherSet:       sudoku.NewValueSet(1, 3, 4, 5),
		expectedResult: false,
	}
}

func scenarioSubsetOfSetOnSubset() scenarioSubsetOf {
	return scenarioSubsetOf{
		description:    "set_subset_of_subset",
		set:            sudoku.NewValueSet(1, 2, 3, 4),
		otherSet:       sudoku.NewValueSet(1, 2, 3),
		expectedResult: false,
	}
}
