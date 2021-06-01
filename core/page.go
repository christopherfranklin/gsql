package core

const PageSize int = 1024 // Max Rows per page
const MaxPages int = 100
const MaxRows int = PageSize * MaxPages

type PageResult int

const (
	PageSuccess PageResult = iota
	PageFull
)

type Page struct {
	rows []Row
}

func (p *Page) PageSize() int {
	return len(p.rows)
}

func (p *Page) AddRow(r Row) PageResult {
	if len(p.rows) >= PageSize {
		return PageFull
	}

	p.rows = append(p.rows, r)
	return PageSuccess
}
