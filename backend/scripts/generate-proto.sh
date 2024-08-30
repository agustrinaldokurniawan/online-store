#!/bin/bash

# Error handling function
handle_error() {
  echo "Error occurred in script at line: $1"
  exit 1
}

# Trap any error and call the handle_error function
trap 'handle_error $LINENO' ERR

# Declare an associative array to map proto directories to their corresponding pb directories
declare -A dir_map=(
  ["services/user-service/proto"]="services/user-service/pb"
  ["services/product-service/proto"]="services/product-service/pb"
  ["services/order-service/proto"]="services/order-service/pb"
)

# Define Static Dir Path
GOOGLEAPIS_DIR="googleapis"
API_GATEWAY_PB_DIR="api-gateway/pb"

# Ensure the API_GATEWAY_PB_DIR exists
if [ ! -d "$API_GATEWAY_PB_DIR" ]; then
  mkdir -p "$API_GATEWAY_PB_DIR"
  echo "Directory $API_GATEWAY_PB_DIR created."
else 
  echo "Directory $API_GATEWAY_PB_DIR already exists."
fi

# Loop through the associative array and generate Go code for each proto directory
for proto_dir in "${!dir_map[@]}"; do
  pb_dir="${dir_map[$proto_dir]}"

  # Check if the pb directory exists, if not, create it
  if [ ! -d "$pb_dir" ]; then
    mkdir -p "$pb_dir"
    echo "Directory $pb_dir created."
  else
    echo "Directory $pb_dir already exists."
  fi

  # Process each .proto file in the proto directory
  for proto_file in "$proto_dir"/*.proto; do
    # Ensure the proto file exists
    if [ -f "$proto_file" ]; then
      # Generate Go code for the current proto file
      protoc --proto_path="$proto_dir" \
        --proto_path="$GOOGLEAPIS_DIR" \
        --go_out="$pb_dir" \
        --go-grpc_out="$pb_dir" \
        --grpc-gateway_out="$pb_dir" \
        "$proto_file"
      echo "Generated Go code for $proto_file"

      # Copy the generated Go code to the API Gateway directory
      cp -auv "$pb_dir"/* "$API_GATEWAY_PB_DIR/"
      echo "Copied generated Go code to $API_GATEWAY_PB_DIR from $pb_dir"
    else
      echo "No .proto files found in $proto_dir."
    fi
  done
done
