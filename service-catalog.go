// Code generated by reverse-kube-resource. DO NOT EDIT.

package main

import v1unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

var (
	// Unstructured "service-catalog-controller-manager"
	serviceCatalogControllerManagerUnstructuredServiceAccount = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"name":      "service-catalog-controller-manager",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-catalog-webhook"
	serviceCatalogWebhookUnstructuredServiceAccount = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"name":      "service-catalog-webhook",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-catalog-tests"
	serviceCatalogTestsUnstructuredServiceAccount = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"name":      "service-catalog-tests",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-catalog-catalog-webhook-cert"
	serviceCatalogCatalogWebhookCertUnstructuredSecret = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Secret",
			"metadata": map[string]interface{}{
				"name":      "service-catalog-catalog-webhook-cert",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-catalog-dashboard"
	serviceCatalogDashboardUnstructuredConfigMap = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]interface{}{
				"name":      "service-catalog-dashboard",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "servicecatalog.k8s.io:controller-manager"
	servicecatalogK8SIocontrollerManagerUnstructuredClusterRole = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name":      "servicecatalog.k8s.io:controller-manager",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "servicecatalog.k8s.io:service-catalog-readiness"
	servicecatalogK8SIoserviceCatalogReadinessUnstructuredClusterRole = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name":      "servicecatalog.k8s.io:service-catalog-readiness",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "servicecatalog.k8s.io:webhook"
	servicecatalogK8SIowebhookUnstructuredClusterRole = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name":      "servicecatalog.k8s.io:webhook",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-catalog-tests"
	serviceCatalogTestsUnstructuredClusterRole = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name":      "service-catalog-tests",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "servicecatalog.k8s.io:controller-manager"
	servicecatalogK8SIocontrollerManagerUnstructuredClusterRoleBinding = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name":      "servicecatalog.k8s.io:controller-manager",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "servicecatalog.k8s.io:service-catalog-readiness"
	servicecatalogK8SIoserviceCatalogReadinessUnstructuredClusterRoleBinding = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name":      "servicecatalog.k8s.io:service-catalog-readiness",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "servicecatalog.k8s.io:webhook"
	servicecatalogK8SIowebhookUnstructuredClusterRoleBinding = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name":      "servicecatalog.k8s.io:webhook",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-catalog-tests"
	serviceCatalogTestsUnstructuredClusterRoleBinding = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name":      "service-catalog-tests",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "servicecatalog.k8s.io:cluster-info-configmap"
	servicecatalogK8SIoclusterInfoConfigmapUnstructuredRole = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "Role",
			"metadata": map[string]interface{}{
				"name":      "servicecatalog.k8s.io:cluster-info-configmap",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "servicecatalog.k8s.io:leader-locking-controller-manager"
	servicecatalogK8SIoleaderLockingControllerManagerUnstructuredRole = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "Role",
			"metadata": map[string]interface{}{
				"name":      "servicecatalog.k8s.io:leader-locking-controller-manager",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-catalog-controller-manager-cluster-info"
	serviceCatalogControllerManagerClusterInfoUnstructuredRoleBinding = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "RoleBinding",
			"metadata": map[string]interface{}{
				"name":      "service-catalog-controller-manager-cluster-info",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-catalog-controller-manager-leader-election"
	serviceCatalogControllerManagerLeaderElectionUnstructuredRoleBinding = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "RoleBinding",
			"metadata": map[string]interface{}{
				"name":      "service-catalog-controller-manager-leader-election",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-catalog-catalog-controller-manager"
	serviceCatalogCatalogControllerManagerUnstructuredService = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "service-catalog-catalog-controller-manager",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-catalog-catalog-webhook"
	serviceCatalogCatalogWebhookUnstructuredService = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name":      "service-catalog-catalog-webhook",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-catalog-catalog-controller-manager"
	serviceCatalogCatalogControllerManagerUnstructuredDeployment = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name":      "service-catalog-catalog-controller-manager",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-catalog-catalog-webhook"
	serviceCatalogCatalogWebhookUnstructuredDeployment = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name":      "service-catalog-catalog-webhook",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "servicecatalog"
	servicecatalogUnstructuredBackendModule = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "ui.kyma-project.io/v1alpha1",
			"kind":       "BackendModule",
			"metadata": map[string]interface{}{
				"name":      "servicecatalog",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-catalog-catalog-webhook"
	serviceCatalogCatalogWebhookUnstructuredMutatingWebhookConfiguration = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "admissionregistration.k8s.io/v1beta1",
			"kind":       "MutatingWebhookConfiguration",
			"metadata": map[string]interface{}{
				"name":      "service-catalog-catalog-webhook",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-catalog-catalog-controller-manager"
	serviceCatalogCatalogControllerManagerUnstructuredPeerAuthentication = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "security.istio.io/v1beta1",
			"kind":       "PeerAuthentication",
			"metadata": map[string]interface{}{
				"name":      "service-catalog-catalog-controller-manager",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-catalog-catalog-webhook"
	serviceCatalogCatalogWebhookUnstructuredPeerAuthentication = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "security.istio.io/v1beta1",
			"kind":       "PeerAuthentication",
			"metadata": map[string]interface{}{
				"name":      "service-catalog-catalog-webhook",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-catalog-catalog-controller-manager"
	serviceCatalogCatalogControllerManagerUnstructuredServiceMonitor = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "monitoring.coreos.com/v1",
			"kind":       "ServiceMonitor",
			"metadata": map[string]interface{}{
				"name":      "service-catalog-catalog-controller-manager",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-catalog"
	serviceCatalogUnstructuredTestDefinition = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "testing.kyma-project.io/v1alpha1",
			"kind":       "TestDefinition",
			"metadata": map[string]interface{}{
				"name":      "service-catalog",
				"namespace": "kyma-system",
			},
		},
	}

	// Unstructured "service-catalog-catalog-validating-webhook"
	serviceCatalogCatalogValidatingWebhookUnstructuredValidatingWebhookConfiguration = v1unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "admissionregistration.k8s.io/v1beta1",
			"kind":       "ValidatingWebhookConfiguration",
			"metadata": map[string]interface{}{
				"name":      "service-catalog-catalog-validating-webhook",
				"namespace": "kyma-system",
			},
		},
	}
)