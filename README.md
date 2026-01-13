# For smtp server run this command in terminal and mailpit will be started on docker

docker run -d \
  --name mailpit \
  -p 1025:1025 \
  -p 8025:8025 \
  axllent/mailpit

# mailpit UI
http://localhost:8025/

# run code 
go run .