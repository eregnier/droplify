FROM debian:bookworm-slim

WORKDIR /
COPY droplify .
EXPOSE 8020
ENTRYPOINT ["/droplify"]