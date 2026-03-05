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

	pkgManager := promptPackageManager()
	if pkgManager == "" {
		os.Exit(0)
	}

	command := buildCommand(pkgManager, selected.Name, selected.Command)
	fmt.Printf("Running: %s\n", command)
	if err := runner.RunScript(command); err != nil {
		fmt.Fprintf(os.Stderr, "Script failed: %v\n", err)
		os.Exit(1)
	}
}

func buildCommand(manager, scriptName, rawCommand string) string {
	switch manager {
	case "npm":
		return "npm run " + scriptName
	case "yarn":
		return "yarn " + scriptName
	case "pnpm":
		return "pnpm " + scriptName
	case "bun":
		return "bun " + scriptName
	default:
		return rawCommand
	}
}

func promptPackageManager() string {
	detected := pkg.DetectPackageManager()
	options := []string{"Direct", "npm", "yarn", "pnpm", "bun"}

	defaultOption := "Direct"
	question := "How do you want to run this script?"
	if detected != "" {
		defaultOption = detected
		question = fmt.Sprintf("How do you want to run this script? (recommended: %s)", detected)
	}

	prompt := tui.NewPromptModelWithDefault(
		question,
		options,
		defaultOption,
	)
	p := tea.NewProgram(prompt)
	finalModel, err := p.Run()
	if err != nil {
		return ""
	}
	return finalModel.(tui.PromptModel).Chosen()
}

func promptCreatePackageJSON() bool {
	prompt := tui.NewPromptModel(
		"No package.json found. Create one with default scripts?",
		[]string{"Yes", "No"},
	)
	p := tea.NewProgram(prompt)
	finalModel, err := p.Run()
	if err != nil {
		return false
	}
	return finalModel.(tui.PromptModel).Accepted()
}
