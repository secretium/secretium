# A default Dockerfile for the GoReleaser process. DO NOT EDIT!

FROM scratch
COPY project /
ENTRYPOINT ["/project"]
