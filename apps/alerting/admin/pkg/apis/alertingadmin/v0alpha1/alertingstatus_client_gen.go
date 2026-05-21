package v0alpha1

import (
	"context"

	"github.com/grafana/grafana-app-sdk/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AlertingStatusClient struct {
	client *resource.TypedClient[*AlertingStatus, *AlertingStatusList]
}

func NewAlertingStatusClient(client resource.Client) *AlertingStatusClient {
	return &AlertingStatusClient{
		client: resource.NewTypedClient[*AlertingStatus, *AlertingStatusList](client, AlertingStatusKind()),
	}
}

func NewAlertingStatusClientFromGenerator(generator resource.ClientGenerator) (*AlertingStatusClient, error) {
	c, err := generator.ClientFor(AlertingStatusKind())
	if err != nil {
		return nil, err
	}
	return NewAlertingStatusClient(c), nil
}

func (c *AlertingStatusClient) Get(ctx context.Context, identifier resource.Identifier) (*AlertingStatus, error) {
	return c.client.Get(ctx, identifier)
}

func (c *AlertingStatusClient) List(ctx context.Context, namespace string, opts resource.ListOptions) (*AlertingStatusList, error) {
	return c.client.List(ctx, namespace, opts)
}

func (c *AlertingStatusClient) ListAll(ctx context.Context, namespace string, opts resource.ListOptions) (*AlertingStatusList, error) {
	resp, err := c.client.List(ctx, namespace, resource.ListOptions{
		ResourceVersion: opts.ResourceVersion,
		Limit:           opts.Limit,
		LabelFilters:    opts.LabelFilters,
		FieldSelectors:  opts.FieldSelectors,
	})
	if err != nil {
		return nil, err
	}
	for resp.GetContinue() != "" {
		page, err := c.client.List(ctx, namespace, resource.ListOptions{
			Continue:        resp.GetContinue(),
			ResourceVersion: opts.ResourceVersion,
			Limit:           opts.Limit,
			LabelFilters:    opts.LabelFilters,
			FieldSelectors:  opts.FieldSelectors,
		})
		if err != nil {
			return nil, err
		}
		resp.SetContinue(page.GetContinue())
		resp.SetResourceVersion(page.GetResourceVersion())
		resp.SetItems(append(resp.GetItems(), page.GetItems()...))
	}
	return resp, nil
}

func (c *AlertingStatusClient) Create(ctx context.Context, obj *AlertingStatus, opts resource.CreateOptions) (*AlertingStatus, error) {
	// Make sure apiVersion and kind are set
	obj.APIVersion = GroupVersion.Identifier()
	obj.Kind = AlertingStatusKind().Kind()
	return c.client.Create(ctx, obj, opts)
}

func (c *AlertingStatusClient) Update(ctx context.Context, obj *AlertingStatus, opts resource.UpdateOptions) (*AlertingStatus, error) {
	return c.client.Update(ctx, obj, opts)
}

func (c *AlertingStatusClient) Patch(ctx context.Context, identifier resource.Identifier, req resource.PatchRequest, opts resource.PatchOptions) (*AlertingStatus, error) {
	return c.client.Patch(ctx, identifier, req, opts)
}

func (c *AlertingStatusClient) UpdateStatus(ctx context.Context, identifier resource.Identifier, newStatus AlertingStatusStatus, opts resource.UpdateOptions) (*AlertingStatus, error) {
	return c.client.Update(ctx, &AlertingStatus{
		TypeMeta: metav1.TypeMeta{
			Kind:       AlertingStatusKind().Kind(),
			APIVersion: GroupVersion.Identifier(),
		},
		ObjectMeta: metav1.ObjectMeta{
			ResourceVersion: opts.ResourceVersion,
			Namespace:       identifier.Namespace,
			Name:            identifier.Name,
		},
		Status: newStatus,
	}, resource.UpdateOptions{
		Subresource:     "status",
		ResourceVersion: opts.ResourceVersion,
	})
}

func (c *AlertingStatusClient) Delete(ctx context.Context, identifier resource.Identifier, opts resource.DeleteOptions) error {
	return c.client.Delete(ctx, identifier, opts)
}
