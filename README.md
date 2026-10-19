# zlay

A small terminal picker for Zellij layouts.

zlay lists the `.kdl` files in `~/.config/zellij/layouts/` in a centred
modal. Pick one and zlay replaces itself with `zellij --layout <name>`, so
Zellij owns the terminal and zlay is gone.

Files that do not end exactly in `.kdl` (such as `foo.kdl.bak`) are ignored.

## Keys

| Key            | Action             |
|----------------|--------------------|
| `↑` / `↓`      | Move               |
| `Enter`        | Start the layout   |
| `q` / `Esc`    | Quit without start |

## Build and install

Requires Go 1.25 and `zellij` on `PATH`.

```sh
make test      # run the tests
make build     # build ./zlay
make install   # build and copy to ~/.local/bin/zlay
```

## Layout

```
cmd/zlay/              entry point: list, pick, exec zellij
internal/layouts/      finds the .kdl layouts
internal/ui/tui/       Bubble Tea picker and Lipgloss styles
```

The styles reuse the taelron-tui-base palette. zlay is a personal tool,
not a Taelron product.
