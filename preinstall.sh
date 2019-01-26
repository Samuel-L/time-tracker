#!/bin/bash
set -e

VERSION="0.5.0"

print_help() {
    echo "Usage: bash preinstall.sh [/path/to/serviceAccountFile.json] [https://{firebase-project-name}.firebaseio.com]"
    echo -e "\nArguments:"
    echo -e "  service account      Add the service account json file path environment variable"
    echo -e "  firebase db url      Add the Firebase Realtime Database url environment variable"
    echo -e "\nOptions:"
    echo -e "  --help               show help"
    echo -e "  --version            print the version"
}

if [ "$1" == "--help" ]; then
    print_help
    exit 0
elif [ "$1" == "--version" ]; then
    echo -e $VERSION
    exit 0
fi

if [ $# -lt 2 ]; then
    print_help
    exit 1
fi

if [ -n "`$SHELL -c 'echo $ZSH_VERSION'`" ]; then
    # assume Zsh
    shell_profile="zshrc"
elif [ -n "`$SHELL -c 'echo $BASH_VERSION'`" ]; then
    # assume Bash
    shell_profile="bashrc"
fi

touch "$HOME/.${shell_profile}"
{
    echo -e "\n# Time Tracker"
    echo -e "export TRACKER_SERVICE_ACCOUNT=$1"
    echo -e "export TRACKER_DB_URL=$2"
} >> "$HOME/.${shell_profile}"

