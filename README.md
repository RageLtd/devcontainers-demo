# DevContainer Demo

## Prerequisites

- macOS
- [Homebrew](https://brew.sh/)
- [Visual Studio Code](https://code.visualstudio.com/)

## Installation

### 1. Install OrbStack

OrbStack is a fast, lightweight, and easy-to-use container runtime for macOS.

```sh
brew install --cask orbstack
```

### 2. Start OrbStack

After installation, start OrbStack:

```sh
orbstack start
```

### 3. Install VSCode Extensions

Make sure you have the following VSCode extensions installed:

- [Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

## Open the DevContainer

1. Once the folder is open in VSCode, press `F1` to open the command palette.
2. Type `Remote-Containers: Reopen in Container` and select it.
3. VSCode will now build and open the DevContainer.

You are now ready to start developing in your DevContainer environment!
