# fusion

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
