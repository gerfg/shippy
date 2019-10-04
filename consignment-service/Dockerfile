# Alpine linux image - light-weight linux distribution
# developed and optmised for running dockerised web app (image-size: 8mb)
FROM debian:latest

RUN mkdir /app
WORKDIR /app
ADD consignment-service /app/consignment-service

CMD ["./consignment-service"]
