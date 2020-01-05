FROM alpine
ADD roll_dices roll_dices

EXPOSE 8080 8888
ENTRYPOINT ["/bin/sh", "roll_dices"]