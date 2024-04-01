# find

`find` utility from Linux written in Go.

## How to build?

Clone the project, open **src** directory and execute `make` command. The binary file will be located in the `build` folder in project root directory.

If there is a build error like `package <name> is not in std` then execute next commands in the root of the project:

```shell
go work init
go work use src/args src/find
```

And then repeat the build process.

## How to use?

`find` recursively searches for files and directories from given paths.

### Parameters

This utility supports a few flags which descriptions you can also read using `--help` flag.

- `-d` - print only directories
- `-f` - print only files
- `sl` - print only symlinks in form `link -> destination`. If destination is unreachable then `[broken]` will be shown
- `-ext` - print only files with specified extension. Works only if `-f` is specified. Do not specify any extensions to find only files without extension

Rest non-flag arguments will be treated as paths where to search files.

### Example

```shell
-> % ../build/find -d ..
..
../.git
../.git/branches
../.git/hooks
../.git/info
../.git/logs
../.git/logs/refs
../.git/logs/refs/heads
../.git/logs/refs/remotes
../.git/logs/refs/remotes/origin
../.git/objects
../.git/objects/info
../.git/objects/pack
../.git/refs
../.git/refs/heads
../.git/refs/remotes
../.git/refs/remotes/origin
../.git/refs/tags
../build
../src
../src/args
../src/find
```
