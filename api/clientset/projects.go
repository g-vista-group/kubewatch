package clientset

import (
	"github.com/g-vista-group/kubewatch/api/types/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type TenantInterface interface {
	List(opts metav1.ListOptions) (*v1alpha1.TenantList, error)
	Get(name string, options metav1.GetOptions) (*v1alpha1.Tenant, error)
	Create(*v1alpha1.Tenant) (*v1alpha1.Tenant, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	// ...
}

type tenantClient struct {
	restClient rest.Interface
	ns         string
}

func (c *tenantClient) List(opts metav1.ListOptions) (*v1alpha1.TenantList, error) {
	result := v1alpha1.TenantList{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("tenants").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

func (c *tenantClient) Get(name string, opts metav1.GetOptions) (*v1alpha1.Tenant, error) {
	result := v1alpha1.Tenant{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("tenants").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

func (c *tenantClient) Create(project *v1alpha1.Tenant) (*v1alpha1.Tenant, error) {
	result := v1alpha1.Tenant{}
	err := c.restClient.
		Post().
		Namespace(c.ns).
		Resource("tenants").
		Body(project).
		Do().
		Into(&result)

	return &result, err
}

func (c *tenantClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namespace(c.ns).
		Resource("tenants").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}
