FROM golang:1.13-stretch as builder

WORKDIR /app

COPY . .
RUN go build -o pcr .

FROM continuumio/miniconda3:latest
RUN conda config --set always_yes true && \
    conda update --all && \
    conda install conda-build && \
    conda clean --all

RUN mkdir -p /var/condapkg

WORKDIR /app
COPY --from=builder /app/pcr            pcr
COPY --from=builder /app/config.yaml    /var/private-conda-repo/config.yaml

ENTRYPOINT ["./pcr"]
