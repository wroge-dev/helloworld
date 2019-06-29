# start with a scratch (no layers)
FROM scratch

# copy our static linked library
COPY helloworld helloworld

# tell we are exposing our service on port 8080
EXPOSE 8080

# run it!
CMD ["./helloworld"]