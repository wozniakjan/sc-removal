# Builder image
FROM golang:1.16 as builder

# Workaround because migration tool is private repo
ARG SSH_PRIVATE_KEY
RUN mkdir /root/.ssh/ && \
    echo "${SSH_PRIVATE_KEY}" > /root/.ssh/id_rsa && \
    chmod 600 /root/.ssh/id_rsa && \
    ssh-keyscan -t rsa -H github.com > ~/.ssh/known_hosts

ENV GOPRIVATE=github.com/SAP/sap-btp-service-operator-migration
RUN git config --global url.git@github.com:.insteadOf https://github.com/

RUN CGO_ENABLED=0 go install github.com/SAP/sap-btp-service-operator-migration@v0.1.0

# Run image
FROM gcr.io/distroless/static:latest

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/bin/sap-btp-service-operator-migration /bin/sap-btp-service-operator-migration

ENTRYPOINT ["sap-btp-service-operator-migration"]
