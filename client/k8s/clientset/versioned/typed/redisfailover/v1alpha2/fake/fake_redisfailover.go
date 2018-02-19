package fake

import (
	v1alpha2 "github.com/spotahome/redis-operator/api/redisfailover/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRedisFailovers implements RedisFailoverInterface
type FakeRedisFailovers struct {
	Fake *FakeStorageV1alpha2
	ns   string
}

var redisfailoversResource = schema.GroupVersionResource{Group: "storage.spotahome.com", Version: "v1alpha2", Resource: "redisfailovers"}

var redisfailoversKind = schema.GroupVersionKind{Group: "storage.spotahome.com", Version: "v1alpha2", Kind: "RedisFailover"}

// Get takes name of the redisFailover, and returns the corresponding redisFailover object, and an error if there is any.
func (c *FakeRedisFailovers) Get(name string, options v1.GetOptions) (result *v1alpha2.RedisFailover, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(redisfailoversResource, c.ns, name), &v1alpha2.RedisFailover{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.RedisFailover), err
}

// List takes label and field selectors, and returns the list of RedisFailovers that match those selectors.
func (c *FakeRedisFailovers) List(opts v1.ListOptions) (result *v1alpha2.RedisFailoverList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(redisfailoversResource, redisfailoversKind, c.ns, opts), &v1alpha2.RedisFailoverList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.RedisFailoverList{}
	for _, item := range obj.(*v1alpha2.RedisFailoverList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested redisFailovers.
func (c *FakeRedisFailovers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(redisfailoversResource, c.ns, opts))

}

// Create takes the representation of a redisFailover and creates it.  Returns the server's representation of the redisFailover, and an error, if there is any.
func (c *FakeRedisFailovers) Create(redisFailover *v1alpha2.RedisFailover) (result *v1alpha2.RedisFailover, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(redisfailoversResource, c.ns, redisFailover), &v1alpha2.RedisFailover{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.RedisFailover), err
}

// Update takes the representation of a redisFailover and updates it. Returns the server's representation of the redisFailover, and an error, if there is any.
func (c *FakeRedisFailovers) Update(redisFailover *v1alpha2.RedisFailover) (result *v1alpha2.RedisFailover, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(redisfailoversResource, c.ns, redisFailover), &v1alpha2.RedisFailover{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.RedisFailover), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRedisFailovers) UpdateStatus(redisFailover *v1alpha2.RedisFailover) (*v1alpha2.RedisFailover, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(redisfailoversResource, "status", c.ns, redisFailover), &v1alpha2.RedisFailover{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.RedisFailover), err
}

// Delete takes name of the redisFailover and deletes it. Returns an error if one occurs.
func (c *FakeRedisFailovers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(redisfailoversResource, c.ns, name), &v1alpha2.RedisFailover{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRedisFailovers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(redisfailoversResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha2.RedisFailoverList{})
	return err
}

// Patch applies the patch and returns the patched redisFailover.
func (c *FakeRedisFailovers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.RedisFailover, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(redisfailoversResource, c.ns, name, data, subresources...), &v1alpha2.RedisFailover{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.RedisFailover), err
}