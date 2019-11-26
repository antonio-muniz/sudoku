package sudoku_test

import (
	"testing"

	"github.com/antonio-muniz/sudoku-solver/sudoku"
	"github.com/stretchr/testify/require"
)

type scenarioUnion struct {
	description    string
	set            sudoku.ValueSet
	otherSet       sudoku.ValueSet
	expectedResult sudoku.ValueSet
}

func TestUnion(t *testing.T) {
	scenarios := []scenarioUnion{
		scenarioUnionSetAndNonContainingSet(),
		scenarioUnionSetAndEqualSet(),
		scenarioUnionSetAndEmptySet(),
		scenarioUnionSetAndContainingSet(),
		scenarioUnionSetAndSubset(),
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			result := scenario.set.Union(scenario.otherSet)
			require.Equal(t, scenario.expectedResult, result)
		})
	}
}

func scenarioUnionSetAndNonContainingSet() scenarioUnion {
	return scenarioUnion{
		description:    "union_of_set_and_non_containing_set",
		set:            sudoku.NewValueSet(1, 2, 3, 4),
		otherSet:       sudoku.NewValueSet(1, 2, 3, 5),
		expectedResult: sudoku.NewValueSet(1, 2, 3, 4, 5),
	}
}

func scenarioUnionSetAndEqualSet() scenarioUnion {
	return scenarioUnion{
		description:    "union_of_set_and_equal_set",
		set:            sudoku.NewValueSet(1, 2, 3, 4),
		otherSet:       sudoku.NewValueSet(1, 2, 3, 4),
		expectedResult: sudoku.NewValueSet(1, 2, 3, 4),
	}
}

func scenarioUnionSetAndEmptySet() scenarioUnion {
	return scenarioUnion{
		description:    "union_of_set_and_empty_set",
		set:            sudoku.NewValueSet(1, 2, 3, 4),
		otherSet:       sudoku.NewValueSet(),
		expectedResult: sudoku.NewValueSet(1, 2, 3, 4),
	}
}

func scenarioUnionSetAndContainingSet() scenarioUnion {
	return scenarioUnion{
		description:    "union_of_set_and_containing_set",
		set:            sudoku.NewValueSet(1, 2, 3),
		otherSet:       sudoku.NewValueSet(1, 2, 3, 4),
		expectedResult: sudoku.NewValueSet(1, 2, 3, 4),
	}
}

func scenarioUnionSetAndSubset() scenarioUnion {
	return scenarioUnion{
		description:    "union_of_set_and_subset",
		set:            sudoku.NewValueSet(1, 2, 3, 4),
		otherSet:       sudoku.NewValueSet(2, 3, 4),
		expectedResult: sudoku.NewValueSet(1, 2, 3, 4),
	}
}
