## Documentation
- https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-initialization/

## Description

Init containers runs before your application container starts. We have a shared volume on both init container and our application container. The init container will create a file and we will see the file from our pod when the application starts.

## Example

Deploy:

```
$ kubectl apply -f app.yml
persistentvolume/task-pv-volume created
persistentvolumeclaim/task-pv-claim created
deployment.apps/echo created
```

View pods:

```
$ kubectl get pods
NAME                    READY   STATUS    RESTARTS   AGE
echo-6f7787bbb7-pf8v7   1/1     Running   0          10s
```

View the directory from the pod:

```
$ kubectl exec -it echo-6f7787bbb7-pf8v7 -- ls /data/
created-file.txt
```
