#!/bin/bash

ME=`whoami`
export ATLAS_HOME=${me}/.atlas
export STEAMCMD=~/.steamcmd

cd $STEAMCMD
./steamcmd.sh +login anonymous +force_install_dir $ATLAS_HOME +app_update 1006030 validate +quit
setfacl -Rm u:10000:rwx ${ATLAS_HOME}