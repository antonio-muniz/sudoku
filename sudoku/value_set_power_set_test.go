package sudoku_test

import (
	"testing"

	"github.com/antonio-muniz/sudoku-solver/sudoku"
	"github.com/stretchr/testify/require"
)

type scenarioPowerSet struct {
	description      string
	set              sudoku.ValueSet
	expectedPowerSet []sudoku.ValueSet
}

func TestPowerSet(t *testing.T) {
	scenarios := []scenarioPowerSet{
		scenarioPowerSetEmptySet(),
		scenarioPowerSetSingleElementSet(),
		scenarioPowerSetFourElementSet(),
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			powerSet := scenario.set.PowerSet()
			require.ElementsMatch(t, scenario.expectedPowerSet, powerSet)
		})
	}
}

func scenarioPowerSetEmptySet() scenarioPowerSet {
	return scenarioPowerSet{
		description: "power_set_of_empty_set",
		set:         sudoku.NewValueSet(),
		expectedPowerSet: []sudoku.ValueSet{
			sudoku.NewValueSet(),
		},
	}
}

func scenarioPowerSetSingleElementSet() scenarioPowerSet {
	return scenarioPowerSet{
		description: "power_set_of_single_element_set",
		set:         sudoku.NewValueSet(7),
		expectedPowerSet: []sudoku.ValueSet{
			sudoku.NewValueSet(),
			sudoku.NewValueSet(7),
		},
	}
}

func scenarioPowerSetFourElementSet() scenarioPowerSet {
	return scenarioPowerSet{
		description: "power_set_of_four_element_set",
		set:         sudoku.NewValueSet(1, 2, 3, 4),
		expectedPowerSet: []sudoku.ValueSet{
			sudoku.NewValueSet(),
			sudoku.NewValueSet(1),
			sudoku.NewValueSet(2),
			sudoku.NewValueSet(1, 2),
			sudoku.NewValueSet(3),
			sudoku.NewValueSet(1, 3),
			sudoku.NewValueSet(2, 3),
			sudoku.NewValueSet(1, 2, 3),
			sudoku.NewValueSet(4),
			sudoku.NewValueSet(1, 4),
			sudoku.NewValueSet(2, 4),
			sudoku.NewValueSet(1, 2, 4),
			sudoku.NewValueSet(3, 4),
			sudoku.NewValueSet(1, 3, 4),
			sudoku.NewValueSet(2, 3, 4),
			sudoku.NewValueSet(1, 2, 3, 4),
		},
	}
}
