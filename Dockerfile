FROM scratch

COPY httpecho /

EXPOSE 8080

ENTRYPOINT ["/httpecho"]
