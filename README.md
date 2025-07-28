# tmpl

A CLI tool for quickly creating project boilerplate code from Git repository templates.
tmpl allows you to manage multiple template repositories and generate new projects
with customizable addons.

This project was started with the boot.dev Hackathon 2025.

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

### Missing Commands Implementation

- [ ] **`check` command** - Template repository validation
- [ ] **`upload` command** - Template repository publishing
- [ ] **`version` command** - Version of the tool bound to CI/CD

### Private Repository Support

- [ ] **Authentication system**
  - [ ] SSH key support for private Git repositories
  - [ ] Personal Access Token (PAT) authentication
  - [ ] Git credential helper integration
  - [ ] Secure credential storage (keyring/keychain)
- [ ] **Repository access management**
  - [ ] Support for GitHub/GitLab/Bitbucket private repos
  - [ ] Organization/team repository access
  - [ ] Repository URL validation for private repos
  - [ ] Error handling for authentication failures
- [ ] **Configuration enhancements**
  - [ ] Per-repository authentication settings
  - [ ] Global authentication configuration
  - [ ] Repository-specific access tokens

### Hooks System Implementation

- [ ] **Pre-creation hooks**
  - [ ] Environment validation
  - [ ] Dependency checking
  - [ ] Custom setup scripts
  - [ ] Interactive prompts and confirmations
- [ ] **Post-creation hooks**
  - [ ] Package installation (npm install, go mod tidy)
  - [ ] Git repository initialization
  - [ ] Initial commits and setup
  - [ ] Development environment setup
  - [ ] Custom build processes
- [ ] **Hook execution engine**
  - [ ] Cross-platform script execution
  - [ ] Environment variable passing
  - [ ] Error handling and rollback
  - [ ] Hook timeout management
  - [ ] Logging and debugging output

### Advanced Templating Features

- [ ] **File content modification**
  - [ ] Insert code blocks at specific markers
  - [ ] Replace/modify existing content
  - [ ] Merge configuration files (JSON, YAML, TOML)
  - [ ] Smart import statement insertion
- [ ] **Template inheritance and composition**
  - [ ] Base template extension
  - [ ] Partial template inclusion
  - [ ] Template layering system
  - [ ] Conflict resolution strategies
- [ ] **Advanced variable handling**
  - [ ] Computed variables (derived from other variables)
  - [ ] Variable validation with custom rules
  - [ ] Interactive variable input with prompts
  - [ ] Variable templates and presets
- [ ] **File operation enhancements**
  - [ ] Conditional file inclusion/exclusion
  - [ ] Dynamic file renaming patterns
  - [ ] Binary file handling
  - [ ] Symlink creation and management

### CI/CD Pipeline

- [ ] **GitHub Actions workflows**
  - [ ] Automated testing on multiple Go versions
  - [ ] Cross-platform builds (Linux, macOS, Windows)
  - [ ] Integration tests with real templates
  - [ ] Linting and code quality checks
- [ ] **Release automation**
  - [ ] Semantic versioning with conventional commits
  - [ ] Automated changelog generation
  - [ ] Binary artifact creation and upload
  - [ ] Package manager integration (Homebrew, Chocolatey, etc.)
  - [ ] Docker image creation and publishing
- [ ] **Quality assurance**
  - [ ] Unit test coverage reporting
  - [ ] Security vulnerability scanning
  - [ ] Dependency update automation
  - [ ] Performance benchmarking

### Code Cleanup and Refactoring

- [ ] **Architecture improvements**
  - [ ] Separate concerns into distinct packages
  - [ ] Implement proper dependency injection
  - [ ] Create interfaces for testability
  - [ ] Standardize error handling patterns
- [ ] **Code quality enhancements**
  - [ ] Comprehensive unit test coverage (>90%)
  - [ ] Integration test suite
  - [ ] Code documentation and comments
  - [ ] Go best practices compliance
  - [ ] Memory and performance optimization
- [ ] **Configuration management**
  - [ ] Centralized configuration handling
  - [ ] Environment-based configuration
  - [ ] Configuration validation and defaults
  - [ ] Migration system for config changes
- [ ] **User experience improvements**
  - [ ] Consistent command-line interface
  - [ ] Better error messages and help text
  - [ ] Progress indicators for long operations
  - [ ] Colored output and formatting
  - [ ] Interactive prompts and confirmations

### Documentation and Community

- [ ] **Documentation improvements**
  - [ ] Comprehensive user guide
  - [ ] Template author documentation
  - [ ] API documentation for extensibility
  - [ ] Tutorial videos and examples
- [ ] **Community features**
  - [ ] Template marketplace/registry
  - [ ] Community template sharing
  - [ ] Template rating and reviews
  - [ ] Contributing guidelines and templates

### Performance and Scalability

- [ ] **Optimization**
  - [ ] Parallel template processing
  - [ ] Incremental updates and caching
  - [ ] Memory usage optimization
  - [ ] Large repository handling
- [ ] **Monitoring and observability**
  - [ ] Usage analytics (opt-in)
  - [ ] Performance metrics
  - [ ] Error reporting and tracking
  - [ ] Debug mode and verbose logging

### Security Enhancements

- [ ] **Security features**
  - [ ] Template sandboxing
  - [ ] Malicious code detection
  - [ ] Dependency vulnerability scanning
  - [ ] Secure defaults and practices
- [ ] **Compliance**
  - [ ] License compatibility checking
  - [ ] GDPR compliance for any data collection
  - [ ] Security audit and penetration testing

### Priority Levels

**High Priority (Next Sprint)**:

- Missing commands (`check`, `upload`)
- Private repository support
- Basic hooks system
- CI/CD setup

**Medium Priority (Following Sprints)**:

- Advanced templating features
- Code refactoring and cleanup
- Performance optimizations

**Low Priority (Future Releases)**:

- Community features
- Advanced security features
- Marketplace integration
