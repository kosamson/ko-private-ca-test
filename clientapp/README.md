# Run client app to test custom CA certificate with Ko/Docker

```sh
docker run --rm --network=host -e SSL_CERT_DIR=/var/run/ko $(ko build --local)
```
