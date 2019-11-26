package sudoku

var allValues = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func Solve(inputMatrix [][]int) [][]int {
	matrix := NewMatrix(inputMatrix)

	removeInvalidCandidates(matrix)

	for round := 0; round < 100; round++ {
		removeConflictsWithNakedCandidates(matrix, 3)
		removeConflictsWithNakedCandidates(matrix, 2)
		removeConflictsWithNakedCandidates(matrix, 1)
		revealHiddenSingles(matrix)
	}

	return matrix.Values()
}

func removeInvalidCandidates(matrix Matrix) {
	matrix.ForEachElement(func(element *Element) {
		if !element.HasValue() {
			return
		}

		value := element.Value()

		elementSets := matrix.ElementSetsOf(element)
		for _, elementSet := range elementSets {
			elementSet.ForEachElement(func(neighbor *Element) {
				if neighbor == element {
					return
				}
				neighbor.RemoveCandidate(value)
			})
		}
	})
}

func revealHiddenSingles(matrix Matrix) {
	matrix.ForEachElement(func(element *Element) {
		if element.HasValue() {
			return
		}
		elementSets := matrix.ElementSetsOf(element)
		for _, elementSet := range elementSets {
			for _, candidate := range element.Candidates().Values() {
				var occurrences int
				elementSet.ForEachElement(func(neighbor *Element) {
					for _, neighborCandidate := range neighbor.Candidates().Values() {
						if neighbor == element {
							continue
						}
						if neighborCandidate == candidate {
							occurrences++
						}
					}
				})
				if occurrences == 0 {
					element.SetCandidates([]int{candidate})
					return
				}
			}
		}
	})
}

func removeConflictsWithNakedCandidates(matrix Matrix, groupSize int) {
	allPossibleValues := NewValueSet(allValues...)
	matrix.ForEachElementSet(func(elementSet ElementSet) {
		unusedValues := allPossibleValues.Except(elementSet.UsedValues())
		var unusedValueSets []ValueSet
		for _, subset := range unusedValues.PowerSet() {
			if subset.Size() == groupSize {
				unusedValueSets = append(unusedValueSets, subset)
			}
		}

		for _, valueSet := range unusedValueSets {
			candidateElements := make(map[*Element]bool)
			elementSet.ForEachElement(func(element *Element) {
				if valueSet.SubsetOf(element.Candidates()) {
					candidateElements[element] = true
				}
			})
			if len(candidateElements) == groupSize {
				elementSet.ForEachElement(func(element *Element) {
					if candidateElements[element] {
						return
					}
					cleanCandidates := element.Candidates().Except(valueSet)
					element.SetCandidates(cleanCandidates.Values())
				})
			}
		}
	})
}
