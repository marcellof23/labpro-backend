FROM golang:1.16.5-buster

# Making app folder, copy everything to it, and set working directory
RUN mkdir -p /app
COPY . /app
WORKDIR /app

# Update debian and install tools for easy debugging
RUN apt update 

# Install and tidy golang dependency for server 
RUN go mod tidy \
    && go mod vendor 

# Install air
RUN go get github.com/cosmtrek/air

# Install latest lts nodejs for server
RUN curl -sL https://deb.nodesource.com/setup_14.x | bash - \
    && apt-get install nodejs  -y

# Change script permission
RUN chmod +x ./scripts/init-prod-server.sh

# Run the server
CMD ["bash", "./scripts/init-prod-server.sh"]

EXPOSE 8080