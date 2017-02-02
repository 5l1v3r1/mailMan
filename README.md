# mailMan
An HTTPS specific internal request router, designed to function as a compact, transparent reverse HTTP proxy.

The following command is an example of using mailMan to redirect HTTPs requests directed at https://localhost:443/server1 to localhost:8080
```
./mailMan "server1://127.0.0.1:8080"
```
