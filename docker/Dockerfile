FROM golang:1.15

WORKDIR /code

COPY entrypoint.sh /
RUN chmod u+x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

CMD ["tail", "-f", "/var/log/lastlog"]
