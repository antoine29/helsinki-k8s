# to setup k8s api access from out-cluster
kubectl proxy --port=8080 &
curl http://localhost:8080/api/


resourcedefinition.yaml crd
countdown.yaml

serviceaccount.yaml
deployment.yaml controller
clusterrole.yaml
clusterrolebinding.yaml


jakousa/dwk-app10-controller:sha-4256579


https://10.43.0.1:443/apis/stable.dwk/v1/countdowns?watch=true



- controller (go server) to receive a post request to create a dummy-site deployment 

-----------------------------------

kubectl proxy --port=8080 &
curl http://localhost:8080/apis/stable.anth/v1/dummysites\?watch\=true

