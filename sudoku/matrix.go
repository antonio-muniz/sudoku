package sudoku

type Matrix struct {
	elements [][]*Element
	rows     []Row
	columns  []Column
	regions  []Region
}

func NewMatrix(source [][]int) Matrix {
	var (
		rows    []Row
		columns []Column
		regions []Region
	)

	elements := make([][]*Element, 9)

	for rowIndex, row := range source {
		for columnIndex, value := range row {
			point := Point{rowIndex: rowIndex, columnIndex: columnIndex}
			var element *Element
			if value == 0 {
				element = &Element{point: point, candidates: NewValueSet(allValues...)}
			} else {
				element = &Element{point: point, candidates: NewValueSet(value)}
			}
			elements[rowIndex] = append(elements[rowIndex], element)
		}
	}

	for index := 0; index < 9; index++ {
		rows = append(rows, Row{index: index, elements: elements})
		columns = append(columns, Column{index: index, elements: elements})
	}

	for rowIndex := 0; rowIndex < 9; rowIndex += 3 {
		for columnIndex := 0; columnIndex < 9; columnIndex += 3 {
			region := Region{
				start:    Point{rowIndex: rowIndex, columnIndex: columnIndex},
				end:      Point{rowIndex: rowIndex + 2, columnIndex: columnIndex + 2},
				elements: elements,
			}
			regions = append(regions, region)
		}
	}

	matrix := Matrix{
		elements: elements,
		rows:     rows,
		columns:  columns,
		regions:  regions,
	}

	return matrix
}

func (matrix Matrix) Values() [][]int {
	values := make([][]int, 9)

	matrix.ForEachElement(func(element *Element) {
		values[element.point.rowIndex] = append(values[element.point.rowIndex], element.Value())
	})

	return values
}

func (matrix Matrix) rowOf(element *Element) ElementSet {
	return matrix.rows[element.point.rowIndex]
}

func (matrix Matrix) columnOf(element *Element) ElementSet {
	return matrix.columns[element.point.columnIndex]
}

func (matrix Matrix) regionOf(element *Element) ElementSet {
	regionRowStartIndex := element.point.rowIndex / 3
	regionColumnStartIndex := element.point.columnIndex / 3
	regionIndex := (regionRowStartIndex * 3) + regionColumnStartIndex
	return matrix.regions[regionIndex]
}

func (matrix Matrix) ElementSetsOf(element *Element) []ElementSet {
	elementSets := []ElementSet{
		matrix.rowOf(element),
		matrix.columnOf(element),
		matrix.regionOf(element),
	}
	return elementSets
}

func (matrix Matrix) ForEachElement(function func(*Element)) {
	for _, elementRow := range matrix.elements {
		for _, element := range elementRow {
			function(element)
		}
	}
}

func (matrix Matrix) ForEachElementSet(function func(ElementSet)) {
	for _, row := range matrix.rows {
		function(row)
	}
	for _, column := range matrix.columns {
		function(column)
	}
	for _, region := range matrix.regions {
		function(region)
	}
}
