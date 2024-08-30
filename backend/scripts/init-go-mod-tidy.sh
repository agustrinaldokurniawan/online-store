#!/bin/bash

# Set the parent directories
PARENT_DIR="services"
API_GATEWAY_DIR="api-gateway"

# Change to the parent directory
cd "$PARENT_DIR" || exit

# Loop through each service directory and run go mod tidy
for dir in */; do
  if [ -f "${dir}go.mod" ]; then
    echo "Running go mod tidy in $dir"
    (cd "$dir" && go mod tidy)
  fi
done

# Run go mod tidy for the api-gateway directory
cd "../$API_GATEWAY_DIR" || exit
if [ -f "go.mod" ]; then
  echo "Running go mod tidy in $API_GATEWAY_DIR"
  go mod tidy
fi

echo "Done."
