// Copyright 2019 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"time"

	v1beta1 "github.com/vmware-tanzu/antrea/pkg/apis/networking/v1beta1"
	scheme "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AppliedToGroupsGetter has a method to return a AppliedToGroupInterface.
// A group's client should implement this interface.
type AppliedToGroupsGetter interface {
	AppliedToGroups() AppliedToGroupInterface
}

// AppliedToGroupInterface has methods to work with AppliedToGroup resources.
type AppliedToGroupInterface interface {
	Get(name string, options v1.GetOptions) (*v1beta1.AppliedToGroup, error)
	List(opts v1.ListOptions) (*v1beta1.AppliedToGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	AppliedToGroupExpansion
}

// appliedToGroups implements AppliedToGroupInterface
type appliedToGroups struct {
	client rest.Interface
}

// newAppliedToGroups returns a AppliedToGroups
func newAppliedToGroups(c *NetworkingV1beta1Client) *appliedToGroups {
	return &appliedToGroups{
		client: c.RESTClient(),
	}
}

// Get takes name of the appliedToGroup, and returns the corresponding appliedToGroup object, and an error if there is any.
func (c *appliedToGroups) Get(name string, options v1.GetOptions) (result *v1beta1.AppliedToGroup, err error) {
	result = &v1beta1.AppliedToGroup{}
	err = c.client.Get().
		Resource("appliedtogroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AppliedToGroups that match those selectors.
func (c *appliedToGroups) List(opts v1.ListOptions) (result *v1beta1.AppliedToGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.AppliedToGroupList{}
	err = c.client.Get().
		Resource("appliedtogroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested appliedToGroups.
func (c *appliedToGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("appliedtogroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}
