# config.json examples

## Reality / TLS

Assuming local proxy is socks5://127.0.0.1:1080 and url-encoded vpn config from your provider is vless://`aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee`@`123.123.123.123`:**443**?type=**tcp**&security=reality&encryption=none&flow=xtls-rprx-vision&fp=`browser`&sni=`some.site.com`&sid=&pbk=`KeyKeyKeyKeyKeyKeyKeyKey`#SomeName

```
{
    "log": {
        "disabled": true
    },
    "inbounds": [
        {
            "type": "socks",
            "tag": "socks-in",
            "listen": "127.0.0.1",
            "listen_port": 1080,
            "users": []
        }
    ],
    "outbounds": [
        {
            "type": "vless",
            "tag": "vless-out",
            "server": "123.123.123.123",
            "server_port": 443,
            "uuid": "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
            "flow": "xtls-rprx-vision",
            "tls": {
                    "enabled": true,
                    "server_name": "some.site.com",
                    "utls": {
                    "enabled": true,
                    "fingerprint": "browser"
                },
                "reality": {
                    "enabled": true,
                    "public_key": "KeyKeyKeyKeyKeyKeyKeyKey"
                }
            },
            "packet_encoding": "xudp"
        }
    ],
    "route": {
        "rules": [
            {
                "inbound": "socks-in",
                "outbound": "vless-out"
            }
        ]
    }
  }
```

## xhttp

Assuming local proxy is socks5://127.0.0.1:1080 and url-encoded vpn config from your provider is vless://`aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee`@`123.123.123.123`:**8443**?type=**xhttp**&security=reality&encryption=none&fp=`browser`&sni=`some.site.com`&sid=&pbk=`KeyKeyKeyKeyKeyKeyKeyKey`&path=`%2F`#SomeName

```
{
    "log": {
        "disabled": true
    },
    "inbounds": [
        {
            "type": "socks",
            "tag": "socks-in",
            "listen": "127.0.0.1",
            "listen_port": 1080,
            "users": []
        }
    ],
    "outbounds": [
        {
            "type": "vless",
            "tag": "vless-out",
            "server": "123.123.123.123",
            "server_port": 8443,
            "uuid": "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
            "tls": {
                "enabled": true,
                "server_name": "some.site.com",
                "alpn": [
                    "h2",
                    "http/1.1"
                ],
                "reality": {
                    "enabled": true,
                    "public_key": "KeyKeyKeyKeyKeyKeyKeyKey"
                },
                "utls": {
                    "enabled": true,
                    "fingerprint": "browser"
                }
            },
            "transport": {
                "type": "xhttp",
                "path": "/",
                "mode": "auto",
                "x_padding_bytes": "100-1000",
                "host": "some.site.com"
            }
        }
    ],
    "route": {
        "rules": [
            {
                "inbound": "socks-in",
                "outbound": "vless-out"
            }
        ]
    }
}
```
