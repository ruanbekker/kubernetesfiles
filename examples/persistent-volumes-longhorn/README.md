# longhorn-persistent-volumes

## Pre-Requisites

See https://longhorn.io/docs/1.3.1/deploy/install/

```bash
sudo apt install open-iscsi -y # on all nodes
helm repo add longhorn https://charts.longhorn.io
helm repo update
helm install longhorn longhorn/longhorn --namespace longhorn-system --create-namespace
kubectl -n longhorn-system get pod
```
