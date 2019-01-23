#!/bin/bash
set -e
#"${VARIABLE:=DEFAULT_VALUE}"
MYIP=`wget -qO - https://ifconfig.me`
"${SIP:=MYIP}"
"${MAXPLAYERS:='10'}"
"${GAMEPORT:='27005'}"
"${QUERYPORT:='27015'}"
"${ADMINPASS:='changeme'}"
"${RCON:='True'}"
"${RCONPORT:='27025'}"
"${RESPLAYERS:='0'}"
"${SLOG:='-log'}"
"${ALLHOME:='-ForceAllHomeServer'}"
"${MAP:='Ocean'}"
"${SVRX:='0'}"
"${SVRY:='0'}"
cd /home/steam/.atlas/ShooterGame/Binaries/Linux/
# if [ "$1" = 'update' ]; then
    
# fi

exec ./ShooterGameServer Ocean?ServerX=0?ServerY=0?MaxPlayers=10?Port=27005?QueryPort=27015?SeamlessIP=${SIP}?ServerAdminPassword=changeme?RCONEnabled=True?RCONPort=27025?ReservedPlayerSlots=0 -log -server -ForceAllHomeServer