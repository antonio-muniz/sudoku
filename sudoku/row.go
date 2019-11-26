package sudoku

type Row struct {
	index    int
	elements [][]*Element
}

func (row Row) ForEachElement(function func(*Element)) {
	for _, element := range row.elements[row.index] {
		function(element)
	}
}

func (row Row) UsedValues() ValueSet {
	usedValues := NewValueSet()
	row.ForEachElement(func(element *Element) {
		if element.HasValue() {
			usedValues.Add(element.Value())
		}
	})
	return usedValues
}
