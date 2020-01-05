FROM alpine
ADD target/game game
ADD target/run.sh run.sh

EXPOSE 8080 8888
ENTRYPOINT ["/bin/sh", "game"]
#ENTRYPOINT ["/bin/sh", "./run.sh"]