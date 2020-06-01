FROM golang:1.14-stretch as builder

WORKDIR /app

COPY . .
RUN go build -o pcr .

FROM continuumio/miniconda3:latest
RUN conda config --set always_yes true && \
    conda update --all && \
    conda install conda-build conda-verify && \
    conda clean --all

RUN mkdir -p /var/condapkg

WORKDIR /app

# this migration line is tied to /infrastructure/database/postgres/migrate.go implementation.
# it enables users to use the latest migration source without needing to download it from github
# more of a convenience
COPY infrastructure/database/migrations     infrastructure/database/migrations
COPY --from=builder /app/pcr                pcr
COPY --from=builder /app/config.yaml        /var/private-conda-repo/config.yaml

ENTRYPOINT ["./pcr"]
