# go-fyne-markdown-editor

A Markdown editor built with [Fyne](https://fyne.io/) and Go. Features a split-pane interface with live Markdown preview.

## Features

- **Live Preview** — Real-time Markdown rendering as you type
- **Split View** — Side-by-side editor and preview panes
- **File Operations** — Open, Save, and Save As for `.md` files
- **Cross-Platform** — Runs on macOS, Windows, and Linux
- **Native Look** — Uses system-native widgets via Fyne

## Prerequisites

- **Go 1.22+** — [Install Go](https://go.dev/doc/install)
- **C Compiler** — Required for Fyne (Xcode on macOS, MinGW on Windows, gcc on Linux)
- **Icon** — Place `Icon.png` or `icon.png` (512×512 recommended) in the project root for packaging

## Quick Start

```bash
# Clone the repository
git clone https://github.com/adarshsrinivasan/go-fyne-markdown-editor.git
cd go-fyne-markdown-editor

# Download dependencies
make deps

# Build and run
make run
```

## Build Commands

| Command | Description |
|---------|-------------|
| `make build` | Compile the application binary |
| `make run` | Build and run the application |
| `make clean` | Remove build artifacts and generated files |
| `make deps` | Download and verify dependencies |
| `make test` | Run tests |
| `make test-cover` | Run tests with coverage report |
| `make fmt` | Format source code |
| `make vet` | Run static analysis |
| `make lint` | Run golangci-lint (requires separate install) |
| `make package` | Create distributable `.app` package (requires fyne CLI) |
| `make help` | Show all available targets |

## Packaging

Create a release-ready macOS app bundle with `Icon.png` in the project root:

```bash
make package
open MarkDown.app
```

Requires the [fyne CLI](https://docs.fyne.io/started/) to be installed:

```bash
go install fyne.io/fyne/v2/cmd/fyne@latest
```

## Development

### Project Structure

```
.
├── main.go          # Application entry point and UI logic
├── Icon.png         # Application icon (512×512 PNG)
├── go.mod           # Go module definition
├── Makefile         # Build automation
└── README.md        # This file
```

### Running Tests

```bash
make test
# Or with coverage:
make test-cover
```

### Code Quality

```bash
make fmt    # Format code
make vet    # Run go vet
make lint   # Run golangci-lint (install separately)
```

## Requirements

- **Fyne v2** — GUI framework
- **Goldmark** — Markdown parsing (via Fyne's RichText)

## License

See [LICENSE](LICENSE) if present.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Run `make fmt test` before committing
4. Submit a pull request
