

kind create cluster

kubectl apply -f hack/resources/crds

helm repo add svc-cat https://kubernetes-sigs.github.io/service-catalog
kubectl create ns kyma-system

helm install service-catalog hack/charts/service-catalog --namespace kyma-system --set asyncBindingOperationsEnabled=true

kubectl apply -f hack/charts/podpresets.settings.svcat.crd.yaml
helm install pod-preset hack/charts/pod-preset --namespace kyma-system

# helm broker
helm install helm-broker hack/charts/helm-broker --namespace kyma-system

kubectl apply -f hack/resources/sample-addons.yaml

# create an instance and a binding
kubectl apply -f hack/resources/instance.yaml
kubectl apply -f hack/resources/binding.yaml

# Run the tool
go run main.go cleaner.go

# the secret must still exists
kubectl get secret testing