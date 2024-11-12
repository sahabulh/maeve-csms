docker compose -f docker-compose.nolb.yml down
docker compose -f docker-compose.nolb.yml up -d
# curl http://localhost:9410/api/v0/cs/cp001 -H 'content-type: application/json' -d '{"securityProfile": 0, "base64SHA256Password": "3oGi4B5I+Y9iEkYtL7xvuUxrvGOXM/X2LQrsCwf/knA="}'
curl http://localhost:9410/api/v0/cs/cp001 -H 'content-type: application/json' -d '{"securityProfile": 2}'
curl http://localhost:9410/api/v0/token -H 'content-type: application/json' -d '{"countryCode": "US", "partyId": "Sandia National Labs", "contractId": "USCPIC001LTON3", "uid": "USCPIC001LTON3", "issuer": "Sandia National Labs", "valid": true, "cacheMode": "ALWAYS"}'
curl http://localhost:9410/api/v0/token -H 'content-type: application/json' -d '{"countryCode": "US", "partyId": "Sandia National Labs", "contractId": "UKSWI123456789G", "uid": "UKSWI123456789G", "issuer": "Sandia National Labs", "valid": true, "cacheMode": "ALWAYS"}'
