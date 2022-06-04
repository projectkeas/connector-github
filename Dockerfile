FROM gcr.io/distroless/static
ARG TARGETOS
ARG TARGETARCH
WORKDIR /app
COPY ./connector-github-${TARGETOS}-${TARGETARCH} /app/connector-github
ENTRYPOINT [ "/app/connector-github" ]