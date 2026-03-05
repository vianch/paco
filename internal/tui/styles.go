package tui

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#9B59B6")).
			MarginBottom(1)

	baseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#9B59B6")).
			Padding(0, 1)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			MarginTop(1)

	promptSelectedStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#04B575")).
				Bold(true)

	promptNormalStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#626262"))
)

func TableStyles() table.Styles {
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#9B59B6")).
		BorderBottom(true).
		Bold(true).
		Foreground(lipgloss.Color("#9B59B6"))
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#FFF")).
		Background(lipgloss.Color("#9B59B6")).
		Bold(true)
	return s
}
