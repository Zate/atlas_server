#!/bin/bash

#IP=`wget -qO - https://ifconfig.me`
IP="159.203.52.169"

cd /home/steam/.atlas/ShooterGame/Binaries/Linux/
#./ShooterGameServer Ocean?RCONEnabled=True?RCONPort=32332?SessionName=1x0?Port=5751?QueryPort=57552?ServerPassword?ServerAdminPassword=passwordgoeshere?MaxPlayers=10?ServerX=1?ServerY=0?SeamlessIP=63.224.241.83?culture=en?MapPlayerLocation=True?AltSaveDirectoryName=A2?listen -log -server -NoCrashDialog -BattlEye


./ShooterGameServer "Ocean?ServerX=0?ServerY=0?AltSaveDirectoryName=A1?MaxPlayers=10?ReservedPlayerSlots=0?SeamlessIP=$IP?Port=5701?QueryPort=57501?RconPassword=test -NoBattlEye -log -server"
