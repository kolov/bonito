#!/usr/bin/env bash

curl -X POST \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer $DO_TOKEN_SARDINE" \
     -d '{"name":"sardine.com",
          "region": "fra1",
          "size": "2gb",
          "image": "21403067",
          "ssh_keys": ["763745"],
          "backups": false,
          "ipv6": false,
          "user_data": null,
          "private_networking": null,
          "volumes": null,
          "tags":["work"]}' \
      "https://api.digitalocean.com/v2/droplets"
