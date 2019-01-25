#!/bin/bash


export ATLAS_HOME=/home/${USER}/.atlas
export REDIS_DATA=${ATLAS_HOME}/data
export STEAMCMD=~/.steamcmd
export SHOOTERGAME=${ATLAS_HOME}/ShooterGame

LOCALBID=`cat ${ATLAS_HOME}/steamapps/appmanifest_1006030.acf | grep buildid | cut -d'"' -f4`
#        "buildid"               "3497453"
        
REMOTEBID=`${STEAMCMD}/steamcmd.sh +app_info_update 1 +app_info_print 1006030 +quit | grep buildid | cut -d'"' -f4`
#                                "buildid"               "3497453"
if [ "$LOCALBID" == "$REMOTEBID" ]; then
    echo "No update available"
    exit 1
fi

echo "Update Availablr .."
