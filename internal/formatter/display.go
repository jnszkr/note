package formatter

import (
	"bytes"
	"time"

	"github.com/jnszkr/note/internal/reader"
)

func Format(r *reader.NoteReader) string {
	d := dateFormat{}

	buf := bytes.Buffer{}
	for r.Next() {
		t, msg, err := r.Read()
		if err != nil {
			continue
		}

		dateStr := d.format(t)

		buf.WriteString(dateStr)
		buf.WriteString(" ")
		buf.WriteString(msg)
		buf.WriteString("\n")
	}

	return buf.String()
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
