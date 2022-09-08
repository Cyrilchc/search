# search
A command line search tool
## Build
```
go build
```
## Install
```
go install
```
## Get help
```
search -h
```
## Flags

- -s : Provided string will be used to search files and directories by name
- -r : The search will extend recursively to sub-folders
- -c : Enables case sensitivity (disabled by default)
- -d : Will display directories (display both files and directories if none of the -d and -f flags are used)
- -f : Will display files (display both files and directories if none of the -d and -f flags are used)
- -p : Will search in the provided path (if not provided, search will occur in the working directory)
- -ew : Will search files and folders that ends with provided search string
- -sw : Will search files and folders that starts with provided search string
- -v : Will display more information

