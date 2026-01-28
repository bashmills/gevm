# Godot Engine Version Manager (gevm) - README

Welcome to the repository for Godot Engine Version Manager! This tool is built using [Go](https://go.dev/) and [InnoSetup](https://jrsoftware.org/isinfo.php). It allows for the downloading of [Godot Engine](https://godotengine.org/) from it's [godot-builds](https://github.com/godotengine/godot-builds) repository via the terminal. We use it as part of our CI/CD pipeline for game development. It can also be used for personal use but we would recommend [Godots](https://github.com/MakovWait/godots) for that case.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Uninstallation](#uninstallation)
- [Contributing](#contributing)
- [License](#license)

## Installation

### Linux / Mac

You can use the very basic helper install script to download and install the binary:

```
curl -o- https://raw.githubusercontent.com/bashmills/gevm/master/install.sh | bash
```

You will need to make sure `~/.local/bin` is in your `PATH` environment variable:

```
export PATH="~/.local/bin:$PATH"
```

### Windows

1. Download the latest installer from the latest release page [here](https://github.com/bashmills/gevm/releases/latest).
2. Run the installer.
3. Follow the on screen instructions.

### Manual

Alternatively you can download the binary and use it manually:

1. Download the `zip` file for your system from the release page.
2. Extract the binary to a location of your choosing.
3. Add location to your `PATH` environment variable.

## Usage

Below is some very basic usage for the more common commands. You can use `gevm --help` to get a full list of commands.

### `versions`

Use the `list` command for listing available versions for your platform:

```
gevm versions list --mono --all
```

| Flag | Short | Description |
| --- | --- | --- |
| `--mono` | `-m` | List the mono versions instead. |
| `--all` | `-a` | Also list non-stable releases. |

View versions for all platforms using the `detailed` command:

```
gevm versions detailed -m -a
```

### `godot`

Install a version of godot using the `install` command:

```
gevm godot install 4.3 --release beta1 --mono
```

| Flag | Short | Description |
| --- | --- | --- |
| `--exclude-export-templates` | `-e` | Exclude export templates from the command. |
| `--release` | `-r` | Specify a non-stable release to use. |
| `--mono` | `-m` | Use the mono version. |

Uninstall a specific version and the export templates by using the `uninstall` command:

```
gevm godot uninstall 4.3 -r beta1 -m
```

You can use the `path` command to print the path to the specified version if installed. You can use this from external tools to get the godot path for running builds:

```
gevm godot path 4.3 -r beta1 -m
```

Use the `list` command to show all currently installed versions:

```
gevm godot list
```

Uninstall all versions and export templates by using the `clear` command:

```
gevm godot clear
```

### `settings`

There are configuration settings you can change such where to put installed versions. Use the `list` command to list all the settings you can change and their current values:

```
gevm settings list
```

You can then change any of these settings by using the `set` command:

```
gevm settings set godot-root-directory <path>
```

Use the `reset` command to reset all settings to defaults:

```
gevm settings reset
```

### `cache`

This tool uses a download cache to make reinstalling versions quicker. You may want to free up space by using the `clear` command:

```
gevm cache clear
```

## Uninstallation

The uninstallation process will not remove any installed versions or cached downloads so you may want to that first to free up space:

```
gevm godot clear
```
```
gevm cache clear
```

### Linux / Mac

You can uninstall by just removing the binary:

```
rm ~/.local/bin/gevm
```

### Windows

1. Open `Settings -> Apps -> Installed Apps` or search for `Add or remove programs` in the start menu.
2. Look for `Godot Engine Version Manager (gevm)` and uninstall.
3. Follow the on screen instructions.

## Contributing

See [contributing](CONTRIBUTING.md) for more details.

## License

This project is licensed under the [MIT License](LICENSE).
