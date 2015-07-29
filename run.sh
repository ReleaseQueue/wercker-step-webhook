#!/bin/bash

if [ ! -n "$WERCKER_WEBHOOK_URL" ]; then
  echo 'Please specify url property'
  exit 1
fi

./post_to_server "$WERCKER_WEBHOOK_URL"

