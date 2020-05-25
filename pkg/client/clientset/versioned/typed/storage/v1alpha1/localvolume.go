/*
Copyright 2019 JD Cloud

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/kubernetes-local-volume/kubernetes-local-volume/pkg/apis/storage/v1alpha1"
	scheme "github.com/kubernetes-local-volume/kubernetes-local-volume/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// LocalVolumesGetter has a method to return a LocalVolumeInterface.
// A group's client should implement this interface.
type LocalVolumesGetter interface {
	LocalVolumes(namespace string) LocalVolumeInterface
}

// LocalVolumeInterface has methods to work with LocalVolume resources.
type LocalVolumeInterface interface {
	Create(*v1alpha1.LocalVolume) (*v1alpha1.LocalVolume, error)
	Update(*v1alpha1.LocalVolume) (*v1alpha1.LocalVolume, error)
	UpdateStatus(*v1alpha1.LocalVolume) (*v1alpha1.LocalVolume, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.LocalVolume, error)
	List(opts v1.ListOptions) (*v1alpha1.LocalVolumeList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LocalVolume, err error)
	LocalVolumeExpansion
}

// localVolumes implements LocalVolumeInterface
type localVolumes struct {
	client rest.Interface
	ns     string
}

// newLocalVolumes returns a LocalVolumes
func newLocalVolumes(c *LocalV1alpha1Client, namespace string) *localVolumes {
	return &localVolumes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the localVolume, and returns the corresponding localVolume object, and an error if there is any.
func (c *localVolumes) Get(name string, options v1.GetOptions) (result *v1alpha1.LocalVolume, err error) {
	result = &v1alpha1.LocalVolume{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("localvolumes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LocalVolumes that match those selectors.
func (c *localVolumes) List(opts v1.ListOptions) (result *v1alpha1.LocalVolumeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LocalVolumeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("localvolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested localVolumes.
func (c *localVolumes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("localvolumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a localVolume and creates it.  Returns the server's representation of the localVolume, and an error, if there is any.
func (c *localVolumes) Create(localVolume *v1alpha1.LocalVolume) (result *v1alpha1.LocalVolume, err error) {
	result = &v1alpha1.LocalVolume{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("localvolumes").
		Body(localVolume).
		Do().
		Into(result)
	return
}

// Update takes the representation of a localVolume and updates it. Returns the server's representation of the localVolume, and an error, if there is any.
func (c *localVolumes) Update(localVolume *v1alpha1.LocalVolume) (result *v1alpha1.LocalVolume, err error) {
	result = &v1alpha1.LocalVolume{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("localvolumes").
		Name(localVolume.Name).
		Body(localVolume).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *localVolumes) UpdateStatus(localVolume *v1alpha1.LocalVolume) (result *v1alpha1.LocalVolume, err error) {
	result = &v1alpha1.LocalVolume{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("localvolumes").
		Name(localVolume.Name).
		SubResource("status").
		Body(localVolume).
		Do().
		Into(result)
	return
}

// Delete takes name of the localVolume and deletes it. Returns an error if one occurs.
func (c *localVolumes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("localvolumes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *localVolumes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("localvolumes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched localVolume.
func (c *localVolumes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LocalVolume, err error) {
	result = &v1alpha1.LocalVolume{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("localvolumes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}