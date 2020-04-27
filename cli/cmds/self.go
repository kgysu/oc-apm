package cmds

import (
	"fmt"
	"github.com/kgysu/oc-apm/config"
	"github.com/kgysu/oc-wrapper/appitem"
	"github.com/kgysu/oc-wrapper/application"
	"github.com/kgysu/oc-wrapper/client"
	"github.com/kgysu/oc-wrapper/items"
	v1 "github.com/openshift/api/apps/v1"
	v13 "github.com/openshift/api/route/v1"
	v12 "k8s.io/api/core/v1"
	v14 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"os"
)

const ImageTag = "registry.hub.docker.com/delai21/oc-apm:0.1.2"

func RunCliCmdInstallSelf() error {
	restConf, err := client.GetRestConfig(false)
	if err != nil {
		return nil
	}

	apmApp := GetApmApp()
	apmApp.CreateItems(os.Stdout, config.GetNamespaceOrDefault(), restConf)
	fmt.Println("oc-apm successfully created.")
	return nil
}

func RunCliCmdRemoveSelf() error {
	restConf, err := client.GetRestConfig(false)
	if err != nil {
		return nil
	}

	apmApp := GetApmApp()
	apmApp.DeleteItems(os.Stdout, config.GetNamespaceOrDefault(), restConf, &metav1.DeleteOptions{})
	fmt.Println("oc-apm successfully removed.")
	return nil
}

var apmDeploymentConfig = v1.DeploymentConfig{
	TypeMeta: metav1.TypeMeta{},
	ObjectMeta: metav1.ObjectMeta{
		Name:   "oc-apm",
		Labels: map[string]string{"app": "oc-apm", "type": "operations"},
	},
	Spec: v1.DeploymentConfigSpec{
		Strategy: v1.DeploymentStrategy{},
		Replicas: 1,
		Selector: map[string]string{"app": "oc-apm"},
		Template: &v12.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Name:   "oc-apm",
				Labels: map[string]string{"app": "oc-apm", "type": "operations"},
			},
			Spec: v12.PodSpec{
				Containers: []v12.Container{
					{
						Name:  "oc-apm",
						Image: ImageTag,
						Ports: []v12.ContainerPort{
							{
								Name:          "basic",
								ContainerPort: 8080,
							},
						},
						Env: []v12.EnvVar{
							{
								Name:  "RUN_IN_CLUSTER",
								Value: "true",
							},
							{
								Name:  "NAMESPACE",
								Value: "",
								ValueFrom: &v12.EnvVarSource{
									FieldRef: &v12.ObjectFieldSelector{
										FieldPath: "metadata.namespace",
									},
								},
							},
						},
						Resources:       v12.ResourceRequirements{},
						VolumeMounts:    nil,
						VolumeDevices:   nil,
						ImagePullPolicy: v12.PullAlways,
					},
				},
				RestartPolicy:      v12.RestartPolicyAlways,
				ServiceAccountName: "oc-apm-sa",
			},
		},
	},
	Status: v1.DeploymentConfigStatus{},
}

var apmService = v12.Service{
	TypeMeta: metav1.TypeMeta{
		Kind:       "Service",
		APIVersion: "v1",
	},
	ObjectMeta: metav1.ObjectMeta{
		Name:        "oc-apm",
		Labels:      map[string]string{"app": "oc-apm", "type": "operations"},
		Annotations: map[string]string{"app": "oc-apm"},
	},
	Spec: v12.ServiceSpec{
		Ports: []v12.ServicePort{
			{
				Name:       "basic",
				Protocol:   v12.ProtocolTCP,
				Port:       8080,
				TargetPort: intstr.FromInt(8080),
			},
		},
		Selector:                 map[string]string{"app": "oc-apm"},
		Type:                     v12.ServiceTypeClusterIP,
		SessionAffinity:          v12.ServiceAffinityNone,
		PublishNotReadyAddresses: false,
	},
	Status: v12.ServiceStatus{},
}

var defaultWeight = int32(100)

var apmRoute = v13.Route{
	TypeMeta: metav1.TypeMeta{
		Kind:       "Route",
		APIVersion: "route.openshift.io/v1",
	},
	ObjectMeta: metav1.ObjectMeta{
		Name:        "oc-apm",
		Labels:      map[string]string{"app": "oc-apm", "type": "operations"},
		Annotations: map[string]string{"app": "oc-apm"},
	},
	Spec: v13.RouteSpec{
		Path: "",
		To: v13.RouteTargetReference{
			Kind:   "Service",
			Name:   "oc-apm",
			Weight: &defaultWeight,
		},
		Port: &v13.RoutePort{TargetPort: intstr.FromInt(8080)},
		TLS: &v13.TLSConfig{
			Termination:                   v13.TLSTerminationEdge,
			InsecureEdgeTerminationPolicy: v13.InsecureEdgeTerminationPolicyRedirect,
		},
		WildcardPolicy: v13.WildcardPolicyNone,
	},
	Status: v13.RouteStatus{},
}

var apmServiceAccount = v12.ServiceAccount{
	TypeMeta: metav1.TypeMeta{
		Kind:       "ServiceAccount",
		APIVersion: "v1",
	},
	ObjectMeta: metav1.ObjectMeta{
		Name:        "oc-apm-sa",
		Labels:      map[string]string{"app": "oc-apm"},
		Annotations: map[string]string{"app": "oc-apm"},
	},
}

var apmRole = v14.Role{
	TypeMeta: metav1.TypeMeta{
		Kind:       "Role",
		APIVersion: "rbac.authorization.k8s.io/v1",
	},
	ObjectMeta: metav1.ObjectMeta{
		Name:        "oc-apm-role",
		Labels:      map[string]string{"app": "oc-apm"},
		Annotations: map[string]string{"app": "oc-apm"},
	},
	Rules: []v14.PolicyRule{
		{
			Verbs:     []string{"get", "list", "create", "update", "delete"},
			APIGroups: []string{""},
			Resources: []string{"pods", "services", "serviceaccounts", "configmaps", "persistentvolumeclaims"},
		},
		{
			Verbs:     []string{"get", "list", "create", "update", "delete"},
			APIGroups: []string{"apps"},
			Resources: []string{"statefulsets"},
		},
		{
			Verbs:     []string{"get", "list", "create", "update", "delete"},
			APIGroups: []string{"apps.openshift.io"},
			Resources: []string{"deploymentconfigs"},
		},
		{
			Verbs:     []string{"get", "list", "create", "update", "delete"},
			APIGroups: []string{"route.openshift.io"},
			Resources: []string{"routes"},
		},
		{
			Verbs:     []string{"get", "list", "create", "update", "delete"},
			APIGroups: []string{"rbac.authorization.k8s.io"},
			Resources: []string{"roles", "rolebindings"},
		},
	},
}

var apmRoleBinding = v14.RoleBinding{
	TypeMeta: metav1.TypeMeta{
		Kind:       "RoleBinding",
		APIVersion: "rbac.authorization.k8s.io/v1",
	},
	ObjectMeta: metav1.ObjectMeta{
		Name:        "oc-apm-rb",
		Labels:      map[string]string{"app": "oc-apm"},
		Annotations: map[string]string{"app": "oc-apm"},
	},
	Subjects: []v14.Subject{
		{
			Kind: "ServiceAccount",
			Name: "oc-apm-sa",
		},
	},
	RoleRef: v14.RoleRef{
		Kind:     "Role",
		Name:     "oc-apm-role",
		APIGroup: "rbac.authorization.k8s.io",
	},
}

func GetApmApp() application.Application {
	return application.Application{
		Name:  "oc-apm",
		Label: "app=oc-apm",
		Items: []appitem.AppItem{
			items.NewOpServiceAccount(apmServiceAccount),
			items.NewOpRole(apmRole),
			items.NewOpRoleBinding(apmRoleBinding),
			items.NewOpDeploymentConfig(apmDeploymentConfig),
			items.NewOpService(apmService),
			items.NewOpRoute(apmRoute),
		},
	}
}
