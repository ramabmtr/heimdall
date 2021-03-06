# Heimdall

<p align="center"><img src="doc-asset/heimdall.jpg" width="360"></p>

A flexible service to handle OTP

## ENV variable

| ENV var name    | Default     | Note                                              |
|-----------------|-------------|---------------------------------------------------|
| APP_ADDRESS     | :1323       |                                                   |
| APP_DEBUG       | 1           | available value: `0` or `1`                       |
| APP_MODE        | development | available value: `development` or `production`    |
| DATABASE_TYPE   | redis       | available value: `redis` or `memcached`           |
| SERVICE_TYPE    | _empty_     | available value: `twilio`, `nexmo`, or `postmark` |
| OTP_EXPIRY_TIME | 5           | in minutes                                        |

## Requirements

Currently we support this component to send OTP:

#### Database

You can use database below and config them using environment variable (`DATABASE_TYPE`)

- Redis (`DATABASE_TYPE="redis"`, recommended)

  Recommended for multi machine. To use this DB, you must set up env variable below:
  
  | ENV var name   | Default        | Note |
  |----------------|----------------|------|
  | REDIS_URL      | localhost:6379 |      |
  | REDIS_PASSWORD | _empty_        |      |
  | REDIS_DB       | 0              |      |

- Memcached (`DATABASE_TYPE="redis"`)

  To use this DB, you must set up env variable below:
  
  | ENV var name               | Default | Note  |
  |----------------------------|---------|-------|
  | MEMCACHED_CLEANUP_INTERVAL | 600000  | in ms |


#### 3rd Party Service

You can use this service below to send OTP. Configurable via payload in API call

- [Twilio](repository/external/twilio)
- [Nexmo](repository/external/nexmo)
- [Postmark](repository/external/postmark)

### [Add more component?](COMPONENT.md)

## Endpoint

#### GET /v1/ping

#### POST /v1/verification/send

body:

```json
{
  "send_to": "<service send_to>"
}
```

**NOTE:**

- For `service send_to`, you can look in each link provided in [this section](#3rd-party-service)

return:

```json
{
  "status": "success"
}
```

#### POST /v1/verification/check

body:

```json
{
  "code": "<verification code>",
  "check_key": "<service check_key>"
}
```

**NOTE:**

- For `service check_key`, you can look in each link provided in [this section](#3rd-party-service)

return:

```json
{
  "status": "success"
}
```