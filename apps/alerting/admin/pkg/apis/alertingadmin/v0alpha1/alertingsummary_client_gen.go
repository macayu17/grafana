package v0alpha1

import (
	"context"

	"github.com/grafana/grafana-app-sdk/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AlertingSummaryClient struct {
	client *resource.TypedClient[*AlertingSummary, *AlertingSummaryList]
}

func NewAlertingSummaryClient(client resource.Client) *AlertingSummaryClient {
	return &AlertingSummaryClient{
		client: resource.NewTypedClient[*AlertingSummary, *AlertingSummaryList](client, AlertingSummaryKind()),
	}
}

func NewAlertingSummaryClientFromGenerator(generator resource.ClientGenerator) (*AlertingSummaryClient, error) {
	c, err := generator.ClientFor(AlertingSummaryKind())
	if err != nil {
		return nil, err
	}
	return NewAlertingSummaryClient(c), nil
}

func (c *AlertingSummaryClient) Get(ctx context.Context, identifier resource.Identifier) (*AlertingSummary, error) {
	return c.client.Get(ctx, identifier)
}

func (c *AlertingSummaryClient) List(ctx context.Context, namespace string, opts resource.ListOptions) (*AlertingSummaryList, error) {
	return c.client.List(ctx, namespace, opts)
}

func (c *AlertingSummaryClient) ListAll(ctx context.Context, namespace string, opts resource.ListOptions) (*AlertingSummaryList, error) {
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

func (c *AlertingSummaryClient) Create(ctx context.Context, obj *AlertingSummary, opts resource.CreateOptions) (*AlertingSummary, error) {
	// Make sure apiVersion and kind are set
	obj.APIVersion = GroupVersion.Identifier()
	obj.Kind = AlertingSummaryKind().Kind()
	return c.client.Create(ctx, obj, opts)
}

func (c *AlertingSummaryClient) Update(ctx context.Context, obj *AlertingSummary, opts resource.UpdateOptions) (*AlertingSummary, error) {
	return c.client.Update(ctx, obj, opts)
}

func (c *AlertingSummaryClient) Patch(ctx context.Context, identifier resource.Identifier, req resource.PatchRequest, opts resource.PatchOptions) (*AlertingSummary, error) {
	return c.client.Patch(ctx, identifier, req, opts)
}

func (c *AlertingSummaryClient) UpdateStatus(ctx context.Context, identifier resource.Identifier, newStatus AlertingSummaryStatus, opts resource.UpdateOptions) (*AlertingSummary, error) {
	return c.client.Update(ctx, &AlertingSummary{
		TypeMeta: metav1.TypeMeta{
			Kind:       AlertingSummaryKind().Kind(),
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

func (c *AlertingSummaryClient) Delete(ctx context.Context, identifier resource.Identifier, opts resource.DeleteOptions) error {
	return c.client.Delete(ctx, identifier, opts)
}
