package utils

type Cell struct {
	R int
	C int
}

func (c Cell) Up(d int) Cell {
	return Cell{
		R: c.R - d,
		C: c.C,
	}
}
func (c Cell) Down(d int) Cell {
	return Cell{
		R: c.R + d,
		C: c.C,
	}
}
func (c Cell) Left(d int) Cell {
	return Cell{
		R: c.R,
		C: c.C - d,
	}
}
func (c Cell) Right(d int) Cell {
	return Cell{
		R: c.R,
		C: c.C + d,
	}
}
