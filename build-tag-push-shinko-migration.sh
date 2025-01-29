#!/bin/zsh

podman build -t shinko-db-migration -f sql/Containerfile
podman tag shinko-db-migration quay.io/joeavaik/shinko-db-migration
podman push quay.io/joeavaik/shinko-db-migration
podman image prune -a
