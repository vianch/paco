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

## Upgrade

If you already have paco installed via Homebrew, upgrade to the latest version:

```bash
brew update
brew upgrade vianch/tap/paco
```

> **Note:** Use the fully qualified name `vianch/tap/paco` to ensure Homebrew upgrades from the correct tap.

## Usage

Run `paco` in any directory with a `package.json`:

```bash
paco
```

- Navigate with `j`/`k` or arrow keys
- Press `enter` to run the selected script
- Press `q` or `esc` to quit

If no `package.json` is found, paco will offer to create one with default scripts.

## License

MIT
