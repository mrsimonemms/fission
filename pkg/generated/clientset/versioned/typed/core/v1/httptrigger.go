/*
Copyright The Fission Authors.

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

package v1

import (
	"context"

	v1 "github.com/fission/fission/pkg/apis/core/v1"
	corev1 "github.com/fission/fission/pkg/generated/applyconfiguration/core/v1"
	scheme "github.com/fission/fission/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// HTTPTriggersGetter has a method to return a HTTPTriggerInterface.
// A group's client should implement this interface.
type HTTPTriggersGetter interface {
	HTTPTriggers(namespace string) HTTPTriggerInterface
}

// HTTPTriggerInterface has methods to work with HTTPTrigger resources.
type HTTPTriggerInterface interface {
	Create(ctx context.Context, _hTTPTrigger *v1.HTTPTrigger, opts metav1.CreateOptions) (*v1.HTTPTrigger, error)
	Update(ctx context.Context, _hTTPTrigger *v1.HTTPTrigger, opts metav1.UpdateOptions) (*v1.HTTPTrigger, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.HTTPTrigger, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.HTTPTriggerList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.HTTPTrigger, err error)
	Apply(ctx context.Context, _hTTPTrigger *corev1.HTTPTriggerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.HTTPTrigger, err error)
	HTTPTriggerExpansion
}

// hTTPTriggers implements HTTPTriggerInterface
type hTTPTriggers struct {
	*gentype.ClientWithListAndApply[*v1.HTTPTrigger, *v1.HTTPTriggerList, *corev1.HTTPTriggerApplyConfiguration]
}

// newHTTPTriggers returns a HTTPTriggers
func newHTTPTriggers(c *CoreV1Client, namespace string) *hTTPTriggers {
	return &hTTPTriggers{
		gentype.NewClientWithListAndApply[*v1.HTTPTrigger, *v1.HTTPTriggerList, *corev1.HTTPTriggerApplyConfiguration](
			"httptriggers",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.HTTPTrigger { return &v1.HTTPTrigger{} },
			func() *v1.HTTPTriggerList { return &v1.HTTPTriggerList{} }),
	}
}
