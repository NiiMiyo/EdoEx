# EdoEx

**EdoEx** is a [CLI](https://en.wikipedia.org/wiki/Command-line_interface "Command-Line Interface") tool to help you create and add custom cards, archetypes and counters into your [EDOPro](https://projectignis.github.io/ "The bleeding-edge automatic duel simulator").

The goal of EdoEx is to simplify expansion creation for EDOPro by using YAML files to set card properties (name, atk, def, abilities etc.) with a fast and simple to use CLI tool.

## Table of Contents

- [EdoEx](#edoex)
    - [Table of Contents](#table-of-contents)
    - [Installation](#installation)
        - [Pre-built binaries](#pre-built-binaries)
        - [Building from source](#building-from-source)
    - [After installation](#after-installation)
        - [Schemas](#schemas)
    - [Usage](#usage)
    - [Contributing](#contributing)
    - [License](#license)

---

## Installation

### Pre-built binaries

Download a pre-built portable binary release from the [releases page](https://github.com/NiiMiyo/EdoEx/releases/latest) for your system.

Be aware that only the Windows x64 release was tested since I can't test for other operating systems.

### Building from source

If your system is not listed in the releases page but does run [Go](https://go.dev/) you can build it from source.

In this case, clone this repository and run `make build`. EdoEx now should be installed into the `bin` folder.

```bash
git clone https://github.com/NiiMiyo/EdoEx
cd EdoEx
make build
```

In case your system does not support `make` or you can't install it, compile the program into the `bin` folder and copy the content of [`buildfiles`](/buildfiles/) into that directory.

```bash
git clone https://github.com/NiiMiyo/edoex
go build -o bin/ edoex
cp -r buildfiles/* bin/
```

## After installation

After you have a working EdoEx installation on your system, add the installation folder into your PATH. Then you can change the `gamedir` property on `edoex.config.yaml` to the folder you have EDOPro installed (for example, `C:/ProjectIgnis`). This can be configured independently for each expansion.


### Schemas

The installation folder contains a `schemas` subfolder with [JSON Schemas](https://json-schema.org) for your files. Some IDEs and Text Editors support JSON Schemas for YAML files, if that is the case for you, is highly recommended that you apply the schemas in your editor.

- `card_schema.json` should apply to any `.yaml` file inside a `cards` directory (`cards/**/*.yaml`).
- `meta_schema.json` should apply to any `.yaml` file inside a `meta` directory (`meta/**/*.yaml`).
- `config_schema.json` should apply only to `edoex.config.yaml` in the root of your project.

---

## Usage

EdoEx is designed to be a simple tool, running `edoex help` into your terminal will print all available commands. Start by running `edoex init my-first-expansion` on an empty folder.

---

## Contributing

If you found a bug or have some idea on how to make EdoEx better head to the [issues page](https://github.com/NiiMiyo/EdoEx/issues?q=is%3Aissue) and check if your issue has not yet been reported.

- If it was, you can comment to improve the discussion;
- If it wasn't, create a new issue with your problem/idea;
    - If it is a problem, try to put as much information you can on how to reproduce it. At least this items should be on your issue:
        - EdoEx version;
        - Your Operating System and architecture;
        - Error message;
    -  If it is an idea, try to put as much information you can on what it is, what it does and how it works.

You can also contribute by pulling requests in the [pull requests page](https://github.com/NiiMiyo/EdoEx/pulls), although I recommend you to do something worth your time instead.

---

## License

[MIT](/LICENSE)
