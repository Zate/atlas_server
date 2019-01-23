#!/bin/bash


export ATLAS_HOME=/home/${USER}/.atlas
export STEAMCMD=~/.steamcmd
export SHOOTERGAME=${ATLAS_HOME}/ShooterGame

cd $STEAMCMD
setfacl -Rm u:10000:rwx ${ATLAS_HOME}
./steamcmd.sh +login anonymous +force_install_dir $ATLAS_HOME +app_update 1006030 validate +quit

if [ -d "$SHOOTERGAME" ]; then
    cd $SHOOTERGAME
    pwd
    if [ ! -d "$SHOOTERGAME/ServerGrid" ]; then
        mkdir -p $SHOOTERGAME/ServerGrid
    fi 
    wget -q https://raw.githubusercontent.com/Zate/atlas_server/master/atlas/ServerGrid.json
    wget -q https://raw.githubusercontent.com/Zate/atlas_server/master/atlas/ServerGrid.ServerOnly.json
    wget -q https://github.com/Zate/atlas_server/raw/master/atlas/ServerGrid/CellImg_0-0.jpg -O $SHOOTERGAME/ServerGrid/CellImg_0-0.jpg
    wget -q https://github.com/Zate/atlas_server/raw/master/atlas/ServerGrid/MapIMG.jpg -O $SHOOTERGAME/ServerGrid/MapIMG.jpg
    setfacl -Rm u:10000:rwx ${ATLAS_HOME}
    exit 1
fi
echo "$SHOOTERGAME missing, check update_server.sh has run correctly."
