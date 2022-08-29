# Validate CSV
CLI tool to validate CSV by duplicates and sums of rows.

# Overview

The validation is based on two conditions:
- Each row must not contain any duplicate numbers.
  
  Following example will pass validation:
  
  `1,2,3`

  This example is not valid and will not pass validation:
  
  `1,2,2`
- The sums of the rows must be all equal.
  Following example will pass validation:
  
  `1,2,3`

  `3,2,1`

  This example is not valid and will not pass validation:
  
  `1,2,3`

  `3,4,5`


## Installation

The CLI tool validate-csv can be installed in three different ways.

### Go get

The tool can be installed via the `go get` command:

```
go install github.com/robinlieb/validate-csv
```

### Homebrew Package Manager

The tool can be installed via Homebrew: 
```
brew tap robinlieb/brew
brew install validate-csv
```

### Docker

The tool can be installed and run via docker:

```
docker pull robinlieb/validate-csv
docker run -it validate-csv -help
```

## Usage

The CLI tool takes the `validate` command to validate the csv file. The file to specify has to be added via the `--file` flag:

```
validate-csv validate --file example.csv
```

Full overview of available commands can be list with the `-h` flag.
```
 validate-csv -h
 CLI tool to validate CSV by duplicates and sums of rows.
Complete documentation is available at https://github.com/robinlieb/validate-csv

Usage:
  validate-cli [flags]
  validate-cli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  validate    Input file to validate.
  version     Prints the version number.

Flags:
  -h, --help   help for validate-cli

Use "validate-cli [command] --help" for more information about a command.
```