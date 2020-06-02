Sample Application
==================

A sample docker-compose file is provided here for you.

## Quick start

### Without SSL

```
docker-compose -p PCR -f up -d
```

### With SSL

If you'd like to run the application with SSL, execute the following 
commands first

```bash
mkdir -p certs
cd certs

# create private key
openssl ecparam -genkey -name secp384r1 -out server.key

# create public key
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

docker-compose -p PCR -f docker-compose.ssl.yml up -d
```

## Things to Note

Note that if you run the server in a dockerized environment, you must 
set the environment variable `PCR_INDEXER.TYPE` to `shell`. Do note
that you must use upper-case names for environment variable keys.
