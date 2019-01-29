#!/bin/bash


export ATLAS_HOME=/home/${USER}/.atlas
export REDIS_DATA=${ATLAS_HOME}/data
export STEAMCMD=~/.steamcmd
export SHOOTERGAME=${ATLAS_HOME}/ShooterGame

rm -rf /home/${USER}/Steam/appcache/appinfo.vdf

LOCALBID=`cat ${ATLAS_HOME}/steamapps/appmanifest_1006030.acf | grep buildid | cut -d'"' -f4`
#        "buildid"               "3497453"
        
REMOTEBID=`${STEAMCMD}/steamcmd.sh +app_info_update 1 +app_info_print 1006030 +quit | grep buildid | cut -d'"' -f4`
#                                "buildid"               "3497453"
if [ "$LOCALBID" == "$REMOTEBID" ]; then
    echo "No update available"
    exit 1
fi

echo "Update Availablr .."

cd $STEAMCMD
sudo setfacl -Rm u:10000:rwx ${ATLAS_HOME}
./steamcmd.sh +login anonymous +force_install_dir $ATLAS_HOME +app_update 1006030 validate +quit
if [ -d "$REDIS_DATA" ]; then
    mkdir -p $REDIS_DATA
fi
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
    sudo setfacl -Rm u:10000:rwx ${ATLAS_HOME}
    exit 1
fi
echo "$SHOOTERGAME missing, check update_server.sh has run correctly."
