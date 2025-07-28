# tmpl

A CLI tool for quickly creating project boilerplate code from Git repository templates.
tmpl allows you to manage multiple template repositories and generate new projects
with customizable addons.

This project was started with the boot.dev Hackathon 2025.

## The Problem

At work, I noticed we were constantly starting new projects with remarkably similar
code structures. While GitHub templates seemed like an obvious solution, the reality
was more complex - everyone's needs were slightly different. One project needed a
login screen and database integration, another required database connectivity plus
header and footer components, and yet another had its own unique combination of
requirements.

Sure, these additions could be done manually after using a basic template, but this
approach led to inconsistencies and errors, especially among less experienced team
members who struggled with the manual setup process. Junior developers would forget
to configure certain components properly, miss important boilerplate code, or
introduce bugs while trying to integrate multiple pieces together.

What we needed was a system that could handle the composable nature of modern
projects - something that could start with a solid base template and then
intelligently layer on exactly the components each project required, without the
error-prone manual work.

## Features

- Add and manage Git repositories containing project templates
- Create new projects from available templates
- Support for template addons to extend base templates
- Update template repositories to get latest changes
- Simple interactive interface for template selection

## Installation

### From `go install`

```bash
go install https://github.com/IC3P3/tmpl@latest
```

### From Source

Ensure you have Go installed.

```bash
git clone https://github.com/IC3P3/tmpl.git
cd tmpl

go build -o tmpl .
```

Move the binary to a location in your PATH:

```bash
sudo mv tmpl /usr/local/bin/
```

Or for local installation:

```bash
mkdir -p ~/.local/bin
mv tmpl ~/.local/bin/
# Ensure ~/.local/bin is in your PATH
```

## Usage

### Adding Template Repositories

Add a new Git repository containing templates:

```bash
tmpl add <name> <repository-url>
```

Example:

```bash
tmpl add my-templates https://github.com/IC3P3/tmpl-example.git
```

### Listing Available Repositories

View all added template repositories:

```bash
tmpl list
```

### Creating a New Project

Start the interactive project creation process:

```bash
tmpl create
```

This will:

1. Prompt for a project name
2. Show available template repositories
3. Display templates within the selected repository
4. Allow selection of addons for the chosen template
5. Generate the project in the current directory with the project name

### Updating Repositories

Update all template repositories to get the latest changes:

```bash
tmpl update
```

### Removing Repositories

Remove a template repository:

```bash
tmpl remove <name>
```

### Other Commands

- `tmpl version` - Show version information (coming later)
- `tmpl check` - Check template repository syntax (coming later)
- `tmpl init` - Initialize a new template repository (coming later)
- `tmpl addon` - Add addons to existing projects (coming later)
- `tmpl upload` - Upload template repository (coming later)

## Template Repository Structure

Template repositories should follow this structure:

```txt
repository/
├── tmpl.json           # Repository configuration
├── templates/          # Base templates
│   ├── template1/
│   └── template2/
└── addons/            # Optional addons
    ├── addon1/
    └── addon2/
```

### Configuration Format

The `tmpl.json` file should contain:

```json
{
  "displayname": "My Templates",
  "description": "A collection of project templates",
  "templates": [
    {
      "displayname": "Basic Template",
      "name": "basic",
      "description": "A basic project template",
      "addons": ["addon1", "addon2"]
    }
  ],
  "addons": [
    {
      "displayName": "Additional Feature",
      "name": "addon1",
      "description": "Adds extra functionality"
    }
  ],
  "hooks": []
}
```

## Building

### Prerequisites

- Go
- Git (for cloning template repositories)

### Build Commands

```bash
# Build for current platform
go build -o tmpl .

# Cross-compile for different platforms
GOOS=linux GOARCH=amd64 go build -o tmpl-linux-amd64 .
GOOS=windows GOARCH=amd64 go build -o tmpl-windows-amd64.exe .
GOOS=darwin GOARCH=amd64 go build -o tmpl-darwin-amd64 .
```

### Dependencies

The project uses the following main dependencies:

- `github.com/spf13/cobra` - CLI framework
- `github.com/go-git/go-git/v5` - Git operations

## Data Storage

tmpl stores its data in `~/.local/share/tmpl/`:

- `repositories.json` - List of added repositories
- `<repo-name>/` - Cloned template repositories

## Troubleshooting

### Common Issues

#### Permission denied when creating projects

- Ensure you have write permissions in the target directory

#### Repository clone fails

- Verify the repository URL is correct and accessible

#### tmpl: Command not found

- If installed using `go install` ensure `GOPATH` is part of the `PATH`

## TODO List - Post-Hackathon Development

- [ ] **`check` command** - Template repository validation
- [ ] **`upload` command** - Template repository publishing
- [ ] **`version` command** - Version of the tool bound to CI/CD
- [ ] **Authentication system** - Being able to access private repositories
- [ ] **Repository configuration enhancements** - Set global or repository-based authentication
- [ ] **Pre-/Post-creation hooks** - Select and execute script
- [ ] **File content modification** - Let templates change files instead of just copying
- [ ] **GitHub Actions workflows** - Testing and building of new versions
- [ ] **Release automation** - Create releases on SemVer tags
- [ ] **Refacor and code cleanup** - Improve code written by a Go beginner (me) in the hectic of a Hackathon
