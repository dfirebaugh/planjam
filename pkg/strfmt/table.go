package strfmt

import (
	"strings"
)

type Row struct {
	Value string
}

type Column struct {
	Rows []Row
}

type Table struct {
	Columns []Column
}

const (
	horizontalDelimiter = "|"
	verticalDelimiter   = "-"
)

func NewTable() *Table {
	return &Table{
		Columns: make([]Column, 0),
	}
}

func (t *Table) AppendToColumn(column int, text string) {
	if column >= len(t.Columns) {
		t.Columns = append(t.Columns, Column{Rows: make([]Row, 0)})
	}
	t.Columns[column].Rows = append(t.Columns[column].Rows, Row{Value: text})
}

func (t Table) String() string {
	var s strings.Builder

	s.WriteString("\n")

	maxRows := 0
	for _, col := range t.Columns {
		if len(col.Rows) > maxRows {
			maxRows = len(col.Rows)
		}
	}

	columnWidths := make([]int, len(t.Columns))

	for colIndex, col := range t.Columns {
		for _, r := range col.Rows {
			if len(r.Value) > columnWidths[colIndex] {
				columnWidths[colIndex] = len(r.Value) + 2
			}
		}
	}

	for i := 0; i < maxRows; i++ {
		for colIndex, col := range t.Columns {
			s.WriteString(horizontalDelimiter)

			if i < len(col.Rows) {
				padding := columnWidths[colIndex] - len(col.Rows[i].Value) - 1
				if padding < 0 {
					padding = 0
				}
				s.WriteString(" " + col.Rows[i].Value + strings.Repeat(" ", padding))
			} else {
				s.WriteString(strings.Repeat(" ", columnWidths[colIndex]))
			}
		}
		s.WriteString(horizontalDelimiter)

		if i == 0 {
			s.WriteString("\n")
			for _, width := range columnWidths {
				s.WriteString(horizontalDelimiter + strings.Repeat(verticalDelimiter, width))
			}
			s.WriteString(horizontalDelimiter)
		}
		s.WriteString("\n")
	}

	return s.String()
}

func (t Table) ASCIIDocTable() string {
	var s strings.Builder

	s.WriteString("\n")

	maxRows := 0
	for _, col := range t.Columns {
		if len(col.Rows) > maxRows {
			maxRows = len(col.Rows)
		}
	}

	columnWidths := make([]int, len(t.Columns))

	for colIndex, col := range t.Columns {
		for _, r := range col.Rows {
			if len(r.Value) > columnWidths[colIndex] {
				columnWidths[colIndex] = len(r.Value) + 2
			}
		}
	}

	s.WriteString("|===")
	s.WriteString("\n")

	for i := 0; i < maxRows; i++ {
		for colIndex, col := range t.Columns {
			s.WriteString("|")

			if i < len(col.Rows) {
				padding := columnWidths[colIndex] - len(col.Rows[i].Value) - 1
				if padding < 0 {
					padding = 0
				}
				s.WriteString(" " + col.Rows[i].Value + strings.Repeat(" ", padding))
			} else {
				s.WriteString(strings.Repeat(" ", columnWidths[colIndex]))
			}
		}

		if i == 0 {
			s.WriteString("\n")
		}
		s.WriteString("\n")
	}

	s.WriteString("|===")
	s.WriteString("\n")

	return s.String()
}
func (t Table) MarkdownTable() string {
	var s strings.Builder

	s.WriteString("\n")

	maxRows := 0
	for _, col := range t.Columns {
		if len(col.Rows) > maxRows {
			maxRows = len(col.Rows)
		}
	}

	columnWidths := make([]int, len(t.Columns))

	// Calculate column widths
	for colIndex, col := range t.Columns {
		for _, r := range col.Rows {
			if len(r.Value) > columnWidths[colIndex] {
				columnWidths[colIndex] = len(r.Value)
			}
		}
	}

	// Write the header row
	for colIndex, col := range t.Columns {
		s.WriteString(horizontalDelimiter)
		if len(col.Rows) > 0 {
			padding := columnWidths[colIndex] - len(col.Rows[0].Value)
			s.WriteString(" " + col.Rows[0].Value + strings.Repeat(" ", padding))
		} else {
			s.WriteString(strings.Repeat(" ", columnWidths[colIndex]))
		}
	}
	s.WriteString(horizontalDelimiter + "\n")

	// Write the header separator row
	for _, width := range columnWidths {
		s.WriteString(horizontalDelimiter + strings.Repeat(verticalDelimiter, width+1))
	}
	s.WriteString(horizontalDelimiter + "\n")

	// Write the data rows
	for i := 1; i < maxRows; i++ { // Start from 1 to skip the header row
		for colIndex, col := range t.Columns {
			s.WriteString(horizontalDelimiter)
			if i < len(col.Rows) {
				padding := columnWidths[colIndex] - len(col.Rows[i].Value)
				if padding < 0 {
					padding = 0
				}
				s.WriteString(" " + col.Rows[i].Value + strings.Repeat(" ", padding))
			} else {
				s.WriteString(strings.Repeat(" ", columnWidths[colIndex]+1))
			}
		}
		s.WriteString(horizontalDelimiter + "\n")
	}

	return s.String()
}
