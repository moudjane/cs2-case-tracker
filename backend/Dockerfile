# Étape 1 : Build de l'application
FROM golang:1.23 AS builder

WORKDIR /app

# Copier uniquement les fichiers de dépendances pour optimiser le cache Docker
COPY go.mod go.sum ./
RUN go mod tidy

# Copier le reste du code source
COPY . .

# Compiler l'application Go en un binaire statique compatible avec Alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/cs2-case-tracker-go ./main.go

# Étape 2 : Conteneur d'exécution minimal
FROM alpine:latest

# Installer les dépendances nécessaires pour l'exécution
RUN apk --no-cache add ca-certificates

# Définir le répertoire de travail
WORKDIR /app

# Copier l'exécutable depuis l'étape de build
COPY --from=builder /app/cs2-case-tracker-go .

# Rendre l'exécutable utilisable
RUN chmod +x /app/cs2-case-tracker-go

# Exposer le port utilisé par l'application
EXPOSE 5001

# Définir la commande de lancement de l'application
CMD ["/app/cs2-case-tracker-go"]
