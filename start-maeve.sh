#!/usr/bin/bash

use_lb=true

usage() {
  echo "
  Usage: "$0" [-h] [-n]
  Options:
   -h --help          Returns this helper
   -n --nolb          Disables load balancer
  "
  exit 0;
}

while [ -n "$1" ]; do
    case "$1" in
        -h|--help)
            usage
            ;;
        -n|--nolb)
            use_lb=false
            shift
            ;;
         *)
            echo "Unknown option $1"
            usage
            ;;
    esac
    shift
done

if [[ $use_lb == true ]]; then
	docker compose down
	clear
	docker compose up -d
else
	docker compose -f docker-compose.nolb.yml down --remove-orphans
	clear
	docker compose -f docker-compose.nolb.yml up -d
fi

curl http://localhost:9410/api/v0/token -H 'content-type: application/json' -d '{"countryCode": "US", "partyId": "SANDIA", "contractId": "USEMAC00000004", "uid": "USEMAC00000004", "issuer": "EonTi", "valid": true, "cacheMode": "ALWAYS"}'
curl http://localhost:9410/api/v0/token -H 'content-type: application/json' -d '{"countryCode": "US", "partyId": "SANDIA", "contractId": "USEMAC00000001", "uid": "USEMAC00000001", "issuer": "EonTi", "valid": true, "cacheMode": "ALWAYS"}'
