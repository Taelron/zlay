// Package picker is the layout list: move with up/down, Enter picks,
// q or Esc quits without picking.
package picker

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"zlay/internal/ui/tui/styles"
)

const hints = "↑/↓ move · enter start · q/esc quit"

// Model lists layout names and records the one the user picks.
type Model struct {
	names  []string
	cursor int
	chosen string
	width  int
	height int
}

// New returns a picker over names.
func New(names []string) Model {
	return Model{names: names}
}

// Chosen is the picked layout name, or "" if the user quit.
func (m Model) Chosen() string {
	return m.chosen
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		return m, nil
	case tea.KeyMsg:
		return m.handleKey(msg)
	}
	return m, nil
}

func (m Model) handleKey(key tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch key.String() {
	case "up":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down":
		if m.cursor < len(m.names)-1 {
			m.cursor++
		}
	case "enter":
		if len(m.names) > 0 {
			m.chosen = m.names[m.cursor]
		}
		return m, tea.Quit
	case "q", "esc", "ctrl+c":
		return m, tea.Quit
	}
	return m, nil
}

// View renders the modal, centred in the terminal once its size is known.
func (m Model) View() string {
	width := lipgloss.Width(hints)
	for _, name := range m.names {
		width = max(width, lipgloss.Width(name)+3)
	}
	width = max(width, styles.MinContentWidth)

	rule := styles.Rule.Render(strings.Repeat("─", width))
	rows := []string{styles.Title.Render("Zellij layouts"), rule}
	for i, name := range m.names {
		if i == m.cursor {
			row := styles.FocusedMarker.Render("▸") + styles.Selected.Render(name)
			rows = append(rows, row+styles.Selected.Render(strings.Repeat(" ", width-lipgloss.Width(row)-1)))
		} else {
			rows = append(rows, styles.Row.Render(" "+name))
		}
	}
	rows = append(rows, rule, styles.HintBar.Render(hints))

	modal := styles.Modal.Render(lipgloss.JoinVertical(lipgloss.Left, rows...))
	if m.width == 0 || m.height == 0 {
		return modal
	}
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, modal)
}
