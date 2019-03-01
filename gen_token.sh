#!/bin/bash
# Request username and password for connecting to Taiga
read -p "Username or email: " USERNAME
read -r -s -p "Password: " PASSWORD
read -p "Taiga domain: " TAIGA_DOMAIN

DATA=$(jq --null-input \
        --arg username "$USERNAME" \
        --arg password "$PASSWORD" \
        '{ type: "normal", username: $username, password: $password }')

# Get AUTH_TOKEN
USER_AUTH_DETAIL=$( curl -X POST \
  -H "Content-Type: application/json" \
  -d "$DATA" \
  https://$TAIGA_DOMAIN/api/v1/auth 2>/dev/null )

echo $USER_AUTH_DETAIL

AUTH_TOKEN=$( echo ${USER_AUTH_DETAIL} | jq -r '.auth_token' )

# Exit if AUTH_TOKEN is not available
if [ -z ${AUTH_TOKEN} ]; then
    echo "Error: Incorrect username and/or password supplied"
    exit 1
else
    echo "auth_token is ${AUTH_TOKEN}"
fi

