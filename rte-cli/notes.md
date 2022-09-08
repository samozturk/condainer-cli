### Docker Commands
##### This runs a container with a volume binded to ~/tmp/envs. This enables sharing environments by different containers, thus reducing image sizes. It also installs pip for further package addition, upgrade and removal.
```docker run --name taptazi --rm -it -v $HOME/tmp/envs:/home/tazi/miniconda3/envs registry.tazi.ai/tazi-rte:1.0.0 /bin/bash -c "/home/tazi/miniconda3/bin/conda create -y -p /home/tazi/miniconda3/envs/kernel-manager python=3.7.10 pip; bash"```
  <br> 


##### This installs a package to the conda environment. The package will be installed in /miniconda3/envs/<env_name>/lib/python3.7/site-packages/<package_name>
```docker exec taptazi bash -c "/home/tazi/miniconda3/bin/conda init; source /home/tazi/miniconda3/etc/profile.d/conda.sh; conda activate <env_name>; pip install <package_name>"```

<br>

##### This installs a package to the conda environment. The package will be installed in /miniconda3/envs/<env_name>/lib/python3.7/site-packages/<package_name>
```docker exec taptazi bash -c '/home/tazi/miniconda3/bin/conda init'; source /home/tazi/miniconda3/etc/profile.d/conda.sh; conda activate <env_name>; pip uninstall -y <package_name>'```

<br>

##### This clones an existing environment.
```docker exec taptazi  /bin/bash -c '/home/tazi/miniconda3/bin/conda create -y --name kernel-manager-x --clone kernel-manager'```



unzip file to /miniconda3/envs/<env_name>/lib/python3.7/site-packages/<package_name



docker exec tazitest bash -c '/home/tazi/miniconda3/bin/conda init'; source /home/tazi/miniconda3/etc/profile.d/conda.sh; conda env list






testing:

docker run --rm -it -v $HOME/tmp/envs:/home/tazi/miniconda3/envs --name tazitest registry.tazi.ai/tazi-rte:1.0.0 /bin/bash -c '/home/tazi/miniconda3/bin/conda init';



### Python packages
Python packages can be installed manually. For that, wheel files needed which can be found at https://pypi.org/simple/<package_name>
Simply find appropriate version and architecture.

Or pip can do it for you. ```python -m pip download --only-binary :all: --dest . --no-cache <package_name> ```
Downloaded file can tarbal or wheel. If that's a tarball, untar it and find the setup.py
in setup.py, you can dependency files with regex ```install_requires=\[:*(.*?)\]```

If that's a whl; first unzip the whl file. Whl files are pretty much like zip files. It will yield two folders: one is named as <package_name> other is <{package_name}-{version}.dist-info>. Inside the latter, there is a file called METADATA.
You can parse neccessary libraries from METADA using this regex ```Requires-Dist:\s:*(.*?)\n```.

All dependencies needs to be installed before installing the package. Also you need to install dependencies of dependencies and so on.