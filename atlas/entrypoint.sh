#!/bin/bash
set -e
if [ "$1" = 'atlas' ]; then
    exec "screen -dmS atlas -L -U /home/steam/bin/start_server.sh"
    exec "tail -f /dev/null"
fi
exec "$@"
