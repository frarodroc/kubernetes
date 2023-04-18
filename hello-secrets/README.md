# Hello Secrets

This is an example of how to read environment variables from Secrets.

This project contains two folders:

1. A [Helm Chart](./helm/) project
2. The [source code](./src/) of the application

## Build the container image

Build the container image:

```
podman build ./src -t quay.io/frarodroc/hello-secrets:1.0.0
```

Login to the public or private registry:

```
podman login -u frarodroc quay.io
```

Push the container image to a public or private repository:

```
podman push quay.io/frarodroc/hello-secrets:1.0.0
```

