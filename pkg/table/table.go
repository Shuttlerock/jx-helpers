package table

import (
	"fmt"
	"io"
	"unicode/utf8"

	"github.com/shuttlerock/jx-helpers/v3/pkg/stripansi"
)

type Table struct {
	Out          io.Writer
	Rows         [][]string
	ColumnWidths []int
	ColumnAlign  []int
	Separator    string
}

func CreateTable(out io.Writer) Table {
	return Table{
		Out:       out,
		Separator: " ",
	}
}

// Clear removes all rows while preserving the layout
func (t *Table) Clear() {
	t.Rows = [][]string{}
}

// AddRow adds a new row to the table
func (t *Table) AddRow(col ...string) {
	t.Rows = append(t.Rows, col)
}

func (t *Table) Render() {
	// lets figure out the max widths of each column
	for _, row := range t.Rows {
		for ci, rawCol := range row {
			col := stripansi.Strip(rawCol)
			l := utf8.RuneCountInString(col)
			t.ColumnWidths = ensureArrayCanContain(t.ColumnWidths, ci)
			if l > t.ColumnWidths[ci] {
				t.ColumnWidths[ci] = l
			}
		}
	}

	out := t.Out
	for _, row := range t.Rows {
		lastColumn := len(row) - 1
		for ci, col := range row {
			if ci > 0 {
				fmt.Fprint(out, t.Separator)
			}
			l := t.ColumnWidths[ci]

			// lets exclude colors from the length padding/aligning so lets add
			// the number of color characters to the width
			plainCol := stripansi.Strip(col)
			l += len(col) - len(plainCol)

			align := t.GetColumnAlign(ci)
			if ci >= lastColumn && align != ALIGN_CENTER && align != ALIGN_RIGHT {
				fmt.Fprint(out, col)
			} else {
				fmt.Fprint(out, Pad(col, " ", l, align))
			}
		}
		fmt.Fprint(out, "\n")
	}
}

// SetColumnsAligns sets the alignment of the columns
func (t *Table) SetColumnsAligns(colAligns []int) {
	t.ColumnAlign = colAligns
}

// GetColumnAlign return the column alignment
func (t *Table) GetColumnAlign(i int) int {
	t.ColumnAlign = ensureArrayCanContain(t.ColumnAlign, i)
	return t.ColumnAlign[i]
}

// SetColumnAlign sets the column alignment for the given column index
func (t *Table) SetColumnAlign(i int, align int) {
	t.ColumnAlign = ensureArrayCanContain(t.ColumnAlign, i)
	t.ColumnAlign[i] = align
}

func ensureArrayCanContain(array []int, idx int) []int {
	diff := idx + 1 - len(array)
	for i := 0; i < diff; i++ {
		array = append(array, 0)
	}
	return array
}
