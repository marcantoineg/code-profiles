# 🖥 code-profiles 💻

[![CI](https://github.com/marcantoineg/code-profiles/actions/workflows/ci.yml/badge.svg)](https://github.com/marcantoineg/code-profiles/actions/workflows/ci.yml)
<img height="20px" src="https://img.shields.io/badge/Golang-FFFFFF?logo=go&style=flat">

Profiles for VS Code Extensions.

I plan to use this to open only the extensions I need in VS Code depending on projects.

You need 2 things:
- a `code-profiles.yml` where your executable is. This is where you define your profiles.
- a `.code-profile` with a valid profile id/name. This will start code using `--extensions-dir` with the right extensions.

You have multiple commands:
- `open`: which opens code with a profile.
<br>_From file in cwd (`.code-profile`) or args (`open` vs `open some-profile`)_
- `install`: which installs required extensions for the profile.
<br>_From file in cwd (`.code-profile`) or args (`install` vs `install some-profile`)_
- `profile add [name]`: create config and/or appends to config a blank profile
