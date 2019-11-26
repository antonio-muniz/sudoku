package sudoku

import (
	"fmt"
	"strconv"
	"strings"
)

type ValueSet struct {
	values map[int]bool
}

func NewValueSet(values ...int) ValueSet {
	set := ValueSet{values: make(map[int]bool)}
	for _, value := range values {
		set.Add(value)
	}
	return set
}

func (set ValueSet) Add(value int) {
	set.values[value] = true
}

func (set ValueSet) Contains(value int) bool {
	return set.values[value]
}

func (set ValueSet) Except(other ValueSet) ValueSet {
	result := NewValueSet()
	for value := range set.values {
		if !other.Contains(value) {
			result.Add(value)
		}
	}
	return result
}

func (set ValueSet) PowerSet() []ValueSet {
	var subsets []ValueSet

	emptySet := NewValueSet()
	subsets = append(subsets, emptySet)

	for value := range set.values {
		singleSet := NewValueSet(value)

		var increasedSubsets []ValueSet
		for _, subset := range subsets {
			increasedSubset := subset.Union(singleSet)
			increasedSubsets = append(increasedSubsets, increasedSubset)
		}
		subsets = append(subsets, increasedSubsets...)
	}

	return subsets
}

func (set ValueSet) Remove(value int) {
	delete(set.values, value)
}

func (set ValueSet) Size() int {
	return len(set.values)
}

func (set ValueSet) SubsetOf(other ValueSet) bool {
	for value := range set.values {
		if !other.Contains(value) {
			return false
		}
	}
	return true
}

func (set ValueSet) String() string {
	var numbers []string
	for value := range set.values {
		numbers = append(numbers, strconv.FormatInt(int64(value), 10))
	}
	return fmt.Sprintf("{%s}", strings.Join(numbers, ", "))
}

func (set ValueSet) Union(other ValueSet) ValueSet {
	unionSet := NewValueSet()
	for value := range set.values {
		unionSet.Add(value)
	}
	for value := range other.values {
		unionSet.Add(value)
	}
	return unionSet
}

func (set ValueSet) Values() []int {
	var values []int
	for value := range set.values {
		values = append(values, value)
	}
	return values
}
