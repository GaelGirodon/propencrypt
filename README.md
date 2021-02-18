# Propencrypt

[![release](https://img.shields.io/github/v/release/GaelGirodon/propencrypt?style=flat-square)](https://github.com/GaelGirodon/propencrypt/releases/latest)
[![license](https://img.shields.io/github/license/GaelGirodon/propencrypt?color=blue&style=flat-square)](./LICENSE)
[![build](https://img.shields.io/gitlab/pipeline/GaelGirodon/propencrypt/master?style=flat-square)](https://gitlab.com/GaelGirodon/propencrypt/-/pipelines/latest)
[![coverage](https://img.shields.io/gitlab/coverage/GaelGirodon/propencrypt/master?style=flat-square)](https://gitlab.com/GaelGirodon/propencrypt/-/pipelines/latest)

Encrypt and decrypt multiple properties, in multiple files, at once.

## About

**Propencrypt** provides the following features:

- Encrypt and decrypt multiple properties in multiples files at once using the
  AES-256-GCM symmetric algorithm, without requiring to encrypt the entire
  contents of the files, making them _Git-friendly_.
- Handle multiple file formats (`yaml`, `properties`, etc.): the `pattern`
  option is used to find values to encrypt.
- Language-agnostic: encrypt files to store them safely in a Git repository,
  and decrypt them back before using them as you want (e.g. to create a K8s
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

| Option            | Description                       | Default |
| ----------------- | --------------------------------- | ------- |
| `-k`, `--key`     | 256-bit encryption key            |         |
| `-p`, `--pattern` | Sensitive property pattern        |         |
| `-e`, `--ext`     | File extension to append / remove | `.enc`  |

Run `propencrypt --help` to show the full help message and
`propencrypt help <command>` to get more information about a given command.

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
