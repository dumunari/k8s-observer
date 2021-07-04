# K8s Observer

K8s Observer is an application created for the sake of studying observability, Golang and chaos engineering.
It's an ongoing process, and the main goal here is to learn :)

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

### API mode

There are two endpoints exposed:
1) GET /deployments
- Using this endpoint will return a list of all deployments and its replicas count by status, e.g:
```
GET `/deployments`
[
  {
    "name": "first",
    "runningReplicas": 2,
    "desiredReplicas": 1
  },
  {
    "name": "second",
    "runningReplicas": 2,
    "unavailableReplicas": 6,
    "desiredReplicas": 8,
  },
  ...
]
```

ps: the unavailableReplicas attribute will not be presented if there are no unavailabeReplicas.

2) GET /nodes
- Using this endpoint will return a list of all nodes and information about its conditions, e.g:
```
GET `/nodes`
[
  {
    "name": "docker-desktop",
    "memoryPressure": "False",
    "diskPressure": "False",
    "pidPressure": "False",
    "ready": "True"
  }
  ...
]
```

### Alerts mode

There are two configured alerts:
1) DeploymentsCheck
Every 5s, k8s-observer will check if there are any deployments with unavailable replicas and if it finds it, it will print
the information on console.

2) NodesCheck
Every 5s, k8s-observer will check if there are any nodes in bad conditions and if it finds it, it will print its
information on console
