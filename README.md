# Backhome

This cli tool serve two purposes: learning golang and a simple backup tool for my configuration files.

It reads a `.config.yaml` (default to `$HOME/.config.yaml`), copy `targets` files to `local` git rempository, commit the changes and push to a `remote` git repository.

I intend to add:
 - [ ] `sync` command (commit changes and push to remote)
 - [ ] `restore` (copy files from local to their origin)
 - [ ] `setup`  (create basic config file, init local repository, add remote as its origin and pull the most recent version)
 - [ ] `add` and `remove` target files from beign watched
 - [ ] run `copy` and `sync` periodically 
 - [ ] Alerts using macos notification
 - [ ] Support to multiple environments (work and home)
  
## Usage
TDB
