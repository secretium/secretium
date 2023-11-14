#!/bin/bash

# Create a new folder for the SQLite database data.
echo "Create a new folder for the SQLite database data..."
mkdir ./secretium-data

# Create a new docker-compose.yml file with a default configuration.
echo "Create a new 'docker-compose.yml' file with a default configuration..."
cat > ./docker-compose.yml << EOL
version: '3.8'

# Define services.
services:
  # Service for the backend.
  secretium:
    # Configuration for the Docker image for the service.
    image: 'truewebartisans/secretium:latest'
    # Set restart rules for the container.
    restart: unless-stopped
    # Forward the exposed port 8787 on the container to the host machine.
    expose:
      - '8787'
    # Set required environment variables for the backend.
    environment:
      SECRET_KEY: /run/secrets/secretium_key
      MASTER_USERNAME: /run/secrets/secretium_master_username
      MASTER_PASSWORD: /run/secrets/secretium_master_password
      DOMAIN: /run/secrets/secretium_domain
    # Set volumes for the container with SQLite data and the root SSL certificates.
    volumes:
      - ./secretium-data:/secretium-data
      - /etc/ssl/certs:/etc/ssl/certs:ro

# Define the secrets.
secrets:
  # Key for the secret sharer.
  secretium_key:
    # Path to the file with your secret key.
    file: secretium_key.txt
  # Master username used for the Secretium dashboard.
  secretium_master_username:
    # Path to the file with your master username.
    file: secretium_master_username.txt
  # Master password used for the Secretium dashboard.
  secretium_master_password:
    # Path to the file with your master password.
    file: secretium_master_password.txt
  # Domain name for the Secretium's links.
  secretium_domain:
    # Path to the file with your domain name.
    file: secretium_domain.txt
EOL

# Run the 'docker-compose up -d' command to start the container.
echo "Run the 'docker-compose up -d' command to start the container..."
docker-compose up -d

# Remove the TXT files with the sensitive data.
echo "Remove the TXT files with the sensitive data..."
rm secretium_key.txt secretium_master_username.txt secretium_master_password.txt secretium_domain.txt

# Final words.
echo "All tasks done!"