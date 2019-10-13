FROM scratch

ADD build/ ./
ADD views ./views
ENV PORT 8080
EXPOSE 8080
ENTRYPOINT ["/gowiki"]
