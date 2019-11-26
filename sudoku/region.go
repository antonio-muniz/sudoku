package sudoku

type Region struct {
	start    Point
	end      Point
	elements [][]*Element
}

func (region Region) ForEachElement(function func(*Element)) {
	for rowIndex := region.start.rowIndex; rowIndex <= region.end.rowIndex; rowIndex++ {
		for columnIndex := region.start.columnIndex; columnIndex <= region.end.columnIndex; columnIndex++ {
			element := region.elements[rowIndex][columnIndex]
			function(element)
		}
	}
}

func (region Region) UsedValues() ValueSet {
	usedValues := NewValueSet()
	region.ForEachElement(func(element *Element) {
		if element.HasValue() {
			usedValues.Add(element.Value())
		}
	})
	return usedValues
}
