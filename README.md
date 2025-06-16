UserAgent Parser
===============================================

This Caddy Module allows to parse the user agent and provides essential device and bot detection.

You can access the information via this placeholder:
- `{user_agent.type}` - Type string with possible values: `bot`, `mobile`, `tablet`, `desktop`

Priority: Bot detection has priority over device type. If a bot is detected, the type will be `bot` regardless of the device.

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