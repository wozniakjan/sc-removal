module github.com/kyma-incubator/sc-removal

go 1.16

require (
	github.com/kubernetes-sigs/service-catalog v0.3.1
	github.com/kyma-project/kyma/components/function-controller v0.0.0-20220114121444-a134ba7d249f
	github.com/mittwald/go-helm-client v0.8.0
	helm.sh/helm/v3 v3.6.3
	k8s.io/api v0.21.2
	k8s.io/apiextensions-apiserver v0.21.2
	k8s.io/apimachinery v0.21.2
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	sigs.k8s.io/controller-runtime v0.9.2
)

replace k8s.io/client-go => k8s.io/client-go v0.21.2
