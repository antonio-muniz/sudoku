package sudoku

type Column struct {
	index    int
	elements [][]*Element
}

func (column Column) ForEachElement(function func(*Element)) {
	for _, row := range column.elements {
		element := row[column.index]
		function(element)
	}
}

func (column Column) UsedValues() ValueSet {
	usedValues := NewValueSet()
	column.ForEachElement(func(element *Element) {
		if element.HasValue() {
			usedValues.Add(element.Value())
		}
	})
	return usedValues
}
