package v0alpha1

import (
	"context"

	"github.com/grafana/grafana-app-sdk/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SummaryClient struct {
	client *resource.TypedClient[*Summary, *SummaryList]
}

func NewSummaryClient(client resource.Client) *SummaryClient {
	return &SummaryClient{
		client: resource.NewTypedClient[*Summary, *SummaryList](client, SummaryKind()),
	}
}

func NewSummaryClientFromGenerator(generator resource.ClientGenerator) (*SummaryClient, error) {
	c, err := generator.ClientFor(SummaryKind())
	if err != nil {
		return nil, err
	}
	return NewSummaryClient(c), nil
}

func (c *SummaryClient) Get(ctx context.Context, identifier resource.Identifier) (*Summary, error) {
	return c.client.Get(ctx, identifier)
}

func (c *SummaryClient) List(ctx context.Context, namespace string, opts resource.ListOptions) (*SummaryList, error) {
	return c.client.List(ctx, namespace, opts)
}

func (c *SummaryClient) ListAll(ctx context.Context, namespace string, opts resource.ListOptions) (*SummaryList, error) {
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

func (c *SummaryClient) Create(ctx context.Context, obj *Summary, opts resource.CreateOptions) (*Summary, error) {
	// Make sure apiVersion and kind are set
	obj.APIVersion = GroupVersion.Identifier()
	obj.Kind = SummaryKind().Kind()
	return c.client.Create(ctx, obj, opts)
}

func (c *SummaryClient) Update(ctx context.Context, obj *Summary, opts resource.UpdateOptions) (*Summary, error) {
	return c.client.Update(ctx, obj, opts)
}

func (c *SummaryClient) Patch(ctx context.Context, identifier resource.Identifier, req resource.PatchRequest, opts resource.PatchOptions) (*Summary, error) {
	return c.client.Patch(ctx, identifier, req, opts)
}

func (c *SummaryClient) UpdateStatus(ctx context.Context, identifier resource.Identifier, newStatus SummaryStatus, opts resource.UpdateOptions) (*Summary, error) {
	return c.client.Update(ctx, &Summary{
		TypeMeta: metav1.TypeMeta{
			Kind:       SummaryKind().Kind(),
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

func (c *SummaryClient) Delete(ctx context.Context, identifier resource.Identifier, opts resource.DeleteOptions) error {
	return c.client.Delete(ctx, identifier, opts)
}
