#!/bin/bash


export ATLAS_HOME=/home/${USER}/.atlas
export REDIS_DATA=${ATLAS_HOME}/data
export STEAMCMD=~/.steamcmd
export SHOOTERGAME=${ATLAS_HOME}/ShooterGame

rm -rf /home/${USER}/Steam/appcache/appinfo.vdf

LOCALBID=`cat ${ATLAS_HOME}/steamapps/appmanifest_1006030.acf | grep buildid | cut -d'"' -f4`
#        "buildid"               "3497453"
        
REMOTEBID=`${STEAMCMD}/steamcmd.sh +login anonymous +app_info_update 1 +app_info_print 1006030 +quit | grep buildid | cut -d'"' -f4`
#                                "buildid"               "3497453"
if [ "$LOCALBID" == "$REMOTEBID" ]; then
    echo "No update available - L : ${LOCALBID} R: ${REMOTEBID}"
    exit 1
fi

echo "Update Available .. New BID : ${REMOTEBID}"
