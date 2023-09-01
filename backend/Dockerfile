FROM golang:1.20-alpine

LABEL maintainer="tuannm <nguyentuan1696@gmail.com>"
WORKDIR /app
ENV BIN_DIR=/app/bin

RUN apk update && apk add --no-cache gcc build-base

RUN apk add --no-cache tzdata

RUN apk add --no-cache git

ADD . .

RUN sed '/^replace\s*.*=>\s.*/d' go.mod > go1.mod && mv go1.mod go.mod

RUN chmod +x scripts/build.sh && ./scripts/build.sh

FROM alpine:latest
WORKDIR /app
ENV BIN_DIR=/app/bin
ENV SCRIPTS_DIR=/app/scripts

# Build Args
ARG LOG_DIR=/var/log/thichlab

# Create log directory
RUN mkdir -p ${BIN_DIR} \
	mkdir -p ${SCRIPTS_DIR} \
	mkdir -p ${LOG_DIR} \
	&& apk update \
	&& addgroup -S thichlab \
	&& adduser -S thichlab -G thichlab \
	&& chown thichlab:thichlab /app \
	&& chown thichlab:thichlab ${LOG_DIR}

ENV TZ=Asia/Ho_Chi_Minh
RUN apk add --no-cache tzdata
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

USER thichlab

COPY --chown=thichlab:thichlab --from=0 ${BIN_DIR} ${BIN_DIR}
COPY --chown=thichlab:thichlab --from=0 ${SCRIPTS_DIR} /app
COPY --chown=thichlab:thichlab --from=0 /app/config.json /app

RUN chmod +x  startup.sh

ENTRYPOINT ["/app/startup.sh"]

