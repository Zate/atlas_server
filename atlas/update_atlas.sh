#!/bin/bash

export ATLAS_HOME=~/.atlas
export STEAMCMD=~/.steamcmd

cd $STEAMCMD
./steamcmd.sh +login anonymous +force_install_dir $ATLAS_HOME +app_update 1006030 validate +quit