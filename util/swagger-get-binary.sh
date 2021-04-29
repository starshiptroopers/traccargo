#!/bin/bash
#jq is needed
#on debian/ubuntu install it with: sudo apt-get install jq

download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | \
  jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url')
curl -o ./swagger -L'#' "$download_url"
chmod +x ./swagger
