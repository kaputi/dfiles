# cmd:
    if no flags, it runs status
    -d : starts daemon

## subcommands
- status:
    shows changes, and any untracked file inside of the include(without exclude)
- git <...git args>:
    it runs git for the bare repo, NOTE: BE CAREFUL, it can mess up everything. maybe only certain things should be enabled 
- add <file/dir>:
    adds to git and includes full path(relative) as pattern
- remove <file/dir>:
    remove from git and include, will stop tracking but will keep history until now, it doesnt add to exclude.
- include <pattern>:
    adds to include (no git add)
- exclude <pattern>:
    adds to exclude but not doesnt remove, if there are files already tracking, they will remain tracking but any new file that matches the pattern will be excluded
- check <msg>:
    shows changed files
    walks through includes, shows any untracked files (there may be new files inside included dir), and adds them to git (NOT COMMAND TO AVOID INCLUDE)
    if msg is passed it commits, or something like this
- commit <message>:
    stages changes and commit
- push

## daemon
it checks, if changes it commits, if commit it pushes (if remote)
