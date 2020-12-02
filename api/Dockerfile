ARG APP_NAME
ARG GOLANG_VERSION=1.15
ARG GOLANG_IMAGE=golang:${GOLANG_VERSION}

FROM ${GOLANG_IMAGE} as base
ARG APP_NAME
ENV APP_NAME=${APP_NAME}
WORKDIR /opt/${APP_NAME}
COPY go.mod go.sum ./
RUN go mod download -x
COPY . .
RUN go build -o /usr/local/bin/${APP_NAME}

FROM ${GOLANG_IMAGE}
ARG APP_NAME
ENV APP_NAME=${APP_NAME}
COPY --from=base /usr/local/bin/${APP_NAME} /usr/local/bin
CMD ${APP_NAME}
