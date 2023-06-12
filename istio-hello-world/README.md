# hello-world

Hello world plain HTML webpage.

## Getting Started

Deploy to your cluster:

```bash
kubectl create namespace hello-world
kubectl create -f https://raw.githubusercontent.com/frarodroc/kubernetes/master/hello-world/hello-world.yaml
```

I am using Bitnami image because their image expose the port 8080 without a S2I process.


