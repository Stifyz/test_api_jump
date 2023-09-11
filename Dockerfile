FROM golang:1.21.0-alpine3.18 as build

WORKDIR /api

ADD cmd/main.go go.mod go.sum /api/

ADD internal /api/internal

RUN go build

# Utilisez une image légère d'Alpine pour exécuter l'application
FROM alpine:3.18.3

# Installation de ca-certificates pour les requêtes HTTPS (facultatif)
RUN apk --no-cache add ca-certificates

# Copie de l'exécutable depuis la phase de construction précédente
COPY --from=build /api/test_api /api/test_api

CMD ["/api/test_api"]
