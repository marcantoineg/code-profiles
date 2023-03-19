# code-profiles 💻

[![CI](https://github.com/marcantoineg/code-profiles/actions/workflows/ci.yml/badge.svg)](https://github.com/marcantoineg/code-profiles/actions/workflows/ci.yml)
<img height="20px" src="https://img.shields.io/badge/Golang-FFFFFF?logo=go&style=flat">

## Simply, what is it?
Profiles for VS Code Extensions.

I use this to load only the extensions I need in VS Code depending on projects.


## How does it work?
You need 2 things:
- a `code-profiles.yml` where your executable is. This is where you define your profiles.
- a `.code-profile` with a valid profile name. This will start VSCode using `--extensions-dir` with the right extensions.

You have multiple commands:
- `open`: which opens VSCode with a specified profile.
<br>_From the `.code-profile` file in the current working directory or args (`code-profiles open [profile_name]`)_
- `install`: which installs required extensions for a specified profile.
<br>_From the `.code-profile` file in the CWD or args (`code-profiles install [profile_name]`)_

## Usage
Basic commands:
```
Usage:
  code-profiles [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  install     install required VSCode extensions for a given profile
  open        open VSCode using a custom profile for extensions

Flags:
  -h, --help   help for code-profiles
```

### install
```
install required VSCode extensions for a given profile

Usage:
  code-profiles install [profile_name] [flags]

Flags:
  -c, --config-path string   Path to code-profiles config (default "./code-profiles.yml")
  -h, --help                 help for install
  -v, --verbose              prints additional logs
```

### open
```
open VSCode using a custom profile for extensions

Usage:
  code-profiles open [profile_name] [flags]

Flags:
  -c, --config-path string   Path to code-profiles config (default "./code-profiles.yml")
  -h, --help                 help for open
  -i, --install              should install extensions before opening VSCode
  -v, --verbose              prints additional logs
```
