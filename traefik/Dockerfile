FROM traefik:latest as atlas-t

COPY traefik.toml /etc/traefik
COPY acme.json /etc/traefik

EXPOSE 80
ENTRYPOINT ["/traefik"]
