package sudoku

type Element struct {
	point      Point
	candidates ValueSet
}

func (element *Element) Candidates() ValueSet {
	return NewValueSet().Union(element.candidates)
}

func (element *Element) HasValue() bool {
	return element.candidates.Size() == 1
}

func (element *Element) RemoveCandidate(value int) {
	element.candidates.Remove(value)
}

func (element *Element) SetCandidates(values []int) {
	element.candidates = NewValueSet(values...)
}

func (element *Element) Value() int {
	if element.candidates.Size() == 1 {
		return element.candidates.Values()[0]
	}
	return 0
}
