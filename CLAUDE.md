# Paco - Package Commands Runner

A TUI tool that reads `package.json` scripts and lets you select and run them interactively.

## Tech Stack

- **Language:** Go (1.24+)
- **TUI Framework:** [BubbleTea](https://github.com/charmbracelet/bubbletea) (Elm-architecture based TUI framework)
- **Styling:** [Lipgloss](https://github.com/charmbracelet/lipgloss) (CSS-like terminal styling)

## Project Structure

```
paco/
  cmd/paco/          Entry point (main.go)
  internal/
    tui/             TUI components (BubbleTea models, styles)
      model.go       Main script selection menu
      prompt.go      Yes/No confirmation prompt
      styles.go      Lipgloss style definitions
    pkg/             Package.json parsing and manipulation
      package.go     Parse, list scripts, create default package.json
    runner/          Script execution
      runner.go      Shell command runner (sh -c / cmd /C)
```

## Architecture

Follows BubbleTea's Elm architecture pattern:
- **Model** holds state (script list, cursor position, selection)
- **Update** handles key messages and state transitions
- **View** renders the current state to a string

Each TUI screen is a separate BubbleTea model implementing `tea.Model`.

## Build & Run

```bash
go build -o paco ./cmd/paco/   # build binary
./paco                          # run in a directory with package.json
go install ./cmd/paco/          # install to $GOPATH/bin
```

## Conventions

- Use `internal/` for all non-exported packages
- Keep TUI models in `internal/tui/`, one file per screen/component
- Keep styles centralized in `styles.go`
- Business logic (parsing, running) stays separate from TUI code
- Sort scripts alphabetically for consistent display
- Use Lipgloss for all terminal styling (no raw ANSI codes)
- Handle both Unix and Windows shells in runner

## Behavior

1. Look for `package.json` in current directory
2. If not found, prompt user to create one with default scripts
3. Parse scripts and display interactive selection menu
4. Run selected script via shell, inheriting stdin/stdout/stderr
