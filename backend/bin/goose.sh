#!/bin/bash

# ====================== DEFINE PARAMETERS ===============================
MIGRATION_DIR=/backend/migrations
DRIVER=mysql
# =========================================================================

ONE_ARGS_COMMANDS=("up" "up-by-one" "down" "redo" "reset" "status" "version")
TWO_ARGS_COMMANDS=("up-to" "down-to")


function dburl2dsn () {
    DATABASE_USER=$(echo "$DATABASE_URL" | grep -oP "$DRIVER://\K(.+?):" | cut -d: -f1)
    DATABASE_PASSWORD=$(echo "$DATABASE_URL" | grep -oP "$DRIVER://.*:\K(.+?)@" | cut -d@ -f1)
    DATABASE_HOST=$(echo "$DATABASE_URL" | grep -oP "$DRIVER://.*@\K(.+?):" | cut -d: -f1)
    DATABASE_PORT=$(echo "$DATABASE_URL" | grep -oP "$DRIVER://.*@.*:\K(\d+)/" | cut -d/ -f1)
    DATABASE_NAME_AND_PARAMETER=$(echo "$DATABASE_URL" | grep -oP "$DRIVER://.*@.*:.*/\K(.+?)$")
    if [[ "$DATABASE_NAME_AND_PARAMETER" == *"?"* ]]; then
        DATABASE_NAME=$(echo "$DATABASE_NAME_AND_PARAMETER" | cut -d\? -f1)
        DATABASE_PARAMETER=$(echo "$DATABASE_NAME_AND_PARAMETER" | cut -d\? -f2)
    else
        DATABASE_NAME=$DATABASE_NAME_AND_PARAMETER
        DATABASE_PARAMETER=""
    fi

    if [[ "$DATABASE_PARAMETER" == "" ]]; then
        DATABASE_DSN="$DATABASE_USER:$DATABASE_PASSWORD@tcp($DATABASE_HOST:$DATABASE_PORT)/$DATABASE_NAME"
    else
        DATABASE_DSN="$DATABASE_USER:$DATABASE_PASSWORD@tcp($DATABASE_HOST:$DATABASE_PORT)/$DATABASE_NAME?$DATABASE_PARAMETER"
    fi
    
    echo "$DATABASE_DSN"
    return 0
}

function execute () {
    if [ $# = 1 ]; then
        flag=0
        for i in "${!ONE_ARGS_COMMANDS[@]}"; do
            if [ "${ONE_ARGS_COMMANDS[$i]}" = "$1" ]; then
                goose -dir $MIGRATION_DIR $DRIVER "$DSN" "$1"
                flag=1
            fi
        done
        if [ $flag = 0 ]; then
            echo "invalid command."
            return 2
        fi
    elif [ $# = 2 ]; then
        flag=0
        for i in "${!TWO_ARGS_COMMANDS[@]}"; do
            if [ "${TWO_ARGS_COMMANDS[$i]}" = "$1" ]; then
                goose -dir $MIGRATION_DIR $DRIVER "$DSN" "$1" "$2"
                flag=1
            fi
        done
        if [ $flag = 0 ]; then
            echo "invalid command."
            return 2
        fi
    else
        echo "invalid command."
        return 1
    fi

    return 0
}

DSN=$(dburl2dsn)
execute "$@"
