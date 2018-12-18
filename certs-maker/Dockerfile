FROM golang:latest

RUN cd src && git clone https://github.com/jsha/minica.git && cd minica && CGO_ENABLED=0 GOOS=linux go install .

WORKDIR /certs

# ENTRYPOINT [ "minica" ]