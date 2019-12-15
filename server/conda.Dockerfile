FROM continuumio/miniconda3:latest

RUN apt-get update && \
    conda install conda-build conda-verify -y && \
    conda update --all -y

WORKDIR /var/condapkg

ENTRYPOINT ["conda"]
