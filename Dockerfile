FROM scratch

COPY cmd/cmd .

EXPOSE 8104

ENTRYPOINT [ "./cmd" ]