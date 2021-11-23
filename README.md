# Dummy project

## How to reproduce Helm does not wait for Statefulset to be ready

More information can be found: https://github.com/helm/helm/issues/10163

`dummy.go` is a simple Go HTTP service, which takes environment variable `DUMMY_TIMEOUT`
to define how long it will take to terminate service on SIGINT or SIGTERM.

### Reproduce

Docker and kind (Kubernetes in Docker) are needed.

Steps to setup environment:
```
docker build -t localhost:5000/dummy .
docker push localhost:5000/dummy:latest

infra/kind-with-registry.sh
helm upgrade --debug --install --wait dummy infra/dummy
```

Change something in `values.yaml`, eg. bump `dummy.version`.

Run once again upgrade and observe behavior.

```
helm upgrade --debug --install --wait dummy infra/dummy
```
