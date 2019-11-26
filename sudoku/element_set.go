package sudoku

type ElementSet interface {
	ForEachElement(func(*Element))
	UsedValues() ValueSet
}
