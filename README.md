
<h1 align="center">
rte-cli Command Line Tool
</h1>

<p align="center">
<img src="resources/rte-cli-logo.png" alt="rte-cli logo" width="200"/>
</p>

rte-cli is a utility for handling python runtime environment in containers for tazi. It has features to manage conda environments, python packages etc.

**<u>Environments</u>**:
 It enables you to **create** a conda environment, **clone** it and **add**/**remove** packages to/from specified environment.
**<u>Packages</u>**:
It enables you to **install, remove, update** a package or **export environment packages** and **install for offline use**.


You can customize environment or package using a command line flag.

_Note: No matter in which container you install the package, all environments are shared between containers using mounting a volume binding._
<br>

#### Version 2.0
- More verbose errors
- More reliable subcommands
- Container home path(WORKDIR) is now customizable; not fixed to "/home/tazi".
- Volume binding path is now customizable, not fixed to "/tmp/envs/"
- You can now download python packages to your local environment. Then you can use them in an offline environment with rte-cli.(Only for macOs and Linux users)

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

>- , --version              version for rte-cli

>-r, --requirementsFile     Path of requirements.txt to be installed

>-v , --pythonVersion       Python version of the environment


Use **"rte-cli [command] --help"** for more information about a command.

### Example Usage ####
#### Environment Related ####
 - **To create an environment in a container:**
Environment name: my_env
Container name: my_cont
Python version: 3.8.5

```rte-cli environ create -c my_cont -e my_env -v 3.8.5 ```

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
``` rte-cli package addZip -c my_cont -e my_env -f downloads/numpy.zip```

 - **To install multiple packages from a requirements.txt file**
Environment name: my_env
Container name: my_cont
Source file: downloads/requirements.txt
``` rte-cli package installReq -c my_cont -e my_env -r downloads/requirements.txt```

#### Run Related ####
 - **To run a python script**
Environment name: my_env
Container name: my_cont
Script name: script.py
``` rte-cli run -c my_cont -e my_env -p script.py```