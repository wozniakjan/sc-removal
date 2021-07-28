package main

import (
	"context"
	helmclient "github.com/mittwald/go-helm-client"
	"helm.sh/helm/v3/pkg/storage/driver"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"time"

	"github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1beta1"
)

const serviceCatalogReleaseName = "service-catalog"

type Cleaner struct {
	k8sCli            client.Client
	kubeConfigContent []byte
}

func NewCleaner(kubeConfigContent []byte) (*Cleaner, error) {
	kubeconfig, err := clientcmd.NewClientConfigFromBytes(kubeConfigContent)
	if err != nil {
		return nil, err
	}

	rc, err := kubeconfig.ClientConfig()
	if err != nil {
		return nil, err
	}
	k8sCli, err := client.New(rc, client.Options{
		Scheme: scheme.Scheme,
	})
	if err != nil {
		return nil, err
	}
	err = v1beta1.AddToScheme(scheme.Scheme)
	if err != nil {
		return nil, err
	}

	return &Cleaner{
		k8sCli:            k8sCli,
		kubeConfigContent: kubeConfigContent,
	}, nil
}

func (c *Cleaner) RemoveServiceCatalogRelease() error {
	helmCli, err := helmclient.NewClientFromKubeConf(&helmclient.KubeConfClientOptions{
		Options: &helmclient.Options{
			Namespace: "kyma-system",
		},
		KubeConfig: c.kubeConfigContent,
	})

	log.Println("Looking for Service Catalog release...")
	release, err := helmCli.GetRelease(serviceCatalogReleaseName)
	if err == driver.ErrReleaseNotFound {
		log.Println("service-catalog release not found, nothing to do")
		return nil
	}
	if err != nil {
		return err
	}

	log.Printf("Found %s release in the namespace %s: status %s", release.Name, release.Namespace, release.Info.Status.String())
	log.Println(" Uninstalling...")
	err = helmCli.UninstallRelease(&helmclient.ChartSpec{
		ReleaseName: serviceCatalogReleaseName,
		Timeout:     time.Minute,
		Wait:        true,
	})
	if err != nil {
		return err
	}
	log.Println("DONE")
	return nil
}

func (c *Cleaner) RemoveResources() error {
	err := c.k8sCli.DeleteAllOf(context.Background(), &v1beta1.ServiceInstance{}, client.InNamespace("default"))
	if err != nil {
		return err
	}

	err = c.k8sCli.DeleteAllOf(context.Background(), &v1beta1.ServiceBinding{}, client.InNamespace("default"))
	if err != nil {
		return err
	}

	return nil
}

func (c *Cleaner) PrepareForRemoval() error {
	// listing
	log.Println("ClusterServiceBrokers")
	var clusterServiceBrokers = &v1beta1.ClusterServiceBrokerList{}
	err := c.k8sCli.List(context.Background(), clusterServiceBrokers)
	if err != nil {
		return err
	}
	for _, item := range clusterServiceBrokers.Items {
		log.Printf("%s/%s", item.Namespace, item.Name)
		item.Finalizers = []string{}
		err := c.k8sCli.Update(context.Background(), &item)
		if err != nil {
			return err
		}
	}

	log.Println("ServiceBrokers")
	var serviceBrokers = &v1beta1.ServiceBrokerList{}
	err = c.k8sCli.List(context.Background(), serviceBrokers, client.InNamespace(""))
	if err != nil {
		return err
	}
	for _, item := range serviceBrokers.Items {
		log.Printf("%s/%s", item.Namespace, item.Name)
		item.Finalizers = []string{}
		err := c.k8sCli.Update(context.Background(), &item)
		if err != nil {
			return err
		}
	}

	log.Println("ServiceInstances")
	var instances = &v1beta1.ServiceInstanceList{}
	err = c.k8sCli.List(context.Background(), instances)
	if err != nil {
		return err
	}
	for _, item := range instances.Items {
		log.Printf("%s/%s", item.Namespace, item.Name)
		item.Finalizers = []string{}
		err := c.k8sCli.Update(context.Background(), &item)
		if err != nil {
			return err
		}
	}

	log.Println("ServiceBindings")
	var bindings = &v1beta1.ServiceBindingList{}
	err = c.k8sCli.List(context.Background(), bindings, client.InNamespace(""))
	if err != nil {
		return err
	}
	for _, item := range bindings.Items {
		log.Printf("%s/%s", item.Namespace, item.Name)
		item.Finalizers = []string{}
		err := c.k8sCli.Update(context.Background(), &item)
		if err != nil {
			return err
		}

		// find linked secrets
		var secret = &v1.Secret{}
		err = c.k8sCli.Get(context.Background(), client.ObjectKey{
			Namespace: item.Namespace,
			Name:      item.Spec.SecretName,
		}, secret)
		if err != nil {
			return err
		}

		secret.OwnerReferences = []v12.OwnerReference{}
		err = c.k8sCli.Update(context.Background(), secret)
		if err != nil {
			return err
		}
	}

	return nil
}

