#!/usr/bin/env bash

function usage() {
    echo >&2
    echo "Simple script to generate sample file.">&2
    echo >&2
    echo "  Usage:" >&2
    echo "    generate_sample.sh [OPTIONS] <ARGUMENTS>" >&2
    echo >&2
    echo "  Arguments:" >&2  
    echo "      -n|--number                  Number of sample lines to generate. Defaults to 50.">&2  
    echo >&2
    echo "  Options:" >&2  
    echo "      -h|--help                     Display help.">&2 
    echo >&2
}

while [ "$1" != "" ];do
    case $1 in
    -n | --number)
        shift
        NUM_SAMPLES=$1
        ;;
    -h | --help)
        usage
        exit 0
        ;;
    *)
        echo "INVALID FLAG"
        usage
        exit 1
        ;;
    esac
    shift
done

if [ -z "$NUM_SAMPLES" ]; then
   NUM_SAMPLES=50
fi

for (( i=$NUM_SAMPLES ; i>0 ; i-- )); do
   echo "http://api.tech.com/item/${RANDOM} ${i}" >> tmp_sample.txt
done 
