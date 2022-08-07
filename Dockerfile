FROM scratch
COPY hmac-signer /hmac-signer
ENTRYPOINT ["/hmac-signer"]
