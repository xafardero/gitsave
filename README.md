# gitsave - Save your all your organization repositories

## 📜 Table of contents

- [🚀 Usage](#️-usage)
- [🏗️ Installation](#-installation)
- [🛠️ Configuration](#️-configuration)

---

## 🚀 Usage

CLI app to download all the repositories of a github organization.

Usage
```
gitsave --gh-token tour-token-here --org organization-name
```

Flags:

* githubToken: Github token ID.
* githubOrganization: Github organization name.
* folder : Folder to clone repositorires. Default: `/tmp/gitsave` 
* sshPrivateKey: Private key used to clone repositories. Default: `$HOME/.ssh/id_rsa`



---

## 🏗️ Installation

### From binary (Linux)

```bash
sudo wget -O /usr/local/bin/gitsave https://github.com/xafardero/gitsave/releases/latest/download/gitsave
```
```bash
sudo chmod +x /usr/local/bin/gitsave
```

### From binary (Mac)

```bash
sudo wget -O /usr/local/bin/gitsave https://github.com/xafardero/gitsave/releases/latest/download/gitsave_darwin
```
```bash
sudo chmod +x /usr/local/bin/gitsave
```

### From source code

```bash
wget https://github.com/xafardero/gitsave/releases/latest/download/gitsave.zip
```
```bash
go build -o gitsave *.go
```
```bash
sudo mv gitsave /usr/local/bin/gitsave
```

---

## 🛠️ Configuration

To get the github token go to https://github.com/settings/tokens and Generate a new token with the repo scope selected.
