# Health Checks

This is a simple app to test Kubernetes container health check configurations.

## Getting started

```
podman build . -t quay.io/frarodroc/health-checks:latest
podman login quay.io
podman push quay.io/frarodroc/health-checks:latest
```

Do not forget to change your repository visibility if it's the first time you push an 
image into it.

```
oc create -f health-checks.yaml
```

## Authors

* Francisco Rodríguez Rocha
