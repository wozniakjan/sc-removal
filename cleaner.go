package main

import (
	"context"
	"fmt"
	"sort"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/rest"

	"log"

	gerr "github.com/pkg/errors"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1beta1"
)

const (
	HelmBrokerReleaseName           = "helm-broker"
	ServiceCatalogAddonsReleaseName = "service-catalog-addons"
	ServiceCatalogReleaseName       = "service-catalog"
	ServiceManagerProxyReleaseName  = "service-manager-proxy"
)

var (
	resources = map[string][]client.Object{
		HelmBrokerReleaseName:           hbResources,
		ServiceCatalogAddonsReleaseName: svcatAddonsResources,
		ServiceCatalogReleaseName:       svcatResources,
		ServiceManagerProxyReleaseName:  smProxyResources,
	}

	smProxyResources = []client.Object{
		&serviceManagerProxyUnstructuredServiceAccount,
		&serviceManagerProxyRegsecretUnstructuredSecret,
		&serviceManagerProxyConfigUnstructuredConfigMap,
		&serviceManagerProxyUnstructuredClusterRole,
		&serviceManagerProxyUnstructuredClusterRoleBinding,
		&serviceManagerProxyRegsecretviewerUnstructuredRole,
		&serviceManagerProxyUnstructuredRoleBinding,
		&serviceManagerProxyUnstructuredService,
		&serviceManagerProxyUnstructuredDeployment,
		&serviceManagerProxyUnstructuredServiceAccount,
		&serviceManagerProxyRegsecretUnstructuredSecret,
		&serviceManagerProxyConfigUnstructuredConfigMap,
		&serviceManagerProxyUnstructuredClusterRole,
		&serviceManagerProxyUnstructuredClusterRoleBinding,
		&serviceManagerProxyRegsecretviewerUnstructuredRole,
		&serviceManagerProxyUnstructuredRoleBinding,
		&serviceManagerProxyUnstructuredService,
		&serviceManagerProxyUnstructuredDeployment,
	}

	svcatResources = []client.Object{
		&serviceCatalogControllerManagerUnstructuredServiceAccount,
		&serviceCatalogWebhookUnstructuredServiceAccount,
		&serviceCatalogTestsUnstructuredServiceAccount,
		&serviceCatalogCatalogWebhookCertUnstructuredSecret,
		&serviceCatalogDashboardUnstructuredConfigMap,
		&servicecatalogK8SIocontrollerManagerUnstructuredClusterRole,
		&servicecatalogK8SIoserviceCatalogReadinessUnstructuredClusterRole,
		&servicecatalogK8SIowebhookUnstructuredClusterRole,
		&serviceCatalogTestsUnstructuredClusterRole,
		&servicecatalogK8SIocontrollerManagerUnstructuredClusterRoleBinding,
		&servicecatalogK8SIoserviceCatalogReadinessUnstructuredClusterRoleBinding,
		&servicecatalogK8SIowebhookUnstructuredClusterRoleBinding,
		&serviceCatalogTestsUnstructuredClusterRoleBinding,
		&servicecatalogK8SIoclusterInfoConfigmapUnstructuredRole,
		&servicecatalogK8SIoleaderLockingControllerManagerUnstructuredRole,
		&serviceCatalogControllerManagerClusterInfoUnstructuredRoleBinding,
		&serviceCatalogControllerManagerLeaderElectionUnstructuredRoleBinding,
		&serviceCatalogCatalogControllerManagerUnstructuredService,
		&serviceCatalogCatalogWebhookUnstructuredService,
		&serviceCatalogCatalogControllerManagerUnstructuredDeployment,
		&serviceCatalogCatalogWebhookUnstructuredDeployment,
		&servicecatalogUnstructuredBackendModule,
		&serviceCatalogCatalogWebhookUnstructuredMutatingWebhookConfiguration,
		&serviceCatalogCatalogControllerManagerUnstructuredPeerAuthentication,
		&serviceCatalogCatalogWebhookUnstructuredPeerAuthentication,
		&serviceCatalogCatalogControllerManagerUnstructuredServiceMonitor,
		&serviceCatalogUnstructuredTestDefinition,
		&serviceCatalogCatalogValidatingWebhookUnstructuredValidatingWebhookConfiguration,
	}

	svcatAddonsResources = []client.Object{
		&serviceCatalogAddonsServiceBindingUsageControllerCleanupUnstructuredJob,
		&serviceCatalogAddonsServiceCatalogUiUnstructuredPodSecurityPolicy,
		&serviceCatalogAddonsServiceBindingUsageControllerUnstructuredServiceAccount,
		&serviceCatalogAddonsServiceCatalogUiUnstructuredServiceAccount,
		&serviceBindingUsageControllerProcessSbuSpecUnstructuredConfigMap,
		&serviceBindingUsageControllerDashboardUnstructuredConfigMap,
		&serviceCatalogUiUnstructuredConfigMap,
		&serviceCatalogAddonsServiceBindingUsageControllerUnstructuredClusterRole,
		&serviceCatalogAddonsServiceBindingUsageControllerUnstructuredClusterRoleBinding,
		&serviceCatalogAddonsServiceCatalogUiUnstructuredRole,
		&serviceCatalogAddonsServiceCatalogUiUnstructuredRoleBinding,
		&serviceCatalogAddonsServiceBindingUsageControllerUnstructuredService,
		&serviceCatalogAddonsServiceCatalogUiUnstructuredService,
		&serviceCatalogAddonsServiceBindingUsageControllerUnstructuredDeployment,
		&serviceCatalogAddonsServiceCatalogUiUnstructuredDeployment,
		&serviceCatalogAddonsServiceCatalogUiUnstructuredDestinationRule,
		&serviceCatalogAddonsServiceBindingUsageControllerUnstructuredPeerAuthentication,
		&serviceCatalogAddonsServiceBindingUsageControllerUnstructuredServiceMonitor,
		&deploymentUnstructuredUsageKind,
		&serviceCatalogAddonsServiceCatalogUiCatalogUnstructuredVirtualService,
	}

	hbResources = []client.Object{
		&helmBrokerCleanupUnstructuredJob,
		&helmBrokerAddonsUiUnstructuredPodSecurityPolicy,
		&helmBrokerAddonsUiUnstructuredServiceAccount,
		&helmBrokerEtcdStatefulEtcdCertsUnstructuredServiceAccount,
		&helmBrokerUnstructuredServiceAccount,
		&helmSecretUnstructuredSecret,
		&helmBrokerWebhookCertUnstructuredSecret,
		&addonsUiUnstructuredConfigMap,
		&helmBrokerDashboardUnstructuredConfigMap,
		&helmConfigMapUnstructuredConfigMap,
		&sshCfgUnstructuredConfigMap,
		&helmBrokerEtcdStatefulEtcdCertsUnstructuredClusterRole,
		&helmBrokerH3UnstructuredClusterRole,
		&helmBrokerEtcdStatefulEtcdCertsUnstructuredClusterRoleBinding,
		&helmBrokerH3UnstructuredClusterRoleBinding,
		&helmBrokerAddonsUiUnstructuredRole,
		&helmBrokerAddonsUiUnstructuredRoleBinding,
		&helmBrokerAddonsUiUnstructuredService,
		&helmBrokerEtcdStatefulUnstructuredService,
		&helmBrokerEtcdStatefulClientUnstructuredService,
		&helmBrokerMetricsUnstructuredService,
		&addonControllerMetricsUnstructuredService,
		&helmBrokerUnstructuredService,
		&helmBrokerWebhookUnstructuredService,
		&helmBrokerAddonsUiUnstructuredDeployment,
		&helmBrokerUnstructuredDeployment,
		&helmBrokerWebhookUnstructuredDeployment,
		&helmBrokerEtcdStatefulUnstructuredStatefulSet,
		&helmBrokerUnstructuredAuthorizationPolicy,
		&helmReposUrlsUnstructuredClusterAddonsConfiguration,
		&addonsclustermicrofrontendUnstructuredClusterMicroFrontend,
		&addonsmicrofrontendUnstructuredClusterMicroFrontend,
		&helmBrokerAddonsUiUnstructuredDestinationRule,
		&helmBrokerEtcdStatefulClientUnstructuredDestinationRule,
		&helmBrokerMutatingWebhookUnstructuredMutatingWebhookConfiguration,
		&helmBrokerUnstructuredPeerAuthentication,
		&helmBrokerEtcdStatefulUnstructuredServiceMonitor,
		&helmBrokerUnstructuredServiceMonitor,
		&helmBrokerAddonControllerUnstructuredServiceMonitor,
		&helmBrokerAddonsUiUnstructuredVirtualService,
	}
)

type Cleaner struct {
	k8sCli            client.Client
	kubeConfigContent []byte
	stats             map[string]int
}

func NewCleaner(kubeConfigContent []byte) (*Cleaner, error) {
	var restConfig *rest.Config
	if len(kubeConfigContent) > 0 {
		var err error
		kubeconfig, err := clientcmd.NewClientConfigFromBytes(kubeConfigContent)
		if err != nil {
			return nil, err
		}
		restConfig, err = kubeconfig.ClientConfig()
		if err != nil {
			return nil, err
		}
	} else {
		var err error
		restConfig, err = rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
	}

	restConfig.Burst = 100
	restConfig.QPS = 500
	restConfig.RateLimiter = nil

	k8sCli, err := client.New(restConfig, client.Options{
		Scheme: scheme.Scheme,
	})
	if err != nil {
		return nil, err
	}
	err = v1beta1.AddToScheme(scheme.Scheme)
	if err != nil {
		return nil, err
	}
	err = apiextensions.AddToScheme(scheme.Scheme)
	if err != nil {
		return nil, err
	}

	return &Cleaner{
		k8sCli:            k8sCli,
		kubeConfigContent: kubeConfigContent,
		stats:             make(map[string]int),
	}, nil
}

func (c Cleaner) printStats() {
	msgs := make([]string, 0, len(c.stats))
	for msg, _ := range c.stats {
		msgs = append(msgs, msg)
	}
	sort.Strings(msgs)
	for _, msg := range msgs {
		if count := c.stats[msg]; count > 0 {
			log.Printf("*** %v: %v", msg, c.stats[msg])
		}
	}
}

func (c *Cleaner) RemoveRelease(releaseName string) error {
	done := make(chan bool)
	var errs []error
	go func() {
		hasResourcesToCheck := true
		for hasResourcesToCheck {
			hasResourcesToCheck = false
			errs = []error{}
			for _, r := range resources[releaseName] {
				ro := &unstructured.Unstructured{}
				ro.SetGroupVersionKind(r.GetObjectKind().GroupVersionKind())
				if err := c.k8sCli.Get(context.Background(), types.NamespacedName{Name: r.GetName(), Namespace: r.GetNamespace()}, ro); kerrors.IsNotFound(err) || meta.IsNoMatchError(err) {
					continue
				} else if err != nil {
					errs = append(errs, gerr.Wrap(err, "getting resource"))
				} else {
					if len(ro.GetFinalizers()) != 0 {
						rodc := ro.DeepCopy()
						rodc.SetGroupVersionKind(r.GetObjectKind().GroupVersionKind())
						rodc.SetFinalizers([]string{})
						if err := c.k8sCli.Update(context.Background(), rodc); err != nil {
							errs = append(errs, gerr.Wrap(err, "failed patching"))
						}
					}
				}
				if err := c.k8sCli.Delete(context.Background(), r); kerrors.IsNotFound(err) || meta.IsNoMatchError(err) {
					continue
				} else if err != nil {
					errs = append(errs, gerr.Wrap(err, fmt.Sprintf("%T %v", err, "failed deleting")))
				} else {
					log.Printf("deleting resource %v %v/%v\n", r.GetObjectKind().GroupVersionKind(), r.GetNamespace(), r.GetName())
					hasResourcesToCheck = true
				}
			}
		}
		done <- true
	}()

	select {
	case <-done:
		if len(errs) != 0 {
			return errors.NewAggregate(errs)
		}
		return nil
	case <-time.After(10 * time.Minute):
		errs = append(errs, fmt.Errorf("deleting %v timed out after 30 minutes", releaseName))
		return errors.NewAggregate(errs)
	}
}

func (c *Cleaner) RemoveResources() error {
	gvkList := []schema.GroupVersionKind{
		{
			Kind:    "ServiceBindingUsage",
			Group:   "servicecatalog.kyma-project.io",
			Version: "v1alpha1",
		},
		{
			Kind:    "UsageKind",
			Group:   "servicecatalog.kyma-project.io",
			Version: "v1alpha1",
		},
		{
			Kind:    "ServiceBinding",
			Group:   "servicecatalog.k8s.io",
			Version: "v1beta1",
		},
		{
			Kind:    "ServiceInstance",
			Group:   "servicecatalog.k8s.io",
			Version: "v1beta1",
		},
		{
			Kind:    "ServiceBroker",
			Group:   "servicecatalog.k8s.io",
			Version: "v1beta1",
		},
		{
			Kind:    "AddonsConfiguration",
			Group:   "addons.kyma-project.io",
			Version: "v1alpha1",
		},
	}

	namespaces := &v1.NamespaceList{}
	err := c.k8sCli.List(context.Background(), namespaces)
	if err != nil {
		return err
	}

	for _, gvk := range gvkList {
		msg := fmt.Sprintf("deleted %v", gvk.Kind)
		for _, namespace := range namespaces.Items {
			log.Printf("%ss in %s\n", gvk.Kind, namespace.Name)
			u := &unstructured.Unstructured{}
			u.SetGroupVersionKind(gvk)
			ul := &unstructured.UnstructuredList{}
			ul.SetGroupVersionKind(gvk)
			err := c.k8sCli.List(context.Background(), ul, client.InNamespace(namespace.Name))
			if meta.IsNoMatchError(err) {
				log.Printf("CRD for GVK %s not found, skipping resource deletion", gvk)
				break
			}
			if err != nil {
				return err
			}
			err = c.k8sCli.DeleteAllOf(context.Background(), u, client.InNamespace(namespace.Name))
			if err != nil {
				return err
			}
			c.stats[msg] += len(ul.Items)
		}
	}

	clusterGVKList := []schema.GroupVersionKind{
		{
			Kind:    "ClusterAddonsConfiguration",
			Group:   "addons.kyma-project.io",
			Version: "v1alpha1",
		},
		{
			Kind:    "ClusterServiceBroker",
			Group:   "servicecatalog.k8s.io",
			Version: "v1beta1",
		},
	}
	for _, gvk := range clusterGVKList {
		u := &unstructured.Unstructured{}
		u.SetGroupVersionKind(gvk)
		ul := &unstructured.UnstructuredList{}
		ul.SetGroupVersionKind(gvk)
		err := c.k8sCli.List(context.Background(), ul)
		if meta.IsNoMatchError(err) {
			log.Printf("CRD for GVK %s not found, skipping resource deletion", gvk)
			continue
		}
		if err != nil {
			return err
		}
		err = c.k8sCli.DeleteAllOf(context.Background(), u, client.InNamespace(""))
		if err != nil {
			return err
		}
		msg := fmt.Sprintf("deleted %v", gvk.Kind)
		c.stats[msg] += len(ul.Items)
	}
	return nil
}

func (c *Cleaner) removeFinalizers(gvk schema.GroupVersionKind, ns string) error {
	ul := &unstructured.UnstructuredList{}
	ul.SetGroupVersionKind(gvk)
	err := c.k8sCli.List(context.Background(), ul, client.InNamespace(ns))
	if err != nil {
		return gerr.Wrap(err, fmt.Sprintf("listing resources %v", gvk))
	}

	for _, obj := range ul.Items {
		obj.SetFinalizers([]string{})
		err := c.k8sCli.Update(context.Background(), &obj)
		if err != nil {
			return gerr.Wrap(err, fmt.Sprintf("updating resource %v %v/%v", gvk, obj.GetNamespace(), obj.GetName()))
		}
		log.Printf("%s %s/%s: finalizers removed", gvk.Kind, ns, obj.GetName())
		c.stats[fmt.Sprintf("removed finalizers for %v", gvk.Kind)] += 1
	}

	return nil
}

func (c *Cleaner) PrepareSBUForRemoval() error {
	namespaces := &v1.NamespaceList{}
	err := c.k8sCli.List(context.Background(), namespaces)
	if err != nil {
		return err
	}

	for _, ns := range namespaces.Items {
		ul := &unstructured.UnstructuredList{}
		ul.SetGroupVersionKind(schema.GroupVersionKind{
			Kind:    "ServiceBindingUsage",
			Group:   "servicecatalog.kyma-project.io",
			Version: "v1alpha1",
		})
		err := c.k8sCli.List(context.Background(), ul, client.InNamespace(ns.Name))
		if meta.IsNoMatchError(err) {
			log.Printf("CRD for ServiceBindingUsage not found, skipping SBU removal")
			continue
		}
		if err != nil {
			return err
		}

		for _, sbu := range ul.Items {
			log.Printf("Removing owner reference from SBU %s/%s", sbu.GetNamespace(), sbu.GetName())
			sbu.SetOwnerReferences([]metav1.OwnerReference{})
			err := c.k8sCli.Update(context.Background(), &sbu)
			if err != nil {
				return err
			}
			c.stats["removed owner references for ServiceBindingUsage"] += 1
		}
	}
	return nil
}

func (c *Cleaner) PrepareForRemoval() error {
	// listing

	namespaces := &v1.NamespaceList{}
	err := c.k8sCli.List(context.Background(), namespaces)
	if err != nil {
		return gerr.Wrap(err, "listing namespaces")
	}

	gvkList := []schema.GroupVersionKind{
		{
			Group:   "servicecatalog.k8s.io",
			Kind:    "ServiceBindingList",
			Version: "v1beta1",
		},
		{
			Group:   "servicecatalog.k8s.io",
			Kind:    "ServiceInstanceList",
			Version: "v1beta1",
		},
		{
			Group:   "servicecatalog.k8s.io",
			Kind:    "ServiceBrokerList",
			Version: "v1beta1",
		},
		{
			Group:   "servicecatalog.k8s.io",
			Kind:    "ClusterServiceBrokerList",
			Version: "v1beta1",
		},
		{
			Kind:    "UsageKind",
			Group:   "servicecatalog.kyma-project.io",
			Version: "v1alpha1",
		},
	}

	for _, gvk := range gvkList {
		for _, ns := range namespaces.Items {
			err := c.removeFinalizers(gvk, ns.Name)
			if meta.IsNoMatchError(err) {
				log.Printf("CRD for GVK %s not found, skipping finalizer removal", gvk)
				break
			}
			if err != nil {
				return gerr.Wrap(err, "removing finalizers")
			}
		}
	}

	log.Println("ServiceBindings secrets owner references")
	var bindings = &v1beta1.ServiceBindingList{}
	err = c.k8sCli.List(context.Background(), bindings, client.InNamespace(""))
	if meta.IsNoMatchError(err) {
		log.Printf("CRD for ServiceBinding not found, skipping owner reference secret adjustments")
		return nil
	}
	if err != nil {
		return gerr.Wrap(err, "listing bindings")
	}
	for i, _ := range bindings.Items {
		item := bindings.Items[i]
		log.Printf("%s/%s", item.Namespace, item.Name)
		item.Finalizers = []string{}
		err := c.k8sCli.Update(context.Background(), &item)
		if err != nil {
			return gerr.Wrap(err, fmt.Sprintf("updating binding %v/%v", item.Namespace, item.Name))
		}
		c.stats["removed finalizer for ServiceBinding"] += 1

		// find linked secrets
		var secret = &v1.Secret{}
		err = c.k8sCli.Get(context.Background(), client.ObjectKey{
			Namespace: item.Namespace,
			Name:      item.Spec.SecretName,
		}, secret)
		if secret.Name == "" {
			continue
		}
		if err != nil {
			return gerr.Wrap(err, fmt.Sprintf("getting secret %v/%v", secret.Namespace, secret.Name))
		}

		secret.OwnerReferences = []metav1.OwnerReference{}
		err = c.k8sCli.Update(context.Background(), secret)
		if err != nil {
			return gerr.Wrap(err, fmt.Sprintf("updating secret %v/%v", secret.Namespace, secret.Name))
		}
		c.stats["removed owner references for Secrets"] += 1
	}

	return nil
}

func (c *Cleaner) RemoveCRDs() error {
	crdsList := &apiextensions.CustomResourceDefinitionList{}

	err := c.k8sCli.List(context.Background(), crdsList)
	if err != nil {
		return err
	}

	for _, crd := range crdsList.Items {
		if crd.Spec.Group == "servicecatalog.k8s.io" || crd.Spec.Group == "servicecatalog.kyma-project.io" {
			log.Printf("Removing CRD %s", crd.Name)
			err := c.k8sCli.Delete(context.Background(), &crd)
			if err != nil {
				return err
			}
			c.stats["removed CRDs"] += 1
		}
	}

	return nil
}

func (c *Cleaner) waitForPodsGone(dep unstructured.Unstructured) error {
	path := []string{"spec", "selector", "matchLabels"}
	ls, found, err := unstructured.NestedStringMap(dep.Object, path...)
	if err != nil {
		msg := fmt.Sprintf("unstructured dep %v/%v failed to find selector %v: %v", dep.GetNamespace(), dep.GetName(), path, err)
		return gerr.Wrap(err, msg)
	}
	if !found {
		return fmt.Errorf("unstructured dep %v/%v missing selector %v", dep.GetNamespace(), dep.GetName(), path)
	}
	cnd := func() (bool, error) {
		pods := &v1.PodList{}
		opts := []client.ListOption{client.InNamespace(dep.GetNamespace()), client.MatchingLabels(ls)}
		if err := c.k8sCli.List(context.Background(), pods, opts...); err != nil {
			return false, err
		}
		return len(pods.Items) == 0, nil
	}
	return wait.PollImmediate(5*time.Second, 10*time.Minute, cnd)
}

func (c *Cleaner) ensureServiceCatalogNotRunning() error {
	done := make(chan error, 2)
	go func() {
		done <- c.waitForPodsGone(serviceCatalogCatalogControllerManagerUnstructuredDeployment)
	}()
	go func() {
		done <- c.waitForPodsGone(helmBrokerUnstructuredDeployment)
	}()
	var errs []error
	errs = append(errs, <-done)
	errs = append(errs, <-done)
	return errors.NewAggregate(errs)
}
