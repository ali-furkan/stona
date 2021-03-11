# Deployment Image
FROM ubuntu:latest

COPY docker-entrypoint.sh stona /stona/

WORKDIR /stona

# Launching ENVs
ENV PORT=7050
ENV USERNAME=stona
ENV PASSWORD=hello_stona123

EXPOSE 7050 ${PORT}
ENTRYPOINT ["docker-entrypoint.sh"]
