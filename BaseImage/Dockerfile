FROM centos:8

# ARG only works in build time. ENV can work for container.
ARG TZ=UTC

RUN cd /etc/yum.repos.d/
RUN sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-*
RUN sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-*

RUN dnf upgrade -y 
RUN dnf -y install epel-release
RUN dnf upgrade -y
RUN dnf -y install sqlite redis

RUN yum -y install wget vim jq bc uuid zip unzip java-1.8.0-openjdk-devel openssl-devel
RUN yum -y groupinstall 'Development Tools'

RUN useradd -ms /bin/bash tazi
USER tazi
WORKDIR /home/tazi

RUN curl https://repo.anaconda.com/miniconda/Miniconda3-py37_4.12.0-Linux-x86_64.sh -o miniconda.sh
RUN chmod +x miniconda.sh
RUN bash miniconda.sh -b -p $HOME/miniconda3
RUN eval "$($HOME/miniconda3/bin/conda shell.bash hook)" && conda init

