## Documentation
- https://kubernetes.io/docs/tasks/configure-pod-container/configure-persistent-volume-storage/

## Description

PersistentVolume that is backed by physical storage. Cluster user creates a PersistentVolumeClaim, which gets automatically bound to a suitable PersistentVolume. User creates a Pod that uses the PersistentVolumeClaim as storage.

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
echo-8bdc9c489-5q4bw    1/1     Running   0          10s
```

Create the file:

```
$ kubectl exec -it echo-8bdc9c489-5q4bw -- /bin/sh
echo "ok" > /data/file.txt
```

Delete the pod so the scheduler replaces it:

```
$ kubectl delete pod/echo-8bdc9c489-5q4bw
```

View the directory from the pod:

```
$ kubectl exec -it echo-xx-xx -- ls /data/
file.txt
```
