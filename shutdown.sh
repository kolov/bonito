curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $DO_TOKEN_SARDINE" -d '{"type":"shutdown"}' "https://api.digitalocean.com/v2/droplets/3067649/actions"
