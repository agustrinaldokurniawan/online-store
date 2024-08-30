#!/bin/bash

# Navigate to the project directory
cd "$(dirname "$0")/.."

# Build and start Docker containers
docker-compose down
