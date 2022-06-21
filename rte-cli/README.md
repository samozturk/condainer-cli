rte-cli is a utility for handling python runtime environment in containers for tazi. 

It allows you to create a conda environment, delete it, clone it and add/remove packages to/from specified environment.

You can customize environment or package using a command line flag.

Usage:
  rte-cli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  environ     Manage conda environments
  help        Help about any command
  package     Manage python packages

Flags:
  -c, --container string     container name
  -e, --envName string       conda environment name
  -h, --help                 help for rte-cli
  -n, --newEnvName string    environment name for cloning a new environment
  -p, --packageName string   package name
  -f, --sourceFile string    path of compressed package directory
  -v, --version              version for rte-cli

Use "rte-cli [command] --help" for more information about a command.