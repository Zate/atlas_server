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
FROM golang:latest as websocket
ENV DEBIAN_FRONTEND=noninteractive
ARG TZ="Etc/UTC"
RUN apt-get update  && apt-get dist-upgrade -y &&\
        apt-get install -y git && \
        cd /tmp && \
        git clone https://github.com/Zate/websocketd.git && \
        cd websocketd && \
        go build && \
        cd /tmp && \
        git clone https://github.com/Zate/web-vmstats.git
FROM ubuntu:18.04 as steamcmdbase
ENV DEBIAN_FRONTEND=noninteractive
ARG TZ="Etc/UTC"
COPY --from=websocket /tmp/websocketd/websocketd /usr/local/bin
COPY --from=redis-cli /tmp/redis-stable/src/redis-cli /usr/local/bin/
COPY entrypoint.sh /usr/local/bin/
RUN apt-get update  && apt-get dist-upgrade -y &&\
        apt-get install -y iputils-ping glances screen wget unzip vim-tiny bzip2 jq iproute2 lib32stdc++6 lib32gcc1 software-properties-common apt-transport-https libcurl3 libcurl-openssl1.0-dev && \
        apt-get clean &&\
        rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* && \
        echo "$TZ" > /etc/timezone && \
        ln -fs /usr/share/zoneinfo/$TZ /etc/localtime &&\
	chmod +x /usr/local/bin/entrypoint.sh && \
        useradd -r -m -u 10000 -s /bin/bash steam && \
        mkdir -p /home/steam/web/js && \
        chown -R steam:steam /home/steam/web
COPY --from=websocket --chown=steam:steam /tmp/web-vmstats/web/js /home/steam/web/js
COPY --from=websocket --chown=steam:steam /tmp/web-vmstats/web/*.* /home/steam/web/
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
USER root
RUN mkdir -p /home/steam/.atlas && \
        chown -R steam:steam /home/steam/.atlas && \
        chown -R steam:steam /home/steam/web && \
        chmod +x /home/steam/bin/*.sh
#        ./steamcmd.sh +login anonymous +force_install_dir ~/.atlas +app_update 1006030 validate +quit\
#ENTRYPOINT ["sh","/usr/local/bin/entrypoint.sh"]
USER steam
WORKDIR /home/steam/.steamcmd
VOLUME /home/steam/.atlas
CMD tail -f /dev/null
