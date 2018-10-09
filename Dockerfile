FROM alpine
COPY ./urlenricher /usr/local/bin/
RUN mkdir -p /var/lib/urlenricher
VOLUME ["/var/lib/urlenricher"]
ENTRYPOINT ["/usr/local/bin/urlenricher"]
EXPOSE 8081
