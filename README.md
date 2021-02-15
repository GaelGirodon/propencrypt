# Propencrypt

Encrypt and decrypt multiple properties, in multiple files, at once.

## About

There are many solutions out there to manage file encryption:

- **GPG** provides easy file encryption, but the output binary file is not
  _Git-friendly_ (no diff, code review is more difficult, etc.).
- **Ansible Vault** encrypts variables and files. With variable encryption,
  files are still easily legible (plaintext and encrypted variables can be mixed
  in the same file), but each string must be encrypted individually making it
  tedious to work with many properties in multiple files. Furthermore, Ansible
  Vault is primarily made to work with Ansible.
- **Jasypt** (through the integration for Spring Boot) provides a Maven plugin
  allowing encrypting multiple placeholders (`DEC(...)`) at once in an
  `application.properties` file, but Jasypt is meant to be used as a library
  in a Java application, so it is not language-agnostic.

**Propencrypt** aims to (_modestly_) solve these limitations by providing the
following features:

- Encrypt and decrypt multiple properties in multiples files at once (using
  AES-256-GCM), without requiring to encrypt the entire contents of the
  files, making them _Git-friendly_.
- Handle multiple file formats (`yaml`, `properties`, etc.): the `pattern`
  option is used to find values to encrypt.
- Language-agnostic: encrypt files to store them safely in a Git repository,
  and decrypt them back before using them as you want (e.g. to create a K8s
  secret).
- Lightweight (~1 MB to download), dependency-free, easy to install and run.

## Install

> `DOWNLOAD_URL`: <https://github.com/GaelGirodon/propencrypt/releases/latest/download>

Download and extract `propencrypt`:

```shell
# Linux
curl -sL "$DOWNLOAD_URL/propencrypt_linux_amd64.tar.gz"  | tar xvz
```

```powershell
# Windows
Invoke-WebRequest "$DOWNLOAD_URL/propencrypt_windows_amd64.zip" -OutFile "propencrypt.zip"
Expand-Archive "propencrypt.zip" -DestinationPath ./
```

## Usage

Encrypt and decrypt properties:

```shell
propencrypt encrypt -k <key> -p <pattern> [-e <ext>] <files>
propencrypt decrypt -k <key>              [-e <ext>] <files>
```

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

Encrypt passwords using the `encrypt` command:

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