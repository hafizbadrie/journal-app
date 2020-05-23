# journal-app

[![Build Status](https://travis-ci.org/hafizbadrie/journal-app.svg?branch=master)](https://travis-ci.org/hafizbadrie/journal-app)
[![Maintainability](https://api.codeclimate.com/v1/badges/90d638ee8fec8920b623/maintainability)](https://codeclimate.com/github/hafizbadrie/journal-app/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/90d638ee8fec8920b623/test_coverage)](https://codeclimate.com/github/hafizbadrie/journal-app/test_coverage)

My training ground with Golang and Kubernetes.

## Initial setup
1. Create the code base under `$GOPATH`.
2. Need to ensure I have `dep` tool.
3. Init go mod, by running `go mod init`.
4. Then, run `go mod vendor` to enable vendoring in golang.
5. Run `dep init` and `dep ensure`, but I need to remove my ssh passphrase first.
6. Build the code `go build main.go`.
7. Run it `./main`.

## For database
1. Database creation should be handled by ansible script. Database name will exist in ansible variable. That variable is used to create database and an environment variable for the app to connect to DB.
2. I use github.com/golang-migrate/migrate. Install it in mac with `brew install golang-migrate`.
3. Run `create -ext sql -dir db/migrations -seq create_journals_table` to create table.
4. Run `migrate -database postgres://hafizbadrielubis@localhost:5432/journal_app_development?sslmode=disable -path db/migrations up` to apply the schema.

## To Run
```
$> go build
$> ./journal-app
```

## To Run With Docker
```
$> docker build -t hafizbadrie/journalapp:v0.1 .
$> docker run --publish 8080:8080 --detach --name journal-app hafizbadrie/journalapp:v0.1
```

## Kubernetes Notes

### ReplicaSets

It is basically a way for us to create a set of pods easily instead of creating it one by one.

### StatefulSets

It is a workload API in kubernetes to allow us to create stateful application.

Quoting from kubernetes doc
> StatefulSets are valuable for applications that require one or more of the > following.
> 1. Stable, unique network identifiers.
> 2. Stable, persistent storage.
> 3. Ordered, graceful deployment and scaling.
> 4. Ordered, automated rolling updates.

### ReplicationController

It is similar to **ReplicaSets** but it has more rigid way to select which pods that shouold be part of it. Unlike, **ReplicaSets**
which can use `selector` to define the resources.

### CronJob

It is workload that allows us to create cron job in kubernetes.

### Node Pools

It is a way to create group of nodes so that we can have specific use cases in a kubernetes cluster. Therefore, if we want to scale out
a particular need, we don't need to scale out the whole cluster. This way we can also decide to which node pools a pod/service 
should be deployed to.

To be able to do this, we need to add labels into nodes of a particular node pool. Later on, in the resource config file we need to add this:
```
spec:
  containers:
  - name: journalapp
    image: hafizbadrie/journalapp:v0.1
    ports:
    - containerPort: 8080
  nodeSelector:
    <node-label-key>: <node-label-value>

```

## Kubernetes Cheat Sheet
```
$> kubectl config set-context <context-name> # to change kubernetes context
$> kubectl config current-context # see the current active context
$> kubectl create namespace <namespace-name> # to create a namespace
$> kubectl -n <namespace-name> apply -f <resource-file> # whether to apply config for Pod, ReplicaSet, Deployment, Service, etc
$> kubectl -n <namespace-name> scale --replicas=0 deployment <deployment-name> # scale down deployment to 0
$> kubectl -n <namespace-name> delete deployment <deployment-name> # delete a deployment

```

## References

1. How to Add DO Kubernetes Context: https://www.digitalocean.com/docs/kubernetes/how-to/connect-to-cluster/
2. Kubernetes cheat sheet: https://kubernetes.io/docs/reference/kubectl/cheatsheet/#kubectl-context-and-configuration
3. Your first service in kubernetes: https://www.digitalocean.com/community/meetup_kits/getting-started-with-containers-and-kubernetes-a-digitalocean-workshop-kit
4. Use private docker image: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
5. Play with `nodeSelector`: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector
