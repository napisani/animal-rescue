## animal-rescue

This is a simple utility for curating `pet` snippets from various sources including:
1. local pet-snippet.toml files (based on the cwd)
2. package.json files (pnpm / npm supported)
3. Makefiles

### build 
```bash
make
```

### run
```bash
./animal-rescue --config [path/to/pet/config/toml] [--clean] [--debug]
```



### bash_profile entry
```bash
function pet-select() {
  # modify this part of the search command to use `animal-rescue` to get the config file
  # reference the pet documentation for more information: https://github.com/knqyf263/pet
  BUFFER=$(pet search --query "$READLINE_LINE" --config $(animal-rescue --config ~/.config/pet/config.toml))
  READLINE_LINE=$BUFFER
  READLINE_POINT=${#BUFFER}
}
bind -x '"\C-x\C-r": pet-select'
```

