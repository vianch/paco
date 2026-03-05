package tui

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/vianch/paco/internal/pkg"
)

type Model struct {
	table    table.Model
	scripts  []pkg.Script
	selected *pkg.Script
	quitting bool
}

func NewModel(scripts []pkg.Script) Model {
	columns := []table.Column{
		{Title: "Script", Width: 20},
		{Title: "Command", Width: 50},
	}

	rows := make([]table.Row, len(scripts))
	for i, s := range scripts {
		rows[i] = table.Row{s.Name, s.Command}
	}

	height := len(scripts)
	if height > 15 {
		height = 15
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(height),
	)
	t.SetStyles(TableStyles())

	return Model{
		table:   t,
		scripts: scripts,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			idx := m.table.Cursor()
			if idx >= 0 && idx < len(m.scripts) {
				m.selected = &m.scripts[idx]
			}
			return m, tea.Quit
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.quitting {
		return ""
	}

	s := "\n" + titleStyle.Render("  paco - Package Commands Runner") + "\n\n"
	s += baseStyle.Render(m.table.View()) + "\n"
	s += helpStyle.Render("  j/k or arrows to navigate | enter to run | q to quit") + "\n"
	return s
}

func (m Model) SelectedScript() *pkg.Script {
	return m.selected
}
