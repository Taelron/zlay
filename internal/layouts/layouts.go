// Package layouts finds the Zellij layouts a user can start.
package layouts

import (
	"os"
	"path/filepath"
	"strings"
)

const ext = ".kdl"

// Dir returns the Zellij layout directory: ~/.config/zellij/layouts.
func Dir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "zellij", "layouts"), nil
}

// List returns the layout names in dir, sorted, without the .kdl
// extension. Only regular files ending exactly in .kdl count, so
// backups like foo.kdl.bak and directories are skipped.
func List(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var names []string
	for _, e := range entries {
		if !e.Type().IsRegular() || !strings.HasSuffix(e.Name(), ext) {
			continue
		}
		name := strings.TrimSuffix(e.Name(), ext)
		if name == "" {
			continue
		}
		names = append(names, name)
	}
	return names, nil
}
