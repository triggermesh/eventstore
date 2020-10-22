/*
Copyright (c) 2020 TriggerMesh Inc.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	eventstoresv1alpha1 "github.com/triggermesh/eventstore/pkg/apis/eventstores/v1alpha1"
	internalclientset "github.com/triggermesh/eventstore/pkg/generated/client/clientset/internalclientset"
	internalinterfaces "github.com/triggermesh/eventstore/pkg/generated/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/triggermesh/eventstore/pkg/generated/client/listers/eventstores/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// InMemoryStoreInformer provides access to a shared informer and lister for
// InMemoryStores.
type InMemoryStoreInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.InMemoryStoreLister
}

type inMemoryStoreInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewInMemoryStoreInformer constructs a new informer for InMemoryStore type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewInMemoryStoreInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredInMemoryStoreInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredInMemoryStoreInformer constructs a new informer for InMemoryStore type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredInMemoryStoreInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EventstoresV1alpha1().InMemoryStores(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EventstoresV1alpha1().InMemoryStores(namespace).Watch(context.TODO(), options)
			},
		},
		&eventstoresv1alpha1.InMemoryStore{},
		resyncPeriod,
		indexers,
	)
}

func (f *inMemoryStoreInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredInMemoryStoreInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *inMemoryStoreInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&eventstoresv1alpha1.InMemoryStore{}, f.defaultInformer)
}

func (f *inMemoryStoreInformer) Lister() v1alpha1.InMemoryStoreLister {
	return v1alpha1.NewInMemoryStoreLister(f.Informer().GetIndexer())
}