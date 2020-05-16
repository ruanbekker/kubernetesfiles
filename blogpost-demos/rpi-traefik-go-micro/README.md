Deploy the home service:

```
$ kubectl apply -f home/deployment.yml
deployment.apps/go-micro-home created
service/go-micro-home-service created
ingress.extensions/go-micro-home-ingress created
```

Deploy the contact service:

```
$ kubectl apply -f contact/deployment.yml
deployment.apps/go-micro-contact created
service/go-micro-contact-service created
ingress.extensions/go-micro-contact-ingress created
```

Deploy the about service:

```
$ kubectl apply -f about/deployment.yml
deployment.apps/go-micro-about created
service/go-micro-about-service created
ingress.extensions/go-micro-about-ingress created
```

View the deployments:

```
$ kubectl get deployments -o wide
NAME               READY   UP-TO-DATE   AVAILABLE   AGE   CONTAINERS         IMAGES                               SELECTOR
go-micro-home      1/1     1            1           74s   go-micro-home      ruanbekker/rpi-go-micro:home         app=go-micro-home
go-micro-contact   1/1     1            1           54s   go-micro-contact   ruanbekker/rpi-go-micro:contact-us   app=go-micro-contact
go-micro-about     1/1     1            1           39s   go-micro-about     ruanbekker/rpi-go-micro:about-us     app=go-micro-about
```

Test it out:

```
$ curl http://rpi-05.ruan.local/
content: home, served from: go-micro-home-65d5f5f78b-tl8kf

$ http://rpi-05.ruan.local/about
content: about-us, served from: go-micro-about-7ddcbc94bc-h5wt2

$ http://rpi-05.ruan.local/contact
content: contact-us, served from: go-micro-contact-5f4c74d864-p528h
```

Scale each deployment to 5 containers:

```
$ kubectl scale deployment/go-micro-home --replicas 5
deployment.apps/go-micro-home scaled

$ hostname kubectl scale deployment/go-micro-contact --replicas 5
deployment.apps/go-micro-contact scaled

$ hostname kubectl scale deployment/go-micro-about --replicas 5
deployment.apps/go-micro-about scaled
```

Verify that the scaling event completed:

```
$ kubectl get deployments -o wide
NAME               READY   UP-TO-DATE   AVAILABLE   AGE   CONTAINERS         IMAGES                               SELECTOR
go-micro-home      5/5     5            5           15m   go-micro-home      ruanbekker/rpi-go-micro:home         app=go-micro-home
go-micro-contact   5/5     5            5           15m   go-micro-contact   ruanbekker/rpi-go-micro:contact-us   app=go-micro-contact
go-micro-about     5/5     5            5           15m   go-micro-about     ruanbekker/rpi-go-micro:about-us     app=go-micro-about
```

View the pods per app selector:

```
$ get pods -o wide -l app=go-micro-home
NAME                             READY   STATUS    RESTARTS   AGE     IP           NODE     NOMINATED NODE   READINESS GATES
go-micro-home-65d5f5f78b-tl8kf   1/1     Running   0          16m     10.42.1.13   rpi-05   <none>           <none>
go-micro-home-65d5f5f78b-gxz2h   1/1     Running   0          3m56s   10.42.0.71   rpi-06   <none>           <none>
go-micro-home-65d5f5f78b-snq6p   1/1     Running   0          3m56s   10.42.0.72   rpi-06   <none>           <none>
go-micro-home-65d5f5f78b-qkfql   1/1     Running   0          3m56s   10.42.1.16   rpi-05   <none>           <none>
go-micro-home-65d5f5f78b-ww6vv   1/1     Running   0          3m56s   10.42.1.17   rpi-05   <none>           <none>
```

Test it out:

```
$ curl http://rpi-05.ruan.local/contact
content: home, served from: go-micro-home-65d5f5f78b-snq6p

$ hostname curl http://rpi-05.ruan.local/
content: home, served from: go-micro-home-65d5f5f78b-tl8kf

$ hostname curl http://rpi-05.ruan.local/
content: home, served from: go-micro-home-65d5f5f78b-ww6vv
```
