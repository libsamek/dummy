# Dummy project

## How to reproduce Helm does not wait for Statefulset to be ready

More information can be found: https://github.com/helm/helm/issues/10163

Docker and kind (Kubernetes in Docker) are needed 

Steps to setup environment:
```
docker build -t localhost:5000/dummy .
infra/kind-with-registry.sh
helm upgrade --install --wait dummy infra/dummy
```

Change something in `values.yaml`, eg. bump `dummy.version`.

Run once again upgrade and observe behavior.

```
helm upgrade --install --wait dummy infra/dummy
```
