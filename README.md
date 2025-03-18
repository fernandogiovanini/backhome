# Backhome

This cli tool serve two purposes: learning golang and a simple backup tool for my configuration files.

It reads a `.backhome.yaml` (default to `$HOME/.backhome.yaml`), copy `files` to local git repository.

It's not implemented yet but it will commit the changes and push to a `remote` git repository.

Features to add:
 - [ ] Backup local destinaton before copying files to it. Restore local if failed to copy one file.
 - [ ] `push` or `sync` command (commit changes and push to remote)
 - [ ] `restore` (copy files from local to their origin)
 - [ ] `setup`  (create basic config file, init local repository, add remote as its origin and pull the most recent version)
 - [ ] `add` and `remove` files from beign watched
 - [ ] run `copy` and `sync` periodically 
 - [ ] Alerts using macos notification
 - [ ] Support to multiple environments (work and home)
  
## Usage
TDB
