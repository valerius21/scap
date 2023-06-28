# Use the official Go image as the base image
FROM golang:1.20.5-bookworm

# Set the working directory inside the container
WORKDIR /app


# Install build essentials
RUN apt-get update && \
    apt-get install -y wget build-essential pkg-config --no-install-recommends

# Install ImageMagick dependencies
RUN apt-get -q -y install libjpeg-dev libpng-dev libtiff-dev \
    libgif-dev libx11-dev --no-install-recommends

# Install ImageMagick
RUN cd /root && \
    wget https://download.imagemagick.org/ImageMagick/download/ImageMagick-7.1.1-12.tar.gz && \
    tar xvzf *.tar.gz && \
    cd ImageMagick* && \
    ./configure \
        --without-magick-plus-plus \
        --without-perl \
        --disable-openmp \
        --with-gvc=no \
        --disable-docs && \
    make -j$(nproc) && make install && \
    ldconfig /usr/local/lib

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the entire project directory into the container
COPY . .

# Build the Go application
RUN go build -o app ./cmd/main.go

# Set the command to run when the container starts
CMD ["./app"]
