/*
Copyright Red Hat, Inc.

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

// Code generated by xns-informer-gen. DO NOT EDIT.

package extensions

import (
	extensionsv1beta1 "github.com/maistra/xns-informer/pkg/generated/kube/extensions/v1beta1"
	informers "github.com/maistra/xns-informer/pkg/informers"
	extensions "k8s.io/client-go/informers/extensions"
	v1beta1 "k8s.io/client-go/informers/extensions/v1beta1"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
)

type group struct {
	factory          internalinterfaces.SharedInformerFactory
	namespaces       informers.NamespaceSet
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespaces informers.NamespaceSet, tweakListOptions internalinterfaces.TweakListOptionsFunc) extensions.Interface {
	return &group{factory: f, namespaces: namespaces, tweakListOptions: tweakListOptions}
}

// V1beta1 returns a new v1beta1.Interface.
func (g *group) V1beta1() v1beta1.Interface {
	return extensionsv1beta1.New(g.factory, g.namespaces, g.tweakListOptions)
}
