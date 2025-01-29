#!/bin/zsh

podman build -t shinko-app -f Containerfile
podman tag shinko-app quay.io/joeavaik/shinko-app 
podman push quay.io/joeavaik/shinko-app
podman image prune -a
