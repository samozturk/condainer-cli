##rte-cli command line tool
rte-cli is a utility for handling python runtime environment in containers for tazi. 

It allows you to **create** a conda environment, **clone** it and **add**/**remove** packages to/from specified environment.

You can customize environment or package using a command line flag.

Usage:
  ```rte-cli [command]```

Available Commands:
> **completion**  Generate the autocompletion script for the specified shell

> **environ**     Manage conda environments

> **help**        Help about any command

> **package**    Manage python packages


Flags:
>-c, --container string     container name

>-e, --envName string       conda environment name

>-h, --help                 help for rte-cli

>-n, --newEnvName string    environment name for cloning a new environment

>-p, --packageName string   package name

>-f, --sourceFile string    path of compressed package directory

>-v, --version              version for rte-cli


Use **"rte-cli [command] --help"** for more information about a command.

### Example Usage ####
#### Environment Related ####
 - **To create an environment in a container:**
Environment name: my_env
Container name: my_cont

```rte-cli environ create -c my_cont -e my_env ```

 - **To clone an environment in a container:**
Environment name: my_env
Container name: my_cont
Clone environment name: my_clone_env

```rte-cli environ clone --container my_cont --envName my_env --newEnvName my_clone_env```

#### Package Related ####
 - **To install a package**
Environment name: my_env
Container name: my_cont
Clone environment name: numpy
``` rte-cli package add -c my_cont -e my_env -p numpy```

 - **To remove a package**
Environment name: my_env
Container name: my_cont
Clone environment name: numpy
``` rte-cli package remove -c my_cont -e my_env -p numpy```

 - **To update a package**
Environment name: my_env
Container name: my_cont
Clone environment name: numpy
``` rte-cli package update -c my_cont -e my_env -p numpy```

 - **To install a package from a compressed file**
Environment name: my_env
Container name: my_cont
Source file: downloads/numpy.zip
``` rte-cli package add -c my_cont -e my_env -f downloads/numpy.zip```