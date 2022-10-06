# cas-oidc-dogu

Dogu for a simple oidc client. Used for testing purposes only

### Local development

1. Start CAS locally via `.gradlew build run`
2. Start client `go run .`

### Testing CAS-OIDC via Dogu

1. Build and start dogu with `cesapp build .` and `cesapp start cas-oidc-dogu`.
2. Visit `<fqdn>/cas-oidc-dogu/login`.
3. Browser should recieve profile as JSON:

```json
{
  "attributes": {
    "cn": "admin",
    "displayName": "admin",
    "givenName": "admin",
    "groups": [
      "cesManager",
      "asd"
    ],
    "mail": "asd@asd.asd",
    "surname": "admin",
    "username": "asd"
  },
  "auth_time": 1665040091,
  "client_id": "cas-oidc-client",
  "id": "asd",
  "service": "https://192.168.56.2/cas-oidc-client/auth/callback",
  "sub": "asd"
}
```