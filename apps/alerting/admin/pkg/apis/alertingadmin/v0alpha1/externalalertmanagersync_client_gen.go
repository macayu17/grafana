package v0alpha1

import (
	"context"

	"github.com/grafana/grafana-app-sdk/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ExternalAlertmanagerSyncClient struct {
	client *resource.TypedClient[*ExternalAlertmanagerSync, *ExternalAlertmanagerSyncList]
}

func NewExternalAlertmanagerSyncClient(client resource.Client) *ExternalAlertmanagerSyncClient {
	return &ExternalAlertmanagerSyncClient{
		client: resource.NewTypedClient[*ExternalAlertmanagerSync, *ExternalAlertmanagerSyncList](client, ExternalAlertmanagerSyncKind()),
	}
}

func NewExternalAlertmanagerSyncClientFromGenerator(generator resource.ClientGenerator) (*ExternalAlertmanagerSyncClient, error) {
	c, err := generator.ClientFor(ExternalAlertmanagerSyncKind())
	if err != nil {
		return nil, err
	}
	return NewExternalAlertmanagerSyncClient(c), nil
}

func (c *ExternalAlertmanagerSyncClient) Get(ctx context.Context, identifier resource.Identifier) (*ExternalAlertmanagerSync, error) {
	return c.client.Get(ctx, identifier)
}

func (c *ExternalAlertmanagerSyncClient) List(ctx context.Context, namespace string, opts resource.ListOptions) (*ExternalAlertmanagerSyncList, error) {
	return c.client.List(ctx, namespace, opts)
}

func (c *ExternalAlertmanagerSyncClient) ListAll(ctx context.Context, namespace string, opts resource.ListOptions) (*ExternalAlertmanagerSyncList, error) {
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

func (c *ExternalAlertmanagerSyncClient) Create(ctx context.Context, obj *ExternalAlertmanagerSync, opts resource.CreateOptions) (*ExternalAlertmanagerSync, error) {
	// Make sure apiVersion and kind are set
	obj.APIVersion = GroupVersion.Identifier()
	obj.Kind = ExternalAlertmanagerSyncKind().Kind()
	return c.client.Create(ctx, obj, opts)
}

func (c *ExternalAlertmanagerSyncClient) Update(ctx context.Context, obj *ExternalAlertmanagerSync, opts resource.UpdateOptions) (*ExternalAlertmanagerSync, error) {
	return c.client.Update(ctx, obj, opts)
}

func (c *ExternalAlertmanagerSyncClient) Patch(ctx context.Context, identifier resource.Identifier, req resource.PatchRequest, opts resource.PatchOptions) (*ExternalAlertmanagerSync, error) {
	return c.client.Patch(ctx, identifier, req, opts)
}

func (c *ExternalAlertmanagerSyncClient) UpdateStatus(ctx context.Context, identifier resource.Identifier, newStatus ExternalAlertmanagerSyncStatus, opts resource.UpdateOptions) (*ExternalAlertmanagerSync, error) {
	return c.client.Update(ctx, &ExternalAlertmanagerSync{
		TypeMeta: metav1.TypeMeta{
			Kind:       ExternalAlertmanagerSyncKind().Kind(),
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

func (c *ExternalAlertmanagerSyncClient) Delete(ctx context.Context, identifier resource.Identifier, opts resource.DeleteOptions) error {
	return c.client.Delete(ctx, identifier, opts)
}
