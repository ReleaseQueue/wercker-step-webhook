#!/bin/bash

if [ ! -n "$WERCKER_WEBHOOK_URL" ]; then
  echo 'Please specify url property'
  exit 1
fi

SCRIPT_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
$SCRIPT_PATH/post_to_server "$WERCKER_WEBHOOK_URL"

