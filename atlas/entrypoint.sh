#!/bin/bash
set -e
if [ "$1" = 'atlas' ]; then
    exec /usr/bin/screen -dmS atlas /home/${USER}/bin/start_server.sh
fi
exec "$@"
