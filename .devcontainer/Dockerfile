FROM mcr.microsoft.com/devcontainers/go:0-1-bullseye as Corkscrew
RUN apt update && apt install -y build-essential git  autoconf 
# run git clone corkscrew repository
WORKDIR /build
RUN git clone https://github.com/bryanpkc/corkscrew.git \
        && cd corkscrew \
        && autoreconf --install \
        && ./configure && make && make install

FROM mcr.microsoft.com/devcontainers/go:0-1-bullseye
RUN apt update && \
    apt install bash-completion -y &&\
    apt clean && \
    rm -rf /var/lib/apt/lists/*

COPY --from=Corkscrew /usr/local/bin/corkscrew /usr/local/bin/corkscrew