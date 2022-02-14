package notes

import (
	"bytes"
	"fmt"
	"regexp"
	"runtime"
	"time"
)

const (
	space   = ' '
	newLine = '\n'

	DefaultPrimaryDateFormat   = "2006-01-02 15:04:05"
	DefaultSecondaryDateFormat = "           15:04:05"
)

type Option func(*formattedNoteReader)

func WithPrefix(p string) Option {
	return func(ops *formattedNoteReader) {
		ops.prefix = p
	}
}

func WithHighlight(substr string, color Color) Option {
	return func(ops *formattedNoteReader) {
		ops.substr = substr
		ops.color = color
	}
}

type FormatterFunc func(NoteReader) string

func Formatter(opts ...Option) FormatterFunc {
	f := &formattedNoteReader{
		primaryDateFormat:   DefaultPrimaryDateFormat,
		secondaryDateFormat: DefaultSecondaryDateFormat,
	}

	for _, opt := range opts {
		opt(f)
	}

	return f.format
}

type formattedNoteReader struct {
	prefix              string
	primaryDateFormat   string
	secondaryDateFormat string
	substr              string
	color               Color
}

func (f *formattedNoteReader) format(r NoteReader) string {
	d := dateFormat{
		primary:   f.primaryDateFormat,
		secondary: f.secondaryDateFormat,
	}

	buf := bytes.Buffer{}
	for r.Next() {
		note := r.ReadNote()
		if len(f.substr) > 0 {
			note.Text = Highlight(f.substr, f.color, note.Text)
		}

		dateStr := d.formatTime(&note.Created)

		buf.WriteString(f.prefix)
		buf.WriteString(dateStr)
		buf.WriteByte(space)
		buf.WriteString(note.Text)
		buf.WriteByte(newLine)
	}

	return buf.String()
}

type dateFormat struct {
	prev      *time.Time
	primary   string
	secondary string
}

func (d *dateFormat) isSecondary(curr *time.Time) bool {
	return curr.Year() == d.prev.Year() && curr.YearDay() == d.prev.YearDay()
}

func (d *dateFormat) formatTime(t *time.Time) string {
	if d.prev == nil {
		d.prev = t
		return t.Format(d.primary)
	}
	if d.isSecondary(t) {
		return t.Format(d.secondary)
	}
	d.prev = t
	return t.Format(d.primary)
}

type Color string

var (
	ResetColor  Color = "\033[0m"
	RedColor    Color = "\033[31m"
	GreenColor  Color = "\033[32m"
	YellowColor Color = "\033[33m"
	BlueColor   Color = "\033[34m"
	PurpleColor Color = "\033[35m"
	CyanColor   Color = "\033[36m"
	GrayColor   Color = "\033[37m"
	WhiteColor  Color = "\033[97m"
)

var winos = false

func init() {
	winos = runtime.GOOS == "windows"
	if winos {
		ResetColor = ""
		RedColor = ""
		GreenColor = ""
		YellowColor = ""
		BlueColor = ""
		PurpleColor = ""
		CyanColor = ""
		GrayColor = ""
		WhiteColor = ""
	}
}

// Highlight would highlight the substring with the given color.
// When substring is empty or set to *, the full text is highlighted.
func Highlight(substr string, c Color, text string) string {
	if winos {
		return text
	}
	if len(substr) == 0 || substr == "*" {
		return fmt.Sprintf("%s%s%s", c, text, ResetColor)
	}
	re := regexp.MustCompile(fmt.Sprintf("(?i)(%s)", substr))
	return re.ReplaceAllString(text, fmt.Sprintf("%s$1%s", c, ResetColor))
}
