FROM alpine
ADD campaign-service /campaign-service
ENTRYPOINT [ "/campaign-service" ]
