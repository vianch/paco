package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/vianch/paco/internal/pkg"
	"github.com/vianch/paco/internal/runner"
	"github.com/vianch/paco/internal/tui"
)

func main() {
	path, err := pkg.FindPackageJSON()
	if err != nil {
		if !promptCreatePackageJSON() {
			fmt.Println("Goodbye!")
			os.Exit(0)
		}
		dir, _ := os.Getwd()
		path, err = pkg.CreateDefaultPackageJSON(dir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating package.json: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Created package.json!")
	}

	pkgJSON, err := pkg.ParsePackageJSON(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	scripts := pkgJSON.ScriptList()
	if len(scripts) == 0 {
		fmt.Println("No scripts found in package.json")
		os.Exit(0)
	}

	model := tui.NewModel(scripts)
	p := tea.NewProgram(model)
	finalModel, err := p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	m := finalModel.(tui.Model)
	selected := m.SelectedScript()
	if selected == nil {
		os.Exit(0)
	}

	fmt.Printf("Running: %s -> %s\n", selected.Name, selected.Command)
	if err := runner.RunScript(selected.Command); err != nil {
		fmt.Fprintf(os.Stderr, "Script failed: %v\n", err)
		os.Exit(1)
	}
}

func promptCreatePackageJSON() bool {
	prompt := tui.NewPromptModel("No package.json found. Create one with default scripts?")
	p := tea.NewProgram(prompt)
	finalModel, err := p.Run()
	if err != nil {
		return false
	}
	return finalModel.(tui.PromptModel).Accepted()
}
