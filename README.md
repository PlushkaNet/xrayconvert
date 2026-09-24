Small http-server with libxray inside for parsing VLESS://, VMESS://, TROJAN://, SS:// URI's and return JSON-correct outbounds for it.

Usage:
```
./xrayconvert # will start on 127.0.0.1:10100 by default
```

Or to specify address explicitly:
```
./xrayconvert -address 127.0.0.1:4000
```

> [!IMPORTANT]
> This server does *not* support HTTPS and intented to use only in local networks or behind proxy

Licensed under the MIT-0 license, read the LICENSE file for more details