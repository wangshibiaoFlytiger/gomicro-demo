FROM alpine
ADD gomicro-service /gomicro-service
ENTRYPOINT [ "/gomicro-service" ]
