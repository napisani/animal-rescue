## animal-rescue

This is a simple utility for curating `pet` snippets from various sources including:
1. local pet-snippet.toml files (based on the cwd)
2. package.json files (pnpm / npm supported)
3. Makefiles
4. annotated shell aliases and functions (opt-in)

### build 
```bash
make
```

### run
```bash
./animal-rescue --config [path/to/pet/config/toml] [--clean] [--debug]
```

To scan shell aliases and functions, pass one or more `--shell-path` options.
A path may name a shell file or a directory of shell fragments. Only declarations
immediately following a `# pet: ...` annotation are included:

```bash
# pet: Rebuild this machine's Nix configuration
alias nixupgrade='...'

# pet: Run kubectl against the homelab cluster
labkubectl() {
  ...
}
```

The generated Pet command is the alias or function name; its implementation is
not copied into the snippet. Without `--shell-path`, shell files are not read.

### bash_profile entry
```bash
function pet-select() {
  BUFFER=$(pet search --query "$READLINE_LINE" --config $(animal-rescue \
    --config ~/.config/pet/config.toml \
    --shell-path "$SHELL_DOTFILES_DIR/bash/rc.d"))
  READLINE_LINE=$BUFFER
  READLINE_POINT=${#BUFFER}
}
bind -x '"\C-x\C-r": pet-select'
```
