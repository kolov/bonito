curl -X GET -H "Content-Type: application/json" \
     -H "Authorization: Bearer $DO_TOKEN_SARDINE" \
     "https://api.digitalocean.com/v2/images?page=1&per_page=100&private=true" | jq
