# Godot Engine Version Manager - README

Welcome to the repository for Godot Engine Version Manager or gevm for short! This tool is built using [Go](https://go.dev/) and [InnoSetup](https://jrsoftware.org/isinfo.php). It allows for the downloading of [Godot Engine](https://godotengine.org/) from it's [GitHub](https://github.com/godotengine/godot-builds) repository via the terminal. We use it as part of our CI/CD pipeline for game development. It can also be used for personal use but we would recommend [Godots](https://github.com/MakovWait/godots) for that.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Uninstall](#uninstall)
- [Contributing](#contributing)
- [License](#license)

## Installation

### Windows

1. Download the latest installer from the latest release page [here](https://github.com/bashidogames/gevm/releases/latest).
2. Run the installer.
3. Follow the on screen instructions.

### Linux

You can use the very basic helper install script to download and install the Linux binary:

```
curl -o- https://raw.githubusercontent.com/bashidogames/gevm/master/install.sh | bash
```

You will need to make sure `~/.local/bin` is in your `PATH` environment variable.

### Manual

Alternatively you can download the binary and use it manually:

1. Download the `zip` file for your system from the release page.
2. Extract the binary to a location of your choosing.
3. Add location to your `PATH` environment variable.

## Usage

Below is some very basic usage for the more common commands. You can use `gevm --help` to get a full list of commands.

### `versions`

Use the `list` command for listing available stable versions for your platform:

```
gevm versions list --mono --all
```

Additional optional flags include:

- `--mono` to specify mono versions instead.
- `--all` to also list non-stable releases.

### `godot`

Install a version of godot using the `install` command:

```
gevm godot install 4.3 --include-export-templates --release beta1 --mono --application --desktop
```

Additional optional flags include:

- `--include-export-templates` will additionally download and install the export templates for this version.
- `--release` specifies a non-stable release to use.
- `--mono` to specify the mono version of the engine.
- `--application` will attempt to create an application shortcut (start menu, app menu, etc).
- `--desktop` will attempt to create a desktop shortcut.

Uninstall a version by using the `uninstall` command:

```
gevm godot uninstall 4.3
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

### `cache`

This tool uses a download cache to making reinstalling versions quicker. You may want to free up space by using the `clear` command:

```
gevm cache clear
```

## Uninstall

The uninstallation process will not remove any installed versions of the engine or cached downloads. So you may want to uninstall any unwanted versions first and then also clear the cache to free up space:

```
gevm godot list
gevm godot uninstall x.x.x.x
...
gevm cache clear
```

### Windows

1. Open `Settings -> Apps -> Installed Apps` or search for `Add or remove programs` in the start menu.
2. Look for `Godot Engine Version Manager (gevm)` and uninstall.
3. Follow the on screen instructions.

### Linux

You can uninstall on Linux by just removing the binary:

```
rm -f ~/.local/bin/gevm
```

## Contributing

See [contributing](CONTRIBUTING) for more details.

## License

This project is licensed under the [MIT License](LICENSE).