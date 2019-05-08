/*
 * Copyright 2019 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package v1alpha1

import (
	v1alpha1 "github.com/pivotal/build-service-system/pkg/apis/build/v1alpha1"
	scheme "github.com/pivotal/build-service-system/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CNBBuildsGetter has a method to return a CNBBuildInterface.
// A group's client should implement this interface.
type CNBBuildsGetter interface {
	CNBBuilds(namespace string) CNBBuildInterface
}

// CNBBuildInterface has methods to work with CNBBuild resources.
type CNBBuildInterface interface {
	Create(*v1alpha1.CNBBuild) (*v1alpha1.CNBBuild, error)
	Update(*v1alpha1.CNBBuild) (*v1alpha1.CNBBuild, error)
	UpdateStatus(*v1alpha1.CNBBuild) (*v1alpha1.CNBBuild, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CNBBuild, error)
	List(opts v1.ListOptions) (*v1alpha1.CNBBuildList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CNBBuild, err error)
	CNBBuildExpansion
}

// cNBBuilds implements CNBBuildInterface
type cNBBuilds struct {
	client rest.Interface
	ns     string
}

// newCNBBuilds returns a CNBBuilds
func newCNBBuilds(c *BuildV1alpha1Client, namespace string) *cNBBuilds {
	return &cNBBuilds{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cNBBuild, and returns the corresponding cNBBuild object, and an error if there is any.
func (c *cNBBuilds) Get(name string, options v1.GetOptions) (result *v1alpha1.CNBBuild, err error) {
	result = &v1alpha1.CNBBuild{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cnbbuilds").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CNBBuilds that match those selectors.
func (c *cNBBuilds) List(opts v1.ListOptions) (result *v1alpha1.CNBBuildList, err error) {
	result = &v1alpha1.CNBBuildList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cnbbuilds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cNBBuilds.
func (c *cNBBuilds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cnbbuilds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a cNBBuild and creates it.  Returns the server's representation of the cNBBuild, and an error, if there is any.
func (c *cNBBuilds) Create(cNBBuild *v1alpha1.CNBBuild) (result *v1alpha1.CNBBuild, err error) {
	result = &v1alpha1.CNBBuild{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cnbbuilds").
		Body(cNBBuild).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cNBBuild and updates it. Returns the server's representation of the cNBBuild, and an error, if there is any.
func (c *cNBBuilds) Update(cNBBuild *v1alpha1.CNBBuild) (result *v1alpha1.CNBBuild, err error) {
	result = &v1alpha1.CNBBuild{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cnbbuilds").
		Name(cNBBuild.Name).
		Body(cNBBuild).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cNBBuilds) UpdateStatus(cNBBuild *v1alpha1.CNBBuild) (result *v1alpha1.CNBBuild, err error) {
	result = &v1alpha1.CNBBuild{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cnbbuilds").
		Name(cNBBuild.Name).
		SubResource("status").
		Body(cNBBuild).
		Do().
		Into(result)
	return
}

// Delete takes name of the cNBBuild and deletes it. Returns an error if one occurs.
func (c *cNBBuilds) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cnbbuilds").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cNBBuilds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cnbbuilds").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cNBBuild.
func (c *cNBBuilds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CNBBuild, err error) {
	result = &v1alpha1.CNBBuild{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cnbbuilds").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
