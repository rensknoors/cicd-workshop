FROM alpine

COPY bin/HelloGo-amd64 /app/HelloGo

EXPOSE 8080

CMD /app/HelloGo
