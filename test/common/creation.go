/*
Copyright 2019 The Knative Authors

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

package common

import (
	"github.com/knative/eventing/test/base/resources"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TODO(Fredy-Z): break this file into multiple files when it grows too large.

var coreAPIGroup = corev1.SchemeGroupVersion.Group
var coreAPIVersion = corev1.SchemeGroupVersion.Version
var rbacAPIGroup = rbacv1.SchemeGroupVersion.Group
var rbacAPIVersion = rbacv1.SchemeGroupVersion.Version

// CreateChannelOrFail will create a Channel Resource in Eventing.
func (client *Client) CreateChannelOrFail(name string, channelTypeMeta *metav1.TypeMeta, provisionerName string) {
	namespace := client.Namespace
	switch channelTypeMeta.Kind {
	case resources.ChannelKind:
		channel := resources.Channel(name, provisionerName)
		channels := client.Eventing.EventingV1alpha1().Channels(namespace)
		channel, err := channels.Create(channel)
		if err != nil {
			client.T.Fatalf("Failed to create channel %q: %v", name, err)
		}
		client.Cleaner.AddObj(channel)
	case resources.InMemoryChannelKind:
		channel := resources.InMemoryChannel(name)
		channels := client.Eventing.MessagingV1alpha1().InMemoryChannels(namespace)
		channel, err := channels.Create(channel)
		if err != nil {
			client.T.Fatalf("Failed to create %q %q: %v", channelTypeMeta.Kind, name, err)
		}
		client.Cleaner.AddObj(channel)
	case resources.KafkaChannelKind:
		channel := resources.KafkaChannel(name)
		channels := client.Kafka.MessagingV1alpha1().KafkaChannels(namespace)
		channel, err := channels.Create(channel)
		if err != nil {
			client.T.Fatalf("Failed to create %q %q: %v", channelTypeMeta.Kind, name, err)
		}
		client.Cleaner.AddObj(channel)
	case resources.NatssChannelKind:
		channel := resources.NatssChannel(name)
		channels := client.Natss.MessagingV1alpha1().NatssChannels(namespace)
		channel, err := channels.Create(channel)
		if err != nil {
			client.T.Fatalf("Failed to create %q %q: %v", channelTypeMeta.Kind, name, err)
		}
		client.Cleaner.AddObj(channel)
	}
}

// CreateChannelsOrFail will create a list of Channel Resources in Eventing.
func (client *Client) CreateChannelsOrFail(names []string, channelTypeMeta *metav1.TypeMeta, provisionerName string) {
	for _, name := range names {
		client.CreateChannelOrFail(name, channelTypeMeta, provisionerName)
	}
}

// CreateSubscriptionOrFail will create a Subscription or fail the test if there is an error.
func (client *Client) CreateSubscriptionOrFail(
	name, channelName string,
	channelTypeMeta *metav1.TypeMeta,
	options ...resources.SubscriptionOption,
) {
	namespace := client.Namespace
	subscription := resources.Subscription(name, channelName, channelTypeMeta, options...)

	subscriptions := client.Eventing.EventingV1alpha1().Subscriptions(namespace)
	// update subscription with the new reference
	subscription, err := subscriptions.Create(subscription)
	if err != nil {
		client.T.Fatalf("Failed to create subscription %q: %v", name, err)
	}
	client.Cleaner.AddObj(subscription)
}

// CreateSubscriptionsOrFail will create a list of Subscriptions with the same configuration except the name.
func (client *Client) CreateSubscriptionsOrFail(
	names []string,
	channelName string,
	channelTypeMeta *metav1.TypeMeta,
	options ...resources.SubscriptionOption,
) {
	for _, name := range names {
		client.CreateSubscriptionOrFail(name, channelName, channelTypeMeta, options...)
	}
}

// CreateBrokerOrFail will create a Broker or fail the test if there is an error.
func (client *Client) CreateBrokerOrFail(name, provisionerName string) {
	namespace := client.Namespace
	broker := resources.Broker(name, provisionerName)

	brokers := client.Eventing.EventingV1alpha1().Brokers(namespace)
	// update broker with the new reference
	broker, err := brokers.Create(broker)
	if err != nil {
		client.T.Fatalf("Failed to create broker %q: %v", name, err)
	}
	client.Cleaner.AddObj(broker)
}

// CreateBrokersOrFail will create a list of Brokers.
func (client *Client) CreateBrokersOrFail(names []string, provisionerName string) {
	for _, name := range names {
		client.CreateBrokerOrFail(name, provisionerName)
	}
}

// CreateTriggerOrFail will create a Trigger or fail the test if there is an error.
func (client *Client) CreateTriggerOrFail(name string, options ...resources.TriggerOption) {
	namespace := client.Namespace
	trigger := resources.Trigger(name, options...)

	triggers := client.Eventing.EventingV1alpha1().Triggers(namespace)
	// update trigger with the new reference
	trigger, err := triggers.Create(trigger)
	if err != nil {
		client.T.Fatalf("Failed to create trigger %q: %v", name, err)
	}
	client.Cleaner.AddObj(trigger)
}

// CreateCronJobSourceOrFail will create a CronJobSource or fail the test if there is an error.
func (client *Client) CreateCronJobSourceOrFail(
	name,
	schedule,
	data string,
	options ...resources.CronJobSourceOption,
) {
	namespace := client.Namespace
	cronJobSource := resources.CronJobSource(name, schedule, data, options...)

	cronJobSources := client.Eventing.SourcesV1alpha1().CronJobSources(namespace)
	// update cronJobSource with the new reference
	cronJobSource, err := cronJobSources.Create(cronJobSource)
	if err != nil {
		client.T.Fatalf("Failed to create cronjobsource %q: %v", name, err)
	}
	client.Cleaner.AddObj(cronJobSource)
}

// CreateContainerSourceOrFail will create a ContainerSource or fail the test if there is an error.
func (client *Client) CreateContainerSourceOrFail(
	name string,
	options ...resources.ContainerSourceOption,
) {
	namespace := client.Namespace
	containerSource := resources.ContainerSource(name, options...)

	containerSources := client.Eventing.SourcesV1alpha1().ContainerSources(namespace)
	// update containerSource with the new reference
	containerSource, err := containerSources.Create(containerSource)
	if err != nil {
		client.T.Fatalf("Failed to create containersource %q: %v", name, err)
	}
	client.Cleaner.AddObj(containerSource)
}

// WithService returns an option that creates a Service binded with the given pod.
func WithService(name string) func(*corev1.Pod, *Client) error {
	return func(pod *corev1.Pod, client *Client) error {
		namespace := pod.Namespace
		svc := resources.Service(name, pod.Labels)

		svcs := client.Kube.Kube.CoreV1().Services(namespace)
		if _, err := svcs.Create(svc); err != nil {
			return err
		}
		client.Cleaner.Add(coreAPIGroup, coreAPIVersion, "services", namespace, name)
		return nil
	}
}

// CreatePodOrFail will create a Pod or fail the test if there is an error.
func (client *Client) CreatePodOrFail(pod *corev1.Pod, options ...func(*corev1.Pod, *Client) error) {
	// set namespace for the pod in case it's empty
	namespace := client.Namespace
	pod.Namespace = namespace
	// apply options on the pod before creation
	for _, option := range options {
		if err := option(pod, client); err != nil {
			client.T.Fatalf("Failed to configure pod %q: %v", pod.Name, err)
		}
	}
	if _, err := client.Kube.CreatePod(pod); err != nil {
		client.T.Fatalf("Failed to create pod %q: %v", pod.Name, err)
	}
	client.Cleaner.Add(coreAPIGroup, coreAPIVersion, "pods", namespace, pod.Name)
}

// CreateServiceAccountAndBindingOrFail creates both ServiceAccount and ClusterRoleBinding with default
// cluster-admin role.
func (client *Client) CreateServiceAccountAndBindingOrFail(saName, crName string) {
	namespace := client.Namespace
	sa := resources.ServiceAccount(saName, namespace)
	sas := client.Kube.Kube.CoreV1().ServiceAccounts(namespace)
	if _, err := sas.Create(sa); err != nil {
		client.T.Fatalf("Failed to create service account %q: %v", saName, err)
	}
	client.Cleaner.Add(coreAPIGroup, coreAPIVersion, "serviceaccounts", namespace, saName)

	crb := resources.ClusterRoleBinding(saName, crName, namespace)
	crbs := client.Kube.Kube.RbacV1().ClusterRoleBindings()
	if _, err := crbs.Create(crb); err != nil {
		client.T.Fatalf("Failed to create cluster role binding %q: %v", crName, err)
	}
	client.Cleaner.Add(rbacAPIGroup, rbacAPIVersion, "clusterrolebindings", "", crb.GetName())
}