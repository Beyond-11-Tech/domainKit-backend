# API Documentation
this doc is to serve as a living document for local development explaining the api endpoints available and the params and return objects, forming a contract between this app and the others that link to it.

While every care will be taken to avoid breaking changes, early in development there will be changes that could be breaking, so checking here for updates will be the recommended action.

# authentication
the app uses basic auth for all endpoints currently. the password of which will be set when the app starts (in prod that means having a shared secret across the repos). the username will be `web` for the web user, and `app` for the app user.

there is currently no endpoint to sign in, so the basic auth is to be sent as part of the headers in every request

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

### /v1/domain/a
a basic search endpoint, returns the AAAA (IPv6) records for a given domain.

it will search all DNS providers for the specified domain, and return any AAAA records in the result

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
      "2606:4700:3030::6815:3001",
      "2606:4700:3030::6815:7001",
      "2606:4700:3030::6815:1001",
      "2606:4700:3030::6815:5001",
      "2606:4700:3030::6815:2001",
      "2606:4700:3030::6815:6001",
      "2606:4700:3030::6815:4001"
    ]
  },
  {
    "registrar": "8.8.8.8",
    "record": [
      "2606:4700:3030::6815:1001",
      "2606:4700:3030::6815:4001",
      "2606:4700:3030::6815:7001",
      "2606:4700:3030::6815:5001",
      "2606:4700:3030::6815:2001",
      "2606:4700:3030::6815:3001",
      "2606:4700:3030::6815:6001"
    ]
  }
]
```

### /v1/domain/ns
a basic search endpoint, returns the NS (NameServer) records for a given domain.

it will search all DNS providers for the specified domain, and return the found records as the result

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
      "ns2.siteground.net.",
      "ns1.siteground.net."
    ]
  },
  {
    "registrar": "8.8.8.8",
    "record": [
      "ns2.siteground.net.",
      "ns1.siteground.net."
    ]
  }
]
```