#!/bin/bash

SHOOTERGAME="/home/steam/.atlas/ShooterGame"


if [ -d "$SHOOTERGAME" ]; then
    cd $SHOOTERGAME
    pwd
    if [ ! -d "$SHOOTERGAME/ServerGrid" ]; then
        mkdir -p $SHOOTERGAME/ServerGrid
    fi 
    wget -q https://raw.githubusercontent.com/Zate/atlas_server/master/atlas/ServerGrid.json
    wget -q https://raw.githubusercontent.com/Zate/atlas_server/master/atlas/ServerGrid.ServerOnly.json
    wget -q https://github.com/Zate/atlas_server/raw/master/atlas/ServerGrid/CellImg_0-0.jpg -O $SHOOTERGAME/ServerGrid/CellImg_0-0.jpg
    wget -q https://github.com/Zate/atlas_server/raw/master/atlas/ServerGrid/MapImg.jpg -O $SHOOTERGAME/ServerGrid/MapImg.jpg
    ls
    exit 1
fi
echo "$SHOOTERGAME missing, check update_server.sh has run correctly."