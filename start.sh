docker compose -f docker-compose.nolb.yml down
docker compose -f docker-compose.nolb.yml up -d
curl http://localhost:9410/api/v0/cs/cp001 -H 'content-type: application/json' -d '{"securityProfile": 2}'
curl http://localhost:9410/api/v0/token -H 'content-type: application/json' -d '{"countryCode": "UK", "partyId": "Switch", "contractId": "UKSWI123456789G", "uid": "UKSWI123456789G", "issuer": "Switch", "valid": true, "cacheMode": "ALWAYS"}'
