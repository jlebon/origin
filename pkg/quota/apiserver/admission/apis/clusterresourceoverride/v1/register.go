package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/openshift/origin/pkg/quota/apiserver/admission/apis/clusterresourceoverride"
)

func (obj *ClusterResourceOverrideConfig) GetObjectKind() schema.ObjectKind { return &obj.TypeMeta }

var GroupVersion = schema.GroupVersion{Group: "autoscaling.openshift.io", Version: "v1"}

var (
	schemeBuilder = runtime.NewSchemeBuilder(
		addKnownTypes,
		clusterresourceoverride.Install,
	)
	Install = schemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(GroupVersion,
		&ClusterResourceOverrideConfig{},
	)
	return nil
}
