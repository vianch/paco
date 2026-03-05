package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/victorchavarro/paco/internal/pkg"
)

type Model struct {
	scripts  []pkg.Script
	cursor   int
	selected *pkg.Script
	quitting bool
}

func NewModel(scripts []pkg.Script) Model {
	return Model{scripts: scripts}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			m.quitting = true
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.scripts)-1 {
				m.cursor++
			}
		case "enter":
			m.selected = &m.scripts[m.cursor]
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	if m.quitting {
		return ""
	}

	var b strings.Builder

	b.WriteString(titleStyle.Render("  paco") + "\n")
	b.WriteString(titleStyle.Render("  Package Commands Runner") + "\n\n")

	for i, s := range m.scripts {
		if m.cursor == i {
			line := fmt.Sprintf(" %s %s  %s",
				cursorIcon,
				selectedStyle.Render(s.Name),
				commandStyle.Render(s.Command),
			)
			b.WriteString(line + "\n")
		} else {
			line := fmt.Sprintf("   %s  %s",
				itemStyle.Render(s.Name),
				commandStyle.Render(s.Command),
			)
			b.WriteString(line + "\n")
		}
	}

	b.WriteString(helpStyle.Render("\n  j/k or arrows to navigate | enter to run | q to quit"))
	return b.String()
}

func (m Model) SelectedScript() *pkg.Script {
	return m.selected
}
