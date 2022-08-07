# HMAC Signer

This  is  a standalone binary used to compute an HMAC message using either SHA-1, SHA-256 or SHA-512 algorithm.

## Usage

hmac-signer is commonly invoked as
```shell
hmac-signer --key <base64-encoded-key> --text <text-to-sign> 
```

## Options

*hmac-signer* understands these options:

--algorithm string, -a string  
> hash function used for signature (sha1, sha256 or sha512) (default "sha256")  

--key bytesBase64, -k bytesBase64
> key encoded as Base64 used to sign the data  

--text string, -t string
> plain-text that will be signed  

--help, -h
> help for hmac-signer  

