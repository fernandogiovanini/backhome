# Backhome

![CI Build](https://github.com/fernandogiovanini/backhome/actions/workflows/build.yml/badge.svg)

This CLI tool serves two purposes for me: learning Golang and acting as a simple backup tool for my configuration files (mostly dotfiles).

It copies the selected files to a local folder and then pushes it to a Github repository.

When executed for the first time it creates a `local` repository (defaults to `$HOME/backhome`) with a `backhome.yaml` in it.

To add a file to config file
```
backhome add path/to/file1 path/to/file2 ...
```

To copy files in config file to `local`
```
backhome copy
```

Setting `--local` initialize the `local` repository in another local.


**Github remotes are still to be implemented**

Features to add:
 - [ ] `push` / `sync` command (commit changes and push to remote)
 - [ ] `restore` (copy files from local to their origin)
 - [ ] `remove` files from configuration file
 - [ ] run `copy` and `sync`/`push` periodically 
 - [ ] Alerts using native (macos) notification
 - [ ] Support to multiple environments (work and home)
 - [ ] Better error handling (use typed error to decide how to print message)
 - [ ] Interactive init procedure (check if `local` and config file exist and ask if user wants to create before doing so) 
