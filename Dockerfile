FROM golang:1.16-alpine

WORKDIR /app

# Download necessary Go modules
COPY ./haproxy-test-app /app/
# RUN chmod +x /app/haproxy-test-app
EXPOSE 3000

CMD [ "/app/haproxy-test-app" ]