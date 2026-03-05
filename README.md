# Paco

A beautiful TUI tool that reads `package.json` scripts and lets you select and run them interactively. Built with Go, [BubbleTea](https://github.com/charmbracelet/bubbletea) and [Lipgloss](https://github.com/charmbracelet/lipgloss).

## Install

```bash
brew tap vianch/tap
brew install paco
```

Or with Go:

```bash
go install github.com/vianch/paco/cmd/paco@latest
```

## Usage

Run `paco` in any directory with a `package.json`:

```bash
paco
```

- Navigate with `j`/`k` or arrow keys
- Press `enter` to run the selected script
- Press `q` or `esc` to quit

If no `package.json` is found, paco will offer to create one with default scripts.

## Releases

### v0.2.0

- Replaced custom list UI with [Bubbles Table](https://github.com/charmbracelet/bubbles) component
- Script and Command columns with scrollable, focused table
- Purple theme for borders, headers, and selected row highlight
- Added README with install instructions
- Release: https://github.com/vianch/paco/releases/tag/v0.2.0

### v0.1.0

- Initial release
- Interactive TUI script selection menu
- Auto-create `package.json` when missing
- Cross-platform builds (linux/darwin/windows x amd64/arm64)
- Homebrew tap created at https://github.com/vianch/homebrew-tap
- Release: https://github.com/vianch/paco/releases/tag/v0.1.0

## License

MIT
