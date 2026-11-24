# Hardened image for mariana. Build: docker build --build-arg APP_VERSION=$(git describe --tags --always) -t mariana .
# Base images are pinned by digest; bump them deliberately.
FROM cgr.dev/chainguard/go:latest@sha256:b42c64f7559093cf20910fc6fe750884dd55c74057a16cbeedbba42c4f9b31f8 AS builder

WORKDIR /workdir

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
ARG APP_VERSION=dev
RUN CGO_ENABLED=0 go build -trimpath -buildvcs=false -ldflags="-s -w -X main.AppVersion=${APP_VERSION}" -o /app .

# --- Runtime: no shell, no package manager, only the static binary ---
FROM cgr.dev/chainguard/static:latest@sha256:fc42af33e643f20d08baf03cddcc1c02ba57052dfb02edb0d79f3fc47bcea8d9

COPY --from=builder /app /app

# Numeric so Kubernetes runAsNonRoot can verify it.
USER 65532:65532

ENV APP_ENV=production APP_PORT=8080

EXPOSE 8080

# No shell or curl here to run a probe; probe from the orchestrator.
HEALTHCHECK NONE

# DB_DSN is required and must come from the runtime (secret), never the image.
ENTRYPOINT ["/app"]
CMD ["serve"]
