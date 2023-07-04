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

package v1

import (
	"context"
	time "time"

	informers "github.com/maistra/xns-informer/pkg/informers"
	routev1 "github.com/openshift/api/route/v1"
	versioned "github.com/openshift/client-go/route/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/route/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/client-go/route/listers/route/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// RouteInformer provides access to a shared informer and lister for
// Routes.
type RouteInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.RouteLister
}

type routeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespaces       informers.NamespaceSet
}

// NewRouteInformer constructs a new informer for Route type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRouteInformer(client versioned.Interface, namespaces informers.NamespaceSet,
	resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRouteInformer(client, namespaces, resyncPeriod, indexers, nil)
}

// NewFilteredRouteInformer constructs a new informer for Route type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRouteInformer(client versioned.Interface, namespaces informers.NamespaceSet,
	resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	newInformer := func(namespace string) cache.SharedIndexInformer {
		return cache.NewSharedIndexInformer(
			&cache.ListWatch{
				ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
					if tweakListOptions != nil {
						tweakListOptions(&options)
					}
					return client.RouteV1().Routes(namespace).List(context.TODO(), options)
				},
				WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
					if tweakListOptions != nil {
						tweakListOptions(&options)
					}
					return client.RouteV1().Routes(namespace).Watch(context.TODO(), options)
				},
			},
			&routev1.Route{},
			resyncPeriod,
			indexers,
		)
	}

	return informers.NewMultiNamespaceInformer(namespaces, resyncPeriod, newInformer)
}

func (f *routeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRouteInformer(client, f.namespaces, resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *routeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&routev1.Route{}, f.defaultInformer)
}

func (f *routeInformer) Lister() v1.RouteLister {
	return v1.NewRouteLister(f.Informer().GetIndexer())
}
