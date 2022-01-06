package formatter

import (
	"bytes"
	"time"
)

type noteReader interface {
	Next() bool
	ReadNote() (*time.Time, string, error)
}

func FormatWith(r noteReader, prefix string) string {
	d := dateFormat{}

	buf := bytes.Buffer{}
	for r.Next() {
		t, msg, err := r.ReadNote()
		if err != nil {
			continue
		}

		dateStr := d.format(t)

		buf.WriteString(prefix)
		buf.WriteString(dateStr)
		buf.WriteString(" ")
		buf.WriteString(msg)
		buf.WriteString("\n")
	}

	return buf.String()
}

func Format(r noteReader) string {
	return FormatWith(r, "")
}

type dateFormat struct {
	last *time.Time
}

func (d *dateFormat) format(t *time.Time) string {
	if d.last == nil {
		d.last = t
		return t.Format("2006-01-02 15:04:05")
	}
	if t.Year() == d.last.Year() && t.YearDay() == d.last.YearDay() {
		return t.Format("           15:04:05")
	}
	d.last = t
	return t.Format("2006-01-02 15:04:05")
}
