### Docker Commands
##### This runs a container with a volume binded to ~/tmp/envs. This enables sharing environments by different containers, thus reducing image sizes. It also installs pip for further package addition, upgrade and removal.
```docker run --name taptazi --rm -it -v $HOME/tmp/envs:/home/tazi/miniconda3/envs registry.tazi.ai/tazi-rte:1.0.0 /bin/bash -c "/home/tazi/miniconda3/bin/conda create -y -p /home/tazi/miniconda3/envs/kernel-manager python=3.7.10 pip; bash"```
  <br> 


##### This installs a package to the conda environment. The package will be installed in /miniconda3/envs/<env_name>/lib/python3.7/site-packages/<package_name>
```docker exec taptazi bash -c "/home/tazi/miniconda3/bin/conda init; source /home/tazi/miniconda3/etc/profile.d/conda.sh; conda activate <env_name>; pip install <package_name>"```

<br>

##### This installs a package to the conda environment. The package will be installed in /miniconda3/envs/<env_name>/lib/python3.7/site-packages/<package_name>
```docker exec taptazi bash -c '/home/tazi/miniconda3/bin/conda init; source /home/tazi/miniconda3/etc/profile.d/conda.sh; conda activate <env_name>; pip uninstall -y <package_name>'```

<br>

##### This clones an existing environment.
docker exec taptazi  /bin/bash -c '/home/tazi/miniconda3/bin/conda create -y --name kernel-manager-x --clone kernel-manager'



unzip file to /miniconda3/envs/<env_name>/lib/python3.7/site-packages/<package_name