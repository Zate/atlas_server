FROM ubuntu:18.04 as redis-cli
ARG TZ="Etc/UTC"
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update  && apt-get dist-upgrade -y &&\
        apt-get install -y wget build-essential tcl libjemalloc-dev libjemalloc1 && \
	cd /tmp && \
	wget -q http://download.redis.io/redis-stable.tar.gz && \
	tar xvzf redis-stable.tar.gz && \
	cd redis-stable/src && \
	make redis-cli
FROM ubuntu:18.04 as steamcmdbase
ENV DEBIAN_FRONTEND=noninteractive
ARG TZ="Etc/UTC"
COPY --from=redis-cli /tmp/redis-stable/src/redis-cli /usr/local/bin/
COPY entrypoint.sh /usr/local/bin/
RUN apt-get update  && apt-get dist-upgrade -y &&\
        apt-get install -y iputils-ping wget unzip vim-tiny bzip2 jq iproute2 lib32stdc++6 lib32gcc1 software-properties-common apt-transport-https libcurl3 libcurl-openssl1.0-dev && \
        apt-get clean &&\
        rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* && \
        echo "$TZ" > /etc/timezone &&\
        ln -fs /usr/share/zoneinfo/$TZ /etc/localtime &&\
        useradd -r -m -s /bin/bash steam
USER steam
RUN mkdir -p ~/bin ~/.steamcmd  ~/.steam/sdk32 &&\
        export PATH="~/bin:$PATH" && \
        cd ~/.steamcmd &&\
        wget -q https://steamcdn-a.akamaihd.net/client/installer/steamcmd_linux.tar.gz && \
        tar -xf steamcmd_linux.tar.gz && \
        rm steamcmd_linux.tar.gz && \
        ln -s ~/.steamcmd/linux32/steamclient.so ~/.steam/sdk32/steamclient.so && \
        ./steamcmd.sh +login anonymous +quit
COPY --chown=steam:steam start_server.sh update_atlas.sh update_map.sh /home/steam/bin/
FROM steamcmdbase as atlas-s
USER steam
WORKDIR /home/steam/.steamcmd
RUN mkdir -p ~/.atlas && \
        chmod +x /home/steam/bin/*.sh
#        ./steamcmd.sh +login anonymous +force_install_dir ~/.atlas +app_update 1006030 validate +quit\
#ENTRYPOINT ["sh","/usr/local/bin/entrypoint.sh"]
CMD tail -f /dev/null
