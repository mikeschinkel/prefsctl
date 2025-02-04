# Prefsctl

**CURRENTLY A WORK-IN-PROGRESS**

**Prefsctl** is a command-line tool written in Golang for managing macOS preferences using a YAML or
JSON configuration file. Inspired by the conceptual model of `kubectl` for Kubernetes, Prefsctl
allows software engineers, IT professionals, DevOps engineers, and power users to declaratively
apply and manage preferences for their macOS system.

---

## Features

- Declarative application of macOS preferences using YAML/JSON files.
- Automatic generation of YAML/JSON resource documents for macOS preferences.
- Domain-specific queries to reduce noise in preference output.
- Support for applying user-managed preferences while avoiding system- or app-managed settings that
  may be harmful or irrelevant.

## com.apple domains only

Note that Prefsctl **_currently_** only supports [preference domains](https://macos-defaults.com/)
that begin with `com.apple`.

However, we may add support for other domains in future if demand exists and users are willing to
help research the preference names, types and defaults values in those domains.

Please [start a discussion](https://github.com/mikeschinkel/prefsctl/discussions/new?category=ideas)
if you are interested in Presctl supporting other preference domains besides `com.apple`.


---

## Installation

Until pre-built binaries are available via GitHub Releases and Homebrew (planned features), please
build from source as shown below:

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

The file format is determined by its extension (`yaml`, `yml`=>`yaml`, or `json`).

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

Users can use Prefsctl to generate resource documents they can they use with the `apply`
command. Those resource documents will contain only the "user-managed" preference values that differ
from _"known"_ defaults using the `get prefs` command:

```shell
prefsctl get prefs -o=yaml
```

To generate a resource document that ignores the default values and outputs all user-managed
preference values for the specified domain(s), or all domains if no domains are specified — use the
`--include-unchanged` flag:

```shell
prefsctl get prefs -o=yaml --include-unchanged
```

If you want to limit the output to just a specific preference domain, you can use the `--domains`
switch as shown below which limits the output to preferences in the `com.apple.dock`
domain only:

```shell
prefsctl get prefs -o=yaml --domains=com.apple.dock
```

### Example Command

To generate preferences for multiple domains:

```shell
prefsctl get prefs -o=yaml --domains=com.apple.dock,com.apple.finder

# or

prefsctl get prefs -o=yaml --domains=com.apple.dock --domains=com.apple.finder
```

---

## Understanding Default Preference Values

macOS does not provide an explicit way to query default values for preferences. Instead, preferences
that have not been explicitly set may not appear at all, making their default values implicit.
Prefsctl uses the concept of “zero” values, similar to Golang, where the default value is assumed if
no preference value exists.

To address this:

- As part of development and testing, the author installs a pristine version of macOS in a virtual
  machine to record default preferences.
- The author then creates a local account with minimal configuration.
- And finally, the author then uses the Prefsctl `get prefs` command to record default values as a
  Go function to be incorporated back into an updated revision of Prefsctl.

This allows Prefsctl to infer default values for preferences and reduce unnecessary noise when
running `get prefs`.

---

## Classification of Preferences

macOS preferences encompass different types of settings that I classified into five different
groups:

1. User-managed settings,
2. System-managed settings,
3. App-managed settings,
4. Runtime state, and
5. Version migration data.

Prefsctl focuses **user-managed preferences** per my classification to minimize noise when running
`get prefs`, and to avoid generating potentially harmful configurations. Classifying preferences is
a gargantuan task, so any help you can provide there will be appreciated.

Also, users will soon be able to can override default values and classifications with custom YAML
files, if needed.

### Preference Research Methodology

To identify user-managed preferences:

- **Manual Research:** I used pristine macOS virtual machines to determine default values for
  preferences available at initial install.
- **Research with AI:** I have used ChatGPT and Claude for some initial classification of
  preferences.
- **Automated Research:** I plans in incorporate the use of Gen AI APIs into Prefsctl to enable
  automated preference classification for myself when developing but also for anyone who is
  interested in helping document defaults and classify preferences.

---

## Internal Structure and Extensibility

Prefsctl maintains a Golang struct that contains:

- Known preference domains.
- Default values and types.
- Classification (user-managed or otherwise).
- Descriptions and usage notes.

Prefsctl currently includes detailed research only for `com.apple.dock` preferences.

---

## Planned Enhancements

- **Complex Preference Types:** Currently, only scalar values are supported (e.g., integers,
  strings, booleans). Support for array and dictionary types (e.g., `__NSDictionaryI`, `__NSArrayM`,
  etc.) is planned, where applicable.
- **Custom Overrides:** Merge of user YAML files to override hardcoded default values, and also
  filters _(TODO: document filters)._
- **Increased Preference Coverage:** Additional preference domains will be added incrementally.
- **Ease of Installation:** Plans are to provide downloadable binaries via GitHub releases and
  Homebrew support before we exist beta.

---

## Example Use-Cases

- **macOS Setup Automation:** Apply saved preferences when setting up a new Mac or virtual machine.
- **Version Control for Preferences:** Track macOS preferences in version-controlled YAML or JSON
  files.
- **Preference Documentation:** Generate YAML files to document the current state of macOS
  preferences.

---

## Contributing

Contributions to Prefsctl are welcome! Here’s how you can help:

- **Submit Issues:** Report bugs or request new features.
- **Pull Requests:** Submit improvements or new preference classifications.
- **Share your Knowledge:** Help identify preference defaults and classifications to reduce the
  number of unnecessary preference domains and values generated by `get prefs`. Contributions that
  clarify which preferences differ from effective defaults and classify user-managed vs
  system-managed settings are especially valuable.
- **Presenting Use-cases:** I built Prefsctl to scratch my own itch, but I am anxious to learn how
  it can help others meet their needs that might require additional features or enhancements. Start
  a [new discussion](https://github.com/mikeschinkel/prefsctl/discussions/new?category=ideas) to
  tell me about your ideas.

---

## Feedback and Support

If you encounter any issues or have questions, please submit an issue on GitHub or open a
discussion. Contributions of any kind, including new use-cases and enhancements, are greatly
appreciated.

---

## License

Prefsctl is released under the MIT License. See `LICENSE` for more details.

---

## Future Directions

Prefsctl’s vision is to become the `kubectl`for`brew bundle` for macOS preferences. By encoding
known preference attributes and supporting customizable overrides, it aims to empower technical
users to manage macOS configurations with precision and ease.
