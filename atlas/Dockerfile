FROM ubuntu:18.04 as steamcmdbase
ARG TZ="Etc/UTC"
RUN apt-get update  && apt-get dist-upgrade -y &&\
        apt-get install -y wget unzip vim-tiny bzip2 jq iproute2 lib32stdc++6 lib32gcc1 software-properties-common apt-transport-https libcurl3 libcurl-openssl1.0-dev && \
        apt-get clean &&\
        rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* && \
        echo "$TZ" > /etc/timezone &&\
        ln -fs /usr/share/zoneinfo/$TZ /etc/localtime &&\
        useradd -r -m -s /bin/bash steam
USER steam
RUN mkdir -p ~/bin ~/.steamcmd  ~/.steam/sdk32 &&\
        cd ~/.steamcmd &&\
        wget -q https://steamcdn-a.akamaihd.net/client/installer/steamcmd_linux.tar.gz && \
        tar -xf steamcmd_linux.tar.gz && \
        rm steamcmd_linux.tar.gz && \
        ln -s ~/.steamcmd/linux32/steamclient.so ~/.steam/sdk32/steamclient.so && \
        ./steamcmd.sh +login anonymous +quit
COPY update_atlas.sh ~/bin
FROM steamcmdbase as atlas-s
USER steam
WORKDIR /home/steam/.steamcmd
RUN mkdir -p ~/.atlas
#        ./steamcmd.sh +login anonymous +force_install_dir ~/.atlas +app_update 1006030 validate +quit
