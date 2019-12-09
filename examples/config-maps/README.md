
## Config Maps

Create a kubernetes cluster:

```
$ k3d create --name="demo" --workers="2"
```

Export kubeconfig:

```
$ export KUBECONFIG="$(k3d get-kubeconfig --name='demo')"
```

Deploy:

```
$ kubectl apply -f app.yaml
configmap/my-config created
pod/demo created
```

Check:

```
$ kubectl get pods
NAME   READY   STATUS    RESTARTS   AGE
demo   1/1     Running   0          8s
```

View logs:

```
$ kubectl logs -f demo
test
test
```

Delete kubernetes cluster:

```
$ k3d delete --name demo
```

## More on Config Maps:
- [Create a Config Map from Env Vars in a File](https://gist.github.com/ruanbekker/a872666cade6aec79bffebea163ea981)
