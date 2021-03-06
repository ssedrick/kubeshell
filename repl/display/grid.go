package display

/*****************************************
 * Code ported from Rust repo:
 * https://github.com/ogham/rust-term-grid
 * Thanks @ogham for putting that great
 * library together
 *****************************************/

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

type Alignment int

const (
	Left Alignment = iota
	Right
)

type Cell struct {
	content   string
	width     int
	alignment Alignment
}

func NewCell(content string) *Cell {
	width := utf8.RuneCountInString(content)
	alignment := Left

	return &Cell{
		content,
		width,
		alignment,
	}
}

type Direction int

const (
	LeftToRight Direction = iota
	TopToBottom
)

type Padding struct {
	spaces int
	text   string
}

func NewWhitespacePadding(i int) *Padding {
	return &Padding{
		text:   "",
		spaces: i,
	}
}

func NewTextPadding(t string) *Padding {
	return &Padding{
		text:   t,
		spaces: 0,
	}
}

func (p *Padding) Width() int {
	return p.spaces + utf8.RuneCountInString(p.text)
}

func (p *Padding) Padding(alignment Alignment) string {
	switch alignment {
	case Left:
		return fmt.Sprintf("%s%s", strings.Repeat(" ", p.spaces), p.text)
	case Right:
		return fmt.Sprintf("%s%s", p.text, strings.Repeat(" ", p.spaces))
	}
	return ""
}

type GridOptions struct {
	direction Direction
	padding   *Padding
}

func DefaultOptions() *GridOptions {
	return &GridOptions{
		direction: LeftToRight,
		padding:   NewWhitespacePadding(2),
	}
}

func (g *GridOptions) ChangeDirection(direction Direction) {
	g.direction = direction
}

func (g *GridOptions) ChangePadding(padding *Padding) {
	g.padding = padding
}

type Dimensions struct {
	lines  int
	widths []int
}

func sum(arr []int) int {
	result := 0
	for _, i := range arr {
		result += i
	}
	return result
}

func (d *Dimensions) TotalWidth(p *Padding) int {
	cols := len(d.widths)
	if cols == 0 {
		return cols
	}
	sep_width := p.Width() * (cols - 1)
	return sum(d.widths) + sep_width
}

type Grid struct {
	options    *GridOptions
	cells      []*Cell
	widestCell int
	totalWidth int
}

func NewGrid() *Grid {
	cells := []*Cell{}
	return &Grid{
		options:    DefaultOptions(),
		cells:      cells,
		widestCell: 0,
		totalWidth: 0,
	}
}

func NewGridWithOptions(o *GridOptions) *Grid {
	cells := []*Cell{}
	return &Grid{
		options:    o,
		cells:      cells,
		widestCell: 0,
		totalWidth: 0,
	}
}

func (g *Grid) AddCell(c *Cell) {
	w := c.width
	if w > g.widestCell {
		g.widestCell = w
	}
	g.totalWidth += w
	g.cells = append(g.cells, c)
}

func (g *Grid) Sort() {
	sort.Slice(g.cells, func(i, j int) bool {
		return g.cells[i].content < g.cells[j].content
	})
}
