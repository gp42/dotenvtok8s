# dotenvtok8s
Helper tool to convert .env files to Kubernetes ConfigMaps and Secrets. Environment Variables take
precedence over values set in the .env file.

The secrets are separated with the help of a variable name prefix.

## Usage
Generate a ConfigMap from all parameters in .env file:
```
./build/dotenvtok8s -in-file test/testdata/.env -out-dir /tmp
```

Generate a ConfigMap and a Secret from all parameters prefixed with SECRET_:
```
./build/dotenvtok8s -in-file test/testdata/.env -out-dir /tmp -secret-prefix SECRET_
```
