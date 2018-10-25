## Environment Variables Needed to Use Nexmo

| ENV var name       | Default  | Note  |
|--------------------|----------|-------|
| NEXMO_API_KEY      | _empty_  |       |
| NEXMO_API_SECRET   | _empty_  |       |
| NEXMO_TTL_DELIVERY | 300000   | in ms |
| NEXMO_SENDER       | HEIMDALL |       |

### `send_to` and `check_key` Value

```json
{
  "phone_number": "xxxxxxx",
  "country_code": "xx"
}
```