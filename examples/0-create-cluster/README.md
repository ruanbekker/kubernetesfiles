## Local Cluster with k3d

Create:

```
$ k3d create --name="demo" --workers="1" --publish="8080:8080"
$ export KUBECONFIG="$(k3d get-kubeconfig --name='demo')"
```

Remove:

```
$ k3d delete --name demo
```
