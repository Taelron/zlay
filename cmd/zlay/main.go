// zlay lists the Zellij layouts in ~/.config/zellij/layouts and
// replaces itself with `zellij --layout <name>` for the one picked.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	tea "github.com/charmbracelet/bubbletea"

	"zlay/internal/layouts"
	"zlay/internal/ui/tui/picker"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "zlay:", err)
		os.Exit(1)
	}
}

func run() error {
	dir, err := layouts.Dir()
	if err != nil {
		return err
	}
	names, err := layouts.List(dir)
	if err != nil {
		return err
	}
	if len(names) == 0 {
		return fmt.Errorf("no .kdl layouts in %s", dir)
	}

	final, err := tea.NewProgram(picker.New(names), tea.WithAltScreen()).Run()
	if err != nil {
		return err
	}
	name := final.(picker.Model).Chosen()
	if name == "" {
		return nil
	}

	zellij, err := exec.LookPath("zellij")
	if err != nil {
		return err
	}
	// Exec, not a child process: Zellij takes over the terminal and
	// zlay is gone. It only returns on failure.
	return syscall.Exec(zellij, []string{"zellij", "--layout", name}, os.Environ())
}
