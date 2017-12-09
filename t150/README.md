# cert/key generation

```text
$ mkdir -v certs
$ openssl req -newkey rsa:4096 -nodes -sha256 -keyout certs/prvt_registry.key -x509 -days 365 -out certs/prvt_registry.cert
Generating a 4096 bit RSA private key
.................++
.......................................................................................................................................................++
writing new private key to 'certs/prvt_registry.key'
-----
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [AU]:GR
State or Province Name (full name) [Some-State]:
Locality Name (eg, city) []:
Organization Name (eg, company) [Internet Widgits Pty Ltd]:
Organizational Unit Name (eg, section) []:
Common Name (e.g. server FQDN or YOUR name) []:
Email Address []:
```
