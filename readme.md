# DiscogsAssembly (discasm)

![discasm](logo.jpg)

**Before using application you should add DISCASM_HOME enviroment variable which contains path to executable file**

Example for linux:

Add: ```export DISCASM_HOME="/root/DiscogsAssembly/build"``` to the end of .bashrc file and then use ```source .bashrc``` to confirm changes


```
Discasm is a tool for downloading images from discogs release

Usage:
  discasm [command]

Available Commands:
  download    Download release images and metadata by it's discogs id
  help        Help about any command
  metadata    Download release metadata by it's discogs id
  release     Fetch release by it's discogs id
  token       Update token in configuration
  whoami      Display current user information

Flags:
  -h, --help      help for discasm
  -v, --version   version for discasm

Use "discasm [command] --help" for more information about a command.
```