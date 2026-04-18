package formatter

// IndentStyle controls whether indentation uses tabs or spaces.
type IndentStyle int

const (
	IndentTabs   IndentStyle = iota
	IndentSpaces             // uses IndentWidth spaces per level
)

// Config holds all formatting options for the printer.
type Config struct {
	IndentStyle   IndentStyle
	IndentWidth   int // only used when IndentStyle == IndentSpaces
	MaxBlankLines int // max consecutive blank lines allowed (inside blocks)
	PrintWidth    int // max line length before if/while conditions are wrapped
}

// DefaultConfig returns the default formatting config: tabs, 79-char limit.
func DefaultConfig() *Config {
	return &Config{
		IndentStyle:   IndentTabs,
		IndentWidth:   4,
		MaxBlankLines: 1,
		PrintWidth:    79,
	}
}
