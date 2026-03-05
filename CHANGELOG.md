# Changelog

All notable changes to this project will be documented in this file.

## [v0.4.0](https://github.com/vianch/paco/releases/tag/v0.4.0)

- Moved package manager prompt to after script selection
- Show recommended package manager in prompt based on detected lock file
- Example: `How do you want to run this script? (recommended: pnpm)`

## [v0.3.0](https://github.com/vianch/paco/releases/tag/v0.3.0)

- Auto-detect lock file (`pnpm-lock.yaml`, `yarn.lock`, `package-lock.json`, `bun.lockb`) to pre-select package manager
- Package manager prompt with options: Direct, npm, yarn, pnpm, bun
- Added bun support

## [v0.2.0](https://github.com/vianch/paco/releases/tag/v0.2.0)

- Replaced custom list UI with [Bubbles Table](https://github.com/charmbracelet/bubbles) component
- Script and Command columns with scrollable, focused table
- Purple theme for borders, headers, and selected row highlight
- Added README with install instructions

## [v0.1.0](https://github.com/vianch/paco/releases/tag/v0.1.0)

- Initial release
- Interactive TUI script selection menu
- Auto-create `package.json` when missing
- Cross-platform builds (linux/darwin/windows x amd64/arm64)
- Homebrew tap published at https://github.com/vianch/homebrew-tap
