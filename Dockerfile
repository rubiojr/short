FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
ENV PATH=/bin
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY short /usr/local/bin/short

# Create /data directory
WORKDIR /data
# Expose data volume
VOLUME /data

EXPOSE 8080/tcp


# Set the default command
ENTRYPOINT [ "/usr/local/bin/short" ]
#CMD [ "serve" ]

