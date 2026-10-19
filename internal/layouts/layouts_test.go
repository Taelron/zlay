package layouts

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestListKeepsOnlyKdlFiles(t *testing.T) {
	dir := t.TempDir()
	for _, f := range []string{
		"mkx.kdl", "fleetdesk.kdl", "fleetdesk.kdl.bak",
		"topicchime.kdl.bak-20261015", "notes.txt", ".kdl",
	} {
		if err := os.WriteFile(filepath.Join(dir, f), nil, 0o644); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.Mkdir(filepath.Join(dir, "sub.kdl"), 0o755); err != nil {
		t.Fatal(err)
	}

	got, err := List(dir)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"fleetdesk", "mkx"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("List() = %v, want %v", got, want)
	}
}

func TestListMissingDir(t *testing.T) {
	if _, err := List(filepath.Join(t.TempDir(), "absent")); err == nil {
		t.Fatal("List() on a missing dir returned no error")
	}
}
