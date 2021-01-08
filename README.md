# fusion

[![Release](https://img.shields.io/github/release/edgelaboratories/fusion.svg?style=for-the-badge)](https://github.com/edgelaboratories/fusion/releases/latest)
![License](https://img.shields.io/github/license/edgelaboratories/fusion?style=for-the-badge)
[![Go version](https://img.shields.io/github/go-mod/go-version/edgelaboratories/fusion.svg?style=for-the-badge)](https://github.com/edgelaboratories/fusion)
[![Test status](https://img.shields.io/github/workflow/status/edgelaboratories/fusion/Test?label=Tests&style=for-the-badge)](https://github.com/edgelaboratories/fusion/actions?workflow=Test)
[![GolangCI status](https://img.shields.io/github/workflow/status/edgelaboratories/fusion/Golangci-Lint?label=Lint&style=for-the-badge)](https://github.com/edgelaboratories/fusion/actions?workflow=Golangci-Lint)

fusion is CLI tool to merge [JSON](https://www.json.org/json-en.html), [YAML](https://yaml.org/) or [TOML](https://toml.io/en/) files.

## Installation

Download from the [Releases page](https://github.com/edgelaboratories/fusion/releases) and put it somewhere in your `$PATH`.

## Usage

Merge YAML files together with:

```shell
❯ fusion yaml file1.yml file2.yml ... fileN.yml -o result.yml
```

The same goes for JSON or TOML files, simply use the appropriate sub-command. E.g.

```shell
❯ fusion json file1.json file2.json ... fileN.json -o result.json
```

Note that since [YAML is a superset of JSON](https://yaml.org/spec/1.2/spec.html#id2759572), merging YAML and JSON files into YAML is also possible:

```shell
❯ fusion yaml file1.yml file2.json ... fileN.yml -o result.yml
```

## Versioning

The versioning is done using [semver](https://semver.org/).

When a new release must be performed, use [mroth/bump](https://github.com/mroth/bump) to create a Git tag and the corresponding Github release,
then, once pushed, the [Goreleaser Github Action workflow](.github/workflows/release.yml) will create the binaries.

The version can be checked with

```shell
❯ fusion version
```

## Merging semantics

Fusion is based on [gookit/config](github.com/gookit/config/v2), itself based on [imdario/mergo](github.com/imdario/mergo). Therefore, merge rules are the one implemented by Mergo.

### Single key/value

Only the last value is retained. For example:

```shell
❯ cat file1.yml
key1: val1

❯ cat file2.yml
key1: val2

❯ fusion yaml file1.yml file2.yml -o res.yml && cat res.yml
key1: val2
```

### Lists

Just like single values, only the last value of conflicting lists is retained. For example:

```shell
❯ cat file1.yml
list1:
  - val11
  - val12

❯ cat file2.yml
list1:
  - val21
  - val22

❯ fusion yaml file1.yml file2.yml -o res.yml && cat res.yml
list1:
  - val21
  - val22
```

### Maps

Maps are merged together, keeping the last value for conflicting keys. For example:

```shell
❯ cat file1.yml
map1:
  key1: val11
  key2: val12

❯ cat file2.yml
map1:
  key1: val21
  key3: val23

❯ fusion yaml file1.yml file2.yml -o res.yml && cat res.yml
map1:
  key1: val21
  key2: val12
  key3: val23
```
