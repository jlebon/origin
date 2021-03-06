// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/operator/v1"
	scheme "github.com/openshift/client-go/operator/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KubeControllerManagersGetter has a method to return a KubeControllerManagerInterface.
// A group's client should implement this interface.
type KubeControllerManagersGetter interface {
	KubeControllerManagers() KubeControllerManagerInterface
}

// KubeControllerManagerInterface has methods to work with KubeControllerManager resources.
type KubeControllerManagerInterface interface {
	Create(*v1.KubeControllerManager) (*v1.KubeControllerManager, error)
	Update(*v1.KubeControllerManager) (*v1.KubeControllerManager, error)
	UpdateStatus(*v1.KubeControllerManager) (*v1.KubeControllerManager, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.KubeControllerManager, error)
	List(opts meta_v1.ListOptions) (*v1.KubeControllerManagerList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.KubeControllerManager, err error)
	KubeControllerManagerExpansion
}

// kubeControllerManagers implements KubeControllerManagerInterface
type kubeControllerManagers struct {
	client rest.Interface
}

// newKubeControllerManagers returns a KubeControllerManagers
func newKubeControllerManagers(c *OperatorV1Client) *kubeControllerManagers {
	return &kubeControllerManagers{
		client: c.RESTClient(),
	}
}

// Get takes name of the kubeControllerManager, and returns the corresponding kubeControllerManager object, and an error if there is any.
func (c *kubeControllerManagers) Get(name string, options meta_v1.GetOptions) (result *v1.KubeControllerManager, err error) {
	result = &v1.KubeControllerManager{}
	err = c.client.Get().
		Resource("kubecontrollermanagers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KubeControllerManagers that match those selectors.
func (c *kubeControllerManagers) List(opts meta_v1.ListOptions) (result *v1.KubeControllerManagerList, err error) {
	result = &v1.KubeControllerManagerList{}
	err = c.client.Get().
		Resource("kubecontrollermanagers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kubeControllerManagers.
func (c *kubeControllerManagers) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("kubecontrollermanagers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a kubeControllerManager and creates it.  Returns the server's representation of the kubeControllerManager, and an error, if there is any.
func (c *kubeControllerManagers) Create(kubeControllerManager *v1.KubeControllerManager) (result *v1.KubeControllerManager, err error) {
	result = &v1.KubeControllerManager{}
	err = c.client.Post().
		Resource("kubecontrollermanagers").
		Body(kubeControllerManager).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kubeControllerManager and updates it. Returns the server's representation of the kubeControllerManager, and an error, if there is any.
func (c *kubeControllerManagers) Update(kubeControllerManager *v1.KubeControllerManager) (result *v1.KubeControllerManager, err error) {
	result = &v1.KubeControllerManager{}
	err = c.client.Put().
		Resource("kubecontrollermanagers").
		Name(kubeControllerManager.Name).
		Body(kubeControllerManager).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *kubeControllerManagers) UpdateStatus(kubeControllerManager *v1.KubeControllerManager) (result *v1.KubeControllerManager, err error) {
	result = &v1.KubeControllerManager{}
	err = c.client.Put().
		Resource("kubecontrollermanagers").
		Name(kubeControllerManager.Name).
		SubResource("status").
		Body(kubeControllerManager).
		Do().
		Into(result)
	return
}

// Delete takes name of the kubeControllerManager and deletes it. Returns an error if one occurs.
func (c *kubeControllerManagers) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("kubecontrollermanagers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kubeControllerManagers) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Resource("kubecontrollermanagers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kubeControllerManager.
func (c *kubeControllerManagers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.KubeControllerManager, err error) {
	result = &v1.KubeControllerManager{}
	err = c.client.Patch(pt).
		Resource("kubecontrollermanagers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
