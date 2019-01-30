#!/bin/bash

# This is going to be dirty and functional, not going to win any awards for doing it right, or pretty.

# Check if we are on Bionic, error out if not.  (people can comment this out themselves to allow other OS's)

function checkOS {
    rel=`lsb_release -cs`
    if [ ! "$rel" == "bionic" ]; then
        echo "This script is made for use on Ubuntu 18.04 Bionic"
        echo "It may work on other Debian based distros, comment out"
        echo "checkOS as indicated in the script to try it"
        exit 1
    fi

}

#comment out the checkOS call below this line to use on something other than bionic
#checkOS

# Check if we have docker all setup and working right as this user
# if we do notm prompt user to see if they want us to install it
# if they do not, lets exit, this whole thing is built around using docker to 
# abstract the game server from anything on your main OS so that anything we
# do to make the server work, will not affect your main OS permanently.

function checkDocker {
    if [ ! -x "$(command -v docker)" ]; then
        echo "Docker is missing, would you like to instsall it? (Y/N)

        sudo apt-get install \
        apt-transport-https \
        ca-certificates \
        curl \
        gnupg2 \
        software-properties-common

        curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add -

        sudo add-apt-repository \
        "deb [arch=amd64] https://download.docker.com/linux/debian \
        $(lsb_release -cs) \
        stable"

        sudo apt-get update
        sudo apt-get install docker-ce -y

        sudo usermod -aG docker zate
        # Ask if we should install it, and insgtall it if they say yes.
        # Exit if they say no
    fi
    docker -v
}

checkDocker
