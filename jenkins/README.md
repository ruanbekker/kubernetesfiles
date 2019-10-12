### Get a Kubernetes Cluster

Use k3d to deploy a kubernetes cluster:

```
$ k3d create --name="kubernetes" --workers="2" --publish="8080:8080"
$ export KUBECONFIG="$(k3d get-kubeconfig --name='kubernetes')"
$ kubectl cluster-info
```

### Deploy Jenkins to Kubernetes

Deploy:

```
$ kubectl apply -f app.yml
service/jenkins created
service/jenkins-service created
deployment.apps/jenkins created
```

View the deployment:

```
$ kubectl get pods,services
NAME                              READY   STATUS    RESTARTS   AGE
pod/jenkins-7cfc969f86-mjbm6      1/1     Running   0          95s
pod/svclb-jenkins-service-glmbl   1/1     Running   0          95s
pod/svclb-jenkins-service-rv2z8   1/1     Running   0          95s
pod/svclb-jenkins-service-xz522   1/1     Running   0          95s

NAME                      TYPE           CLUSTER-IP     EXTERNAL-IP                        PORT(S)              AGE
service/jenkins           ClusterIP      None           <none>                             8080/TCP,50000/TCP   95s
service/jenkins-service   LoadBalancer   10.43.20.239   172.24.0.2,172.24.0.3,172.24.0.4   8080:30554/TCP       95s
service/kubernetes        ClusterIP      10.43.0.1      <none>                             443/TCP              11m
```

### Access Jenkins

Access Jenkins on http://localhost:8080 and provide the intial admin password, which you can retrieve from the pod:

```
$ kubectl exec -it pod/jenkins-7cfc969f86-mjbm6 -- cat /var/jenkins_home/secrets/initialAdminPassword
cbdd57531af340ab8f2545235fc22942
```

And after you have login with admin you should be able to access:

<img width="1278" alt="image" src="https://user-images.githubusercontent.com/567298/66704860-4c848380-ed20-11e9-89b5-efebb718589f.png">


## Cleanup

Remove the cluster:

```
$ k3d delete --name kubernetes
```
