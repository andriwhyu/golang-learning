#!/bin/bash

# Set interval (in seconds) between checks
INTERVAL=1

while true; do
  # Check connection to Redis
#  response=$(redis-cli -h "$REDIS_HOST" -p "$REDIS_PORT" PING)
#
#  # Output result
#  if [[ "$response" == "PONG" ]]; then
#    echo "$(date): Redis is up"
#  else
#    echo "$(date): Failed to connect to Redis"
#  fi

  python3 socket-libc.py "$1"

  # Wait for the next check
  sleep "$INTERVAL"
done