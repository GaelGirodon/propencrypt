# Propencrypt

[![release](https://img.shields.io/github/v/release/GaelGirodon/propencrypt?style=flat-square)](https://github.com/GaelGirodon/propencrypt/releases/latest)
[![license](https://img.shields.io/github/license/GaelGirodon/propencrypt?color=blue&style=flat-square)](./LICENSE)
[![build](https://img.shields.io/github/actions/workflow/status/GaelGirodon/propencrypt/build.yml?branch=main&style=flat-square)](https://github.com/GaelGirodon/propencrypt/actions/workflows/build.yml)
[![tests](https://img.shields.io/endpoint?style=flat-square&url=https%3A%2F%2Fgist.githubusercontent.com%2FGaelGirodon%2Ffbde4d59b7dd3c4f2cc9c4fea3497ae1%2Fraw%2Fpropencrypt-go-tests.json)](https://github.com/GaelGirodon/propencrypt/actions/workflows/build.yml)
[![coverage](https://img.shields.io/endpoint?style=flat-square&url=https%3A%2F%2Fgist.githubusercontent.com%2FGaelGirodon%2Ffbde4d59b7dd3c4f2cc9c4fea3497ae1%2Fraw%2Fpropencrypt-go-coverage.json)](https://github.com/GaelGirodon/propencrypt/actions/workflows/build.yml)
[![docker](https://img.shields.io/docker/v/gaelgirodon/propencrypt?color=%232496ed&label=docker&logo=docker&logoColor=white&style=flat-square)](https://hub.docker.com/r/gaelgirodon/propencrypt)
[![npm](https://img.shields.io/npm/v/propencrypt?color=%23cb3837&logo=npm&style=flat-square)](https://www.npmjs.com/package/propencrypt)

Encrypt and decrypt multiple properties, in multiple files, at once.

## About

**Propencrypt** provides the following features:

- Encrypt and decrypt multiple properties in multiples files at once using the
  AES-256-GCM symmetric algorithm, without requiring to encrypt the entire
  contents of the files, making them _Git-friendly_.
- Handle multiple file formats (`yaml`, `properties`, etc.): the `pattern`
  option is used to find values to encrypt.
- Language-agnostic: encrypt files to store them safely in a Git repository,
  and decrypt them back before using them as you want (e.g., to create a K8s
  secret).
- Lightweight (~1 MB to download), dependency-free, easy to install and run.

It aims to (_modestly_) solve some limitations of these encryption tools:

- **GPG** provides easy file encryption, but the output binary file is not
  _Git-friendly_ (no diff available, code review is more difficult, etc.).
- **Ansible Vault** encrypts variables and files. With variable encryption,
  files are still easily legible (plaintext and encrypted variables can be mixed
  in the same file), but each string must be encrypted individually making it
  tedious to work with many properties in multiple files. Furthermore, Ansible
  Vault is primarily made to work with Ansible.
- **Jasypt** (through the integration for Spring Boot) provides a Maven plugin
  allowing encrypting multiple placeholders (`DEC(...)`) at once in an
  `application.properties` file, but Jasypt is meant to be used as a library
  in a Java application, so it is not language-agnostic.

## Install

Download and extract the
[latest release](https://github.com/GaelGirodon/propencrypt/releases/latest):

```shell
# Linux (Bash)
DOWNLOAD_URL="https://github.com/GaelGirodon/propencrypt/releases/latest/download"
curl -sL "$DOWNLOAD_URL/propencrypt_linux_amd64.tar.gz" | tar xvz
```

```powershell
# Windows (PowerShell)
$DOWNLOAD_URL = "https://github.com/GaelGirodon/propencrypt/releases/latest/download"
Invoke-WebRequest -OutFile "propencrypt.zip" "$DOWNLOAD_URL/propencrypt_windows_amd64.zip"
Expand-Archive "propencrypt.zip" -DestinationPath ./
```

## Usage

Encrypt and decrypt properties in files:

```shell
propencrypt encrypt -k <key> -p <pattern> [-e <ext>] <files>
propencrypt decrypt -k <key>              [-e <ext>] <files>
```

Run `propencrypt --help` to show the help message and
`propencrypt help <command>` to get more information about a given command.

### Commands

#### encrypt

The `encrypt` command reads input files (a list of file names,
[glob patterns](https://golang.org/pkg/path/filepath/#Match) are supported),
encrypts and encodes each property value matched by the provided pattern (the
capturing group is used to find the value) and creates output files where values
are replaced by their encrypted counterpart, encoded as Base64 and surrounded by
`ENC(<...>)`. The name of each output file is the concatenation of the
associated input file name and the extension.

```shell
encrypt -k <key> -p <pattern> [-e <ext>] <files>
```

#### decrypt

The `decrypt` command reads input files (a list of file names,
[glob patterns](https://golang.org/pkg/path/filepath/#Match) are supported),
decodes and decrypts each property value matched by the `ENC(<...>)` pattern and
creates output files where values are replaced by their unbounded (`ENC()` is
removed), decoded (from Base64) and decrypted counterpart. The name of each
output file is the name of the associated input file without the extension.

```shell
decrypt -k <key> [-e <ext>] <files>
```

### Options

| Option            | Description                       | Default | Environment variable  |
| ----------------- | --------------------------------- | ------- | --------------------- |
| `-k`, `--key`     | 256-bit encryption key            |         | `PROPENCRYPT_KEY`     |
| `-p`, `--pattern` | Sensitive property pattern        |         | `PROPENCRYPT_PATTERN` |
| `-e`, `--ext`     | File extension to append / remove | `.enc`  | `PROPENCRYPT_EXT`     |

> **Note**: options set from the command-line take precedence over the
> environment variables.

#### key

`key` is a 32-bytes string used as the symmetric key for properties values
encryption and decryption with the AES-256-GCM algorithm.

#### pattern

`pattern` is a regular expression used to find values to encrypt in files.
It must contain exactly one capturing group that matches the property value.
This pattern allows finding properties with different names (e.g.
`(?:pass|secret|login)=(.+)`), in multiples file types (e.g., `prop: (.+)` for
YAML, `prop=(.+)` for INI/Properties), etc.

#### ext

`ext` is the extension of the output encrypted file. By default, input files
are not modified during encryption: output files with encrypted values are
created as `<input-file-name><ext>`. This extension is removed from the
encrypted file name during decryption to get back to the input file name
(original unencrypted files are overridden if they exist). This extension
can be set to an empty value (`--ext ""`) to edit files in place.

## Example

Given `config.yml`, a configuration file where passwords need to be encrypted:

```yml
database:
  url: mysql://host/db
  username: app
  password: secret
```

Encrypt passwords using the `encrypt` command (add `--ext ""` to edit the file
in place):

```shell
propencrypt encrypt --key <key> --pattern "password: (.+)" config.yml
```

A new file with encrypted passwords, `config.yml.enc`, is created:

```yml
database:
  url: mysql://host/db
  username: app
  password: ENC(<base64-encrypted-value>)
```

It can be decrypted back to `config.yml` using the `decrypt` command:

```shell
propencrypt decrypt --key <key> config.yml.enc
```

## License

**Propencrypt** is licensed under the GNU General Public License.
