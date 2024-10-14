Client app to test a transaction with a server using custom CA certificate with Ko/Docker

# Run

```sh
docker run --rm --network=host -e SSL_CERT_DIR=/var/run/ko $(ko build --local)
```
