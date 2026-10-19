package picker

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func press(m Model, keys ...tea.KeyMsg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	for _, k := range keys {
		var next tea.Model
		next, cmd = m.Update(k)
		m = next.(Model)
	}
	return m, cmd
}

var (
	up    = tea.KeyMsg{Type: tea.KeyUp}
	down  = tea.KeyMsg{Type: tea.KeyDown}
	enter = tea.KeyMsg{Type: tea.KeyEnter}
	esc   = tea.KeyMsg{Type: tea.KeyEsc}
	q     = tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("q")}
)

func isQuit(cmd tea.Cmd) bool {
	if cmd == nil {
		return false
	}
	_, ok := cmd().(tea.QuitMsg)
	return ok
}

func TestEnterPicksRowUnderCursor(t *testing.T) {
	m, cmd := press(New([]string{"a", "b", "c"}), down, down, up, enter)
	if m.Chosen() != "b" {
		t.Fatalf("Chosen() = %q, want %q", m.Chosen(), "b")
	}
	if !isQuit(cmd) {
		t.Fatal("enter did not quit")
	}
}

func TestCursorStaysInBounds(t *testing.T) {
	m, _ := press(New([]string{"a", "b"}), up, enter)
	if m.Chosen() != "a" {
		t.Fatalf("after up at top: Chosen() = %q, want %q", m.Chosen(), "a")
	}
	m, _ = press(New([]string{"a", "b"}), down, down, down, enter)
	if m.Chosen() != "b" {
		t.Fatalf("after down past end: Chosen() = %q, want %q", m.Chosen(), "b")
	}
}

func TestQuitKeysPickNothing(t *testing.T) {
	for _, k := range []tea.KeyMsg{q, esc} {
		m, cmd := press(New([]string{"a"}), k)
		if m.Chosen() != "" {
			t.Fatalf("%s: Chosen() = %q, want empty", k, m.Chosen())
		}
		if !isQuit(cmd) {
			t.Fatalf("%s did not quit", k)
		}
	}
}

func TestOtherKeysDoNothing(t *testing.T) {
	_, cmd := press(New([]string{"a"}), tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("x")})
	if cmd != nil {
		t.Fatal("x returned a command")
	}
}

func TestViewFillsTheWindow(t *testing.T) {
	next, _ := New([]string{"a", "b"}).Update(tea.WindowSizeMsg{Width: 120, Height: 40})
	view := next.View()
	if got := lipgloss.Height(view); got != 40 {
		t.Fatalf("view height = %d, want 40", got)
	}
	if got := lipgloss.Width(view); got != 120 {
		t.Fatalf("view width = %d, want 120", got)
	}
}
