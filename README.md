UserAgent Parser
===============================================

This Caddy Module allows to parse the user agent and provides essential device and bot detection.

You can access the information via these placeholders:
- `{user_agent.is_bot}` - Boolean indicating if the request is from a bot (true/false)
- `{user_agent.device_type}` - Device type string (mobile/tablet/desktop)

The module uses the parser from [here](https://github.com/mileusna/useragent).

## Install

First, the [xcaddy](https://github.com/caddyserver/xcaddy) command:

```shell
$ go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest
```

Then build Caddy with this Go module plugged in. For example:

```shell
$ xcaddy build --with github.com/neodyme-labs/influx_log
```