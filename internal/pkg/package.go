package pkg

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

type PackageJSON struct {
	Name    string            `json:"name"`
	Version string            `json:"version"`
	Scripts map[string]string `json:"scripts"`
}

type Script struct {
	Name    string
	Command string
}

func FindPackageJSON() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path := filepath.Join(dir, "package.json")
	if _, err := os.Stat(path); err != nil {
		return "", fmt.Errorf("no package.json found in %s", dir)
	}
	return path, nil
}

func ParsePackageJSON(path string) (*PackageJSON, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var pkg PackageJSON
	if err := json.Unmarshal(data, &pkg); err != nil {
		return nil, fmt.Errorf("failed to parse package.json: %w", err)
	}
	return &pkg, nil
}

func (p *PackageJSON) ScriptList() []Script {
	scripts := make([]Script, 0, len(p.Scripts))
	for name, cmd := range p.Scripts {
		scripts = append(scripts, Script{Name: name, Command: cmd})
	}
	sort.Slice(scripts, func(i, j int) bool {
		return scripts[i].Name < scripts[j].Name
	})
	return scripts
}

func CreateDefaultPackageJSON(dir string) (string, error) {
	pkg := PackageJSON{
		Name:    filepath.Base(dir),
		Version: "1.0.0",
		Scripts: map[string]string{
			"start": "echo \"start script\"",
			"build": "echo \"build script\"",
			"test":  "echo \"test script\"",
			"lint":  "echo \"lint script\"",
			"dev":   "echo \"dev script\"",
		},
	}
	data, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		return "", err
	}
	path := filepath.Join(dir, "package.json")
	if err := os.WriteFile(path, append(data, '\n'), 0644); err != nil {
		return "", err
	}
	return path, nil
}
