Yaml config
Compare content periodically (cron?) and make a commit with changes 

Files:
* ~/.zshrc
* ~/.oh-my-posh-fernando.omp.json
* ~/.vimrc
* ~/.zshenv
* ~/.zprofile
* ~/.vim/colors/solarized.vim

Tasks
- [x] Base project (config and compiling)
- [x] Copy files to tmp dir
  - [x] Read a config file with a list of files to be copied
  - [x] Copy
- [ ] Check if there is a diff
  - [ ] Read the remote repo from config
  - [ ] Diff
  - [ ] Commit or error
- [ ] Commit and push
- [ ] Alert using MacOS notification
- [ ] Run periodically
- [ ] Command to setup backup (remote repo)
  - [ ] Create/edit yaml
- [ ] Command to add/remove file from watch
