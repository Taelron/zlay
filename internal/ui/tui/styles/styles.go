// Package styles holds zlay's Lipgloss styles. The palette is copied
// from taelron-tui-base so zlay looks like the other Taelron TUIs.
package styles

import "github.com/charmbracelet/lipgloss"

// Colour palette, ANSI 256-colour codes as in taelron-tui-base.
var (
	ColorInfo      = lipgloss.Color("14")  // cyan
	ColorBorder    = lipgloss.Color("12")  // blue
	ColorHint      = lipgloss.Color("241") // dim grey
	ColorEmphasis  = lipgloss.Color("15")  // bright white
	ColorRule      = lipgloss.Color("238") // dark grey
	ColorHighlight = lipgloss.Color("237") // selected row background
)

// MinContentWidth keeps the modal from shrinking below the hint bar.
const MinContentWidth = 36

// Modal is the centred bordered container.
var Modal = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(ColorBorder).
	Padding(1, 2)

// Title labels the modal.
var Title = lipgloss.NewStyle().
	Foreground(ColorInfo).
	Bold(true)

// Rule separates the title and hint bar from the list.
var Rule = lipgloss.NewStyle().
	Foreground(ColorRule)

// Row is an unselected list row.
var Row = lipgloss.NewStyle().
	PaddingLeft(1)

// Selected is the row under the cursor: a full-width highlight bar.
var Selected = lipgloss.NewStyle().
	Foreground(ColorEmphasis).
	Background(ColorHighlight).
	Bold(true).
	PaddingLeft(1)

// FocusedMarker is the glyph in front of the selected row.
var FocusedMarker = lipgloss.NewStyle().
	Foreground(ColorInfo).
	Background(ColorHighlight).
	Bold(true)

// HintBar renders the keymap reminders at the bottom of the modal.
var HintBar = lipgloss.NewStyle().
	Foreground(ColorHint)
