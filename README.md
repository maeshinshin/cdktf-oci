# cdktf-oci

This is a library for defining infrastructure components for Oracle Cloud Infrastructure using [CDK for Terraform](https://developer.hashicorp.com/terraform/cdktf).

## Technology Stack

- [Oracle Cloud Infrastructure](https://www.oracle.com/cloud/)
- [CDK for Terraform](https://developer.hashicorp.com/terraform/cdktf)
- [Go](https://golang.org/)

## environment setup

### Using direnv with Nix

```bash
$ cd cdktf-oci/
🔨 Welcome to devshell

[[general commands]]

  cdktf     - CDK for Terraform CLI
  go        - Go Programming language
  gotest    - go test with colors
  make      - Tool to control the generation of non-source files from sources
  menu      - prints this menu
  oci       - Command Line Interface for Oracle Cloud Infrastructure
  terraform - Tool for building, changing, and versioning infrastructure
```

### Using devshell with Nix

```bash
nix develop
🔨 Welcome to devshell

[[general commands]]

  cdktf     - CDK for Terraform CLI
  go        - Go Programming language
  gotest    - go test with colors
  make      - Tool to control the generation of non-source files from sources
  menu      - prints this menu
  oci       - Command Line Interface for Oracle Cloud Infrastructure
  terraform - Tool for building, changing, and versioning infrastructure
```

### Using npm

```bash
$ npm install --global cdktf-cli@latest
```

## Getting Started

generate the provider's code

```bash
cdktf get
```

check the provisioning plan

```bash
cdktf plan
```

apply the provisioning plan

```bash
cdktf apply
```
