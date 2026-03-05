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

## Release v0.1.0

1. Homebrew tap repo created - https://github.com/vianch/homebrew-tap
2. Code pushed to `origin/main`
3. Tag `v0.1.0` created and pushed
4. GoReleaser built 6 binaries (linux/darwin/windows x amd64/arm64)
5. GitHub release published - https://github.com/vianch/paco/releases/tag/v0.1.0
6. Homebrew formula pushed to `vianch/homebrew-tap`

## License

MIT
