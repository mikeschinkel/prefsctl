# Prefsctl

**Prefsctl** is a command-line tool written in Golang for managing macOS preferences using a YAML or JSON configuration file. Inspired by the conceptual model of `kubectl` for Kubernetes, Prefsctl allows software engineers, IT professionals, DevOps engineers, and power users to declaratively apply and manage preferences for their macOS system.

---

## Features

- Declarative application of macOS preferences using YAML/JSON files.
- Automatic generation of YAML/JSON resource documents for macOS preferences.
- Domain-specific queries to reduce noise in preference output.
- Support for applying user-managed preferences while avoiding system- or app-managed settings that may be harmful or irrelevant.

---

## Installation

Until pre-built binaries are available via GitHub Releases and Homebrew (planned features), please build from source as shown below:

```shell
# Clone the repository
git clone https://github.com/yourusername/prefsctl.git
cd prefsctl

# Build the binary
make build

# Optionally, move the binary to a directory in your PATH
mv ./bin/prefsctl /usr/local/bin/
```

---

## Usage

### Apply Preferences

To apply preferences from a file:

```shell
prefsctl apply -f path/to/filename.yaml
```

The file format is determined by its extension (`yaml`, `yml`, or `json`).

#### Explicitly Specify Format

You can also specify the file format explicitly using the `-o`/`--output` option:

```shell
prefsctl apply -o json -f path/to/filename
```

### Example YAML File

```yaml
---
apiVersion: prefs/v1beta1
kind: prefs
metadata:
  domain: com.apple.dock
  macos: monterey
spec:
  preferences:
  - name: size-immutable
    value: false
---
apiVersion: prefs/v1beta1
kind: prefs
metadata:
  domain: com.apple.dock
  macos: monterey
spec:
  preferences:
  - name: tilesize
    value: 65
  - name: mineffect
    value: genie
```

---

## Generating YAML for Preferences

Prefsctl can generate YAML documents containing current macOS preference values using the `get prefs` command:

```shell
prefsctl get prefs -o=yaml
```

This will output a YAML document containing all preferences that differ from the default values for the current macOS version. Note that currently only preference domains that begin with `com.apple.` are included in the output.

If you want to limit the output to just a specific preference domain, you can use the `--domains` switch as shown below:

```shell
prefsctl get prefs -o=yaml --domains=com.apple.dock
```

This example limits the output to preferences in the `com.apple.dock` domain only.

#### Example Command

To generate preferences for multiple domains:

```shell
prefsctl get prefs -o=yaml --domains=com.apple.dock,com.apple.finder

# or

prefsctl get prefs -o=yaml --domains=com.apple.dock --domains=com.apple.finder
```

---

## Understanding Default Preference Values

macOS does not provide an explicit way to query default values for preferences. Instead, preferences that have not been explicitly set may not appear at all, making their default values implicit. Prefsctl uses the concept of “zero” values, similar to Golang, where the default value is assumed if no preference value exists.

To address this:

- As part of development and testing, the author installs a pristine version of macOS in a virtual machine to record default preferences.
- The author then creates a local account with minimal configuration.
- And finally, the author then uses the Prefsctl `get prefs` command to record default values as a Go function to be incorporated back into an updated revision of Prefsctl.

This allows Prefsctl to infer default values for preferences and reduce unnecessary noise when running `get prefs`.

---

## Classification of Preferences

macOS preferences encompass user-managed settings, system-managed settings, app-managed settings, runtime state, and version migration data. Prefsctl focuses on user-managed preferences to avoid noise and harmful configurations.

### Preference Research Methodology

To identify user-managed preferences:

- **Manual Research:** Testing in pristine macOS virtual machines.
- **Generative AI Assistance:** Using tools like ChatGPT and Claude for initial classification.
- **Planned feature:** Prefsctl will include APIs for automated preference classification and storage.

---

## Internal Structure and Extensibility

Prefsctl maintains a Golang struct that contains:

- Known preference domains.
- Default values and types.
- Classification (user-managed or otherwise).
- Descriptions and usage notes.

Prefsctl currently includes detailed research only for `com.apple.dock` preferences. Users can override default classifications with custom YAML files if needed.

---

## Planned Enhancements

- **Support for Complex Data Types:** Currently, only scalar values are supported (e.g., integers, strings, booleans). Support for array and dictionary types (e.g., `__NSDictionaryI`, `__NSArrayM`, etc.) is planned, where applicable.
- **Custom Overrides:** Users will be able to merge their own YAML files to override hardcoded default values.
- **Increased Preference Coverage:** Additional preference domains will be added incrementally.
- **Ease of Installation:** Plans to provide downloadable binaries via GitHub releases and Homebrew support are in progress.

---

## Example Use-Cases

- **macOS Setup Automation:** Apply saved preferences when setting up a new Mac or virtual machine.
- **Version Control for Preferences:** Track macOS preferences in version-controlled YAML or JSON files.
- **Preference Documentation:** Generate YAML files to document the current state of macOS preferences.

---

## Contributing

Contributions to Prefsctl are welcome! Here’s how you can help:

- **Submit Issues:** Report bugs or request new features.
- **Pull Requests:** Submit improvements or new preference classifications.
- **Knowledge Sharing:** Help identify preference defaults and classifications to reduce the number of unnecessary preference domains and values generated by `get prefs`. Contributions that clarify which preferences differ from effective defaults and classify user-managed vs system-managed settings are especially valuable.

---

## Feedback and Support

If you encounter any issues or have questions, please submit an issue on GitHub or open a discussion. Contributions of any kind, including new use-cases and enhancements, are greatly appreciated.

---

## License

Prefsctl is released under the MIT License. See `LICENSE` for more details.

---

## Future Directions

Prefsctl’s vision is to become the `kubectl` for  `brew bundle` for macOS preferences. By encoding known preference attributes and supporting customizable overrides, it aims to empower technical users to manage macOS configurations with precision and ease.

