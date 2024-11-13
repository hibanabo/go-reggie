FROM ubuntu:latest
LABEL authors="hiban"

ENTRYPOINT ["top", "-b"]