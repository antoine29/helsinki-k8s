# GitOpsify Cluster with Flux

## Delete previous K3D cluster

## Create a new K3D cluster

## Add gh token to your zsh/bash config

```shell
$ export GITHUB_TOKEN=yourGHToken
```

## Install flux and bootstrap an flux repository with:

```shell
$ flux bootstrap github \
--owner=antoine29 \
--repository=kube-cluster-flux \
--personal \
--private=false
```

## At this point you can run `4.06.Project_v2.0` manifests 

