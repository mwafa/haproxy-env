# HAPROXY-ENV

Simple way to create haproxy based reverse proxy from env.
This method is convenient when you using docker-compose on swarm mode.

### Simple usage:

```yml
version: "3.8"

services:
  proxy:
    image: mwafa/haproxy-env:latest
    ports:
      - "80:80"
    environment:
      - BIND_APP1=domain1.com 10.0.0.1:1234
      - BIND_APP2=domain2.com 10.0.0.2:3000 10.0.0.3:3000
      - ...

    labels:
      - traefik.enable=true
      - traefik.http.routers.haproxy.rule=Host(`domain1.com`) || Host(`domain2.com`)
      - traefik.http.routers.haproxy.entrypoints=websecure
      - traefik.http.routers.haproxy.tls.certresolver=le
      - traefik.http.services.haproxy.loadbalancer.server.port=80
```

Just add

```sh
BIND_<name>=<from_domain> <target_1> <target2> ...
```
