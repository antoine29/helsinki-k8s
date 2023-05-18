# GitOpsify Cluster with Flux

## 1. Bootstraping the todos application

* By bootstrap the application as it is in this repo, the application will use dummy/default secret values for pg password and telegram token.
If you want to set your own values refer to the section 2 of this readme.

### 1.1 Delete previous K3D cluster and recreate it

```shell
$ k3d cluster stop mycluster
$ k3d cluster delete mycluster
$ k3d cluster create mycluster --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2 --registry-use k3d-myregistry.localhost:12345
```

### 1.2 Add your gh credentials to your zsh/bash config

```shell
$ export GITHUB_USER=yourGHUserName
$ export GITHUB_TOKEN=yourGHToken
```

### 1.3 Instal flux and bootstrap the application

* It'll also work if the repo already exists

```shell
$ flux bootstrap github \
 --owner=$GITHUB_USER \
 --repository=fleet-infra \
 --branch=main \
 --path=./clusters/my-cluster \
 --personal
```

* If you're running this command for the first time, flux will create a `fleet-infra` repo on your gh account. Clone this repo and add/create files (check the files on the `fleet-infra` folder) so it ends up looking as follows:

```shell
fleet-infra
└── clusters
    └── my-cluster
        ├── flux-system
        │   ├── gotk-components.yaml
        │   ├── gotk-sync.yaml
        │   └── kustomization.yaml
        ├── todos-kustomization.yaml
        └── todos-source.yaml
```

* notice `todos-source.yaml` is the file pointing to the manifests repo. To target to other or your own fork of this repo, you'll need to modify this file properly.

### 1.4 Set the decryption secret

At this point flux is 'conciliating' the repo manifest with your cluster. However since some resources relies on the encrypted `secrets.yml` file, the conciliating process will fail.

We need to set up the sops-age private key as secret into our cluster, so flux can decrypt the secrets. 
This repo contains the dummy `age.agekey` age private key used the encrypt `secrets.yml`. (Which never should be shared like this on the repo. Did it this way to simplify this guide ;))
We'll use it do set it as a secret on the cluster.

```shell
$ cat age.agekey | kubectl create secret generic sops-age --namespace=flux-system --from-file=age.agekey=/dev/stdin
```

### 1.5 Check logs 

You can check flux logs at any moment while the re-conciliation runs. Specially handy when running the bootstrap command the first time (to see the missing-secret error) and to verify that once the private-key secret is set, the re-conciliation process goes as expected 

```shell
$ flux logs -f --all-namespaces
```

## 2. Setting up your own secrets (pg password and telegram token)

To set a more secure pg password and a working telegram token for the app, you'll need to set these values in a new encrypted K8S secret, which you'll push/replace on your fork of this repo.

2.1. Fork `fleet-infra` and manifests repos. Modify `todos-source.yaml` on your forked `fleet-infra` to point to your fork manifests repo.

2.2. Generate an age key. Save it as `age.agekey`

```shell
$ age-keygen -o age-agekey
Public key: age1eqdraxekuf5586myy60wtetq6344s80d79tvvefk73uxljwpuu7sc3ltdr

$ cat age-agekey 
# created: 2023-05-29T23:31:51-04:00
# public key: age1eqdraxekuf5586myy60wtetq6344s80d79tvvefk73uxljwpuu7sc3ltdr
AGE-SECRET-KEY-1LC6L3PN2XE09JWGFM8JWHVUKA3H6W8E3XC4P24SDJYL68KTRQ55S0L8S77
```

2.3. Create and encrypt a `secrets.plain.yaml` (containing your plain-base64 pg password and telegram token) file using sops and the age public key from your just generated `age.agekey` secret

```shell
$ cat secrets.plain.yaml
piVersion: v1
kind: Secret
metadata:
  name: secrets
  namespace: todos-ns
data:
  PG_PASSWORD: yourBase64Password
  TG_BOT_TOKEN: yourBase64TGToken

$ sops --encrypt \
--age age17mgq9ygh23q0cr00mjn0dfn8msak0apdy0ymjv5k50qzy75zmfkqzjdam4 \
--encrypted-regex '^(data)$' \
secrets.plain.yml > secrets.yml
```

2.4. Replace your encrypted `secrets.aml` secret into your manifests repo

2.5. At this point you can bootstrap the cluster (following section 1. on this guide) using your forked `fleet-infra` repo and by pointing to your forked manifests repo.

### Additional notes 

- to uninstall flux from cluster:

```shell
$ flux uninstall --namespace=flux-system 
```

Then delete the local repo folder

- manual decrypting

```shell
$ export SOPS_AGE_KEY_FILE=~/age.agekey 
$ sops --decrypt secret.yml > secret.plain.yml
```

- K8S secret values are base64 encoded strings

The `data` values on a K8S secret should be base64 encoded strings. To encode this kind of string you can use:

```shell
$ echo -n "postgres" | base64
cG9zdGdyZXM=
```

