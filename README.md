#K8s Observer

K8s Observer is a first step to start our observability.

## How to run it?

### Local mode
You can run k8s observer locally to validate new features or anything like that.
To do so, the followings steps are required.

cd to the k8s-observer folder:

1) execute go run main.go

K8s observer will be started on port 5000

ps: Running locally, we assume you have your kubeconfig configured in the ~/${HOME}/.kube/config location.
In local mode, k8s observer will use your current kube context.

### Cluster mode
In cluster mode, you must deploy the k8s observer manifest contained in the cluster-mode folder.

To do so, the following steps are required:
1) kubectl apply -f cluster-mode/observer.yaml

K8s observer will be applied to the default namespace

## How to use it?

There are two endpoints exposed:
1) GET /services
- Services with no application group will have the application group attribute omitted.
- Using this endpoint will return a list of all services deployed in the cluster, e.g:
```
GET `/services`
[
  {
    "name": "first",
    "applicationGroup": "alpha",
    "runningPodsCount": 2
  },
  {
    "name": "second",
    "runningPodsCount": 1
  },
  ...
]
```

2) GET /services/{applicationGroup}
- Using this endpoint you can filter the services by application group, e.g:
```
GET `/services/{applicationGroup}`
[
  {
    "name": "foobar",
    "applicationGroup": "<applicationGroup>",
    "runningPodsCount": 1
  },
  ...
]
```

- If you want to filter by services with no application group, you can call it as:
```
GET `/services/none`
[
  {
    "name": "foobar",
    "runningPodsCount": 1
  },
  ...
]
```