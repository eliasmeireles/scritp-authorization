## Get authorization token

- By request authorization from script `authorization`, the api must return a response like json bellow, by a `POST`
  method declared in `{{apitHost}}`

```json
{
  "code": 200,
  "success": true,
  "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
  "message": "Authorization successful.",
  "timestamp": 1692678504740
}
```

```shell
./ci/authorization -api_host {{apitHost}} -password {{password}} -username {{username}}
```