FROM continuumio/miniconda3:latest

# Update and clears cache (to reduce image size)
RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get clean -y && \
    conda install conda-build conda-verify -y && \
    conda update --all -y && \
    conda clean --all -y

WORKDIR /var/condapkg

ENTRYPOINT ["conda"]
