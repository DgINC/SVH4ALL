#!/usr/bin/env bash

function usage()
{
    echo "Build scrpit! (fcking qt bind...)"
}

POSITIONAL=()
while [[ $# -gt 0 ]]
do
key="$1"

case $key in
    -b | --build)
        TARGET="$2"
        shift
        shift
        ;;
    -h | --help)
        usage
        exit
        ;;
    *)
    POSITIONAL+=("$1")
        echo "ERROR: unknown parameter ${POSITIONAL[$@]}"
        echo "Type '-h|--help' for information."
        shift
        exit 1
        ;;
esac
done
set -- "${POSITIONAL[@]}" # restore positional parameters
