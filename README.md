# Planni

[![Go Report Card](https://goreportcard.com/badge/github.com/TreGalloway/planni)](https://goreportcard.com/report/github.com/TreGalloway/planni)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://img.shields.io/github/actions/workflow/status/TreGalloway/planni/go.yml?branch=main)](https://github.com/TreGalloway/planni/actions)
[![Latest Release](https://img.shields.io/github/v/release/TreGalloway/planni)](https://github.com/TreGalloway/planni/releases)

> **🚧 This Project is a Work in Progress! 🚧**
> **🚧 This Project is a Work in Progress! 🚧**

`Planni` is a modern, all-in-one CLI designed to be the backbone of your terminal. It reimagines your daily workflow with smart directory navigation, a beautiful and feature-rich file lister, an interactive file browser, and powerful utilities, all within a single, blazingly fast Go binary.

The default terminal commands are powerful, but they haven't fundamentally changed in decades. `planni` is an attempt to build what the terminal experience *could* be—more productive, more intuitive, and visually rich.

<!-- TODO: Record a GIF demonstrating the 'p j ...' -> 'p' workflow and add it here! -->
![demo_gif_placeholder](https://user-images.githubusercontent.com/12345/placeholder.gif)

---

## ✨ Core Features

*   🧠 **Smart Jump (`zoxide` like):** A `jump` command that learns your habits. `p j api` takes you straight to your most frequently used API project, no matter where you are.
*   🚀 **Modern List (`eza` like):** A beautiful, context-aware `ls`. Get file-type icons, Git status integration, a tree view, and more, all by default.
*   📁 **Interactive Browser (`spf` like):** A planned TUI file manager for when you need to do more. Copy, move, delete, and browse your files with a fluid, keyboard-driven interface.
*   🔒 **Secure Backups:** A planned `backup` command that uses SFTP to securely copy files or directories to any remote server you have SSH access to.

## 📦 Installation

**Prerequisites:**
*   Go 1.21+
*   A [Nerd Font](https://www.nerdfonts.com/) installed and enabled in your terminal for icon support.

#### 1. Install the binary

```bash
go install github.com/your-username/planni@latest
```

#### 2. Set up the shell function

To enable the `p` command and the smart `jump` functionality, you **must** add a function to your shell's configuration file (`~/.zshrc`, `~/.bashrc`, etc.).

```bash
# Add this to your ~/.zshrc or ~/.bashrc

p() {
    # The hook automatically logs the current directory for the 'jump' logic.
    # We run it in the background and redirect output to /dev/null to keep the prompt clean.
    command planni hook --add "$(pwd)" &> /dev/null &

    case "$1" in
        ""|ls|l|lt)
            # If 'p' is called with no args, or 'ls'/'l'/'lt', run the list command.
            command planni ls "${@:1}" 
            ;;
        j|jump)
            # The jump command requires changing the shell's directory.
            local dest
            dest="$(command planni _jump_internal "${@:2}")"
            if [[ -n "$dest" ]]; then
                cd "$dest"
            fi
            ;;
        *)
            # Handle all other subcommands like 'p browse', 'p backup'.
            command planni "$@"
            ;;
    esac
}
```

Restart your shell or source the file (`source ~/.zshrc`) to apply the changes.

## 🗺️ Roadmap

- [ ] **Phase 1: The Core Workflow**
    - [ ] Implement Smart Jump (`p j`) logic with frecency scoring.
    - [ ] Implement Modern List (`p ls`) with icons, colors, and a tree view.
- [ ] **Phase 2: The Interactive Layer**
    - [ ] Build the TUI File Browser (`p browse`) with `bubbletea`.
    - [ ] Create a syntax-highlighted file viewer (`p cat`).
- [ ] **Phase 3: Networking Utilities**
    - [ ] Implement the secure SFTP `backup` command.
    - [ ] Build the interactive SSH connection menu.

## 🔭 Future Vision

Once the core features are mature and stable, we are interested in exploring how modern technologies can further enhance terminal productivity. This could include ideas like integrating local, context-aware AI to assist with complex commands or summarizing logs. These are long-term ideas, not committed features.

## 🙌 Contributing

Contributions are welcome! If you're interested in helping build the future of the terminal, please see `CONTRIBUTING.md` for guidelines on how to get started.

## 📜 License

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for details.
