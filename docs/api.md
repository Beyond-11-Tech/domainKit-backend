# API Documentation
this doc is to serve as a living document for local development explaining the api endpoints available and the params and return objects, forming a contract between this app and the others that link to it.

While every care will be taken to avoid breaking changes, early in development there will be changes that could be breaking, so checking here for updates will be the recommended action.

# endpoints

### /v1/domain/a
a basic search endpoint, returns the A (IPv4) records for a given domain.

it will search all DNS providers for the specified domain, and return the A records as the result

#### inputs

Path params (Req)

|name| type | example |
|----|------|---------|
address | string | google.com

#### output

array of object with each object a single query to a DNS registrar

```json
[
  {
    "registrar": "1.1.1.1",
    "record": [
      "17.253.144.10"
    ]
  },
  {
    "registrar": "8.8.8.8",
    "record": [
      "17.253.144.10"
    ]
  }
]
```