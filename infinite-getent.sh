#!/bin/bash

HOSTNAME=$1
INTERVAL=1

while true; do
  # Perform DNS resolution using getent, which uses libc
  IP_ADDRESS=$(getent hosts "$HOSTNAME" | awk '{ print $1 }')

  # Check if we got an IP address
  if [ -n "$IP_ADDRESS" ]; then
      echo "The IP address of $HOSTNAME is: $IP_ADDRESS"
  else
      echo "Failed to resolve $HOSTNAME"
  fi

  # Wait for the next check
  sleep "$INTERVAL"
done