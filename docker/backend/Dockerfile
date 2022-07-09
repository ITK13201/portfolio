ARG go_version=1.18-buster

# Tools
FROM golang:${go_version} AS tools

ARG goose_version=v3.6.1
ARG air_version=1.40.3
ARG cobra_cli_version=v1.3.0

RUN wget https://github.com/pressly/goose/releases/download/${goose_version}/goose_linux_x86_64 \
    && mv goose_linux_x86_64 /usr/local/bin/goose \
    && chmod +x /usr/local/bin/goose

RUN wget https://github.com/cosmtrek/air/releases/download/v${air_version}/air_${air_version}_linux_amd64.tar.gz \
    && tar -C /usr/local/bin -xzvf air_${air_version}_linux_amd64.tar.gz \
    && chmod +x /usr/local/bin/air \
    && rm air_${air_version}_linux_amd64.tar.gz

RUN go install github.com/spf13/cobra-cli@${cobra_cli_version} \
    && mv /go/bin/cobra-cli /usr/local/bin/cobra-cli

# Dvenelopment environment
FROM golang:${go_version} AS development

ARG WORKDIR=/backend
ENV GO111MODULE=on

RUN mkdir -p ${WORKDIR}
WORKDIR ${WORKDIR}

RUN apt-get update

COPY ./backend/go.mod .
COPY ./backend/go.sum .
RUN go mod download

COPY ./backend .
COPY ./backend/bin /usr/local/bin

COPY --from=tools /usr/local/bin/goose /bin/goose
COPY --from=tools /usr/local/bin/air /bin/air
COPY --from=tools /usr/local/bin/cobra-cli /bin/cobra-cli

RUN groupadd -g 1000 golang \
    && useradd -m -s /bin/bash -u 1000 -g 1000 golang
USER golang

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]


# Builder
FROM golang:${go_version} AS builder

ARG WORKDIR=/backend
ENV GO111MODULE=on
WORKDIR ${WORKDIR}

RUN apt-get update

COPY ./backend/go.mod .
COPY ./backend/go.sum .
RUN go mod download

COPY ./backend .

RUN go build -o ./bin/portfolio


# Production environment
FROM gcr.io/distroless/base:debug AS production

COPY --from=tools /usr/local/bin/goose /bin/goose
COPY --from=builder /backend/bin/portfolio /bin/portfolio

CMD ["sh", "-c", "/bin/portfolio runserver"]
