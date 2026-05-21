package v0alpha1

// SummarySpec is intentionally empty. Summary is a read-only synthetic
// kind; clients never write to it. The k8s machinery requires every kind
// to declare a spec type with at least one field, so we carry a single
// reserved placeholder. Writers will be rejected by the synthetic
// rest.Storage in pkg/app/summary_storage.go.
SummarySpec: {
	// Reserved. Writers are rejected by storage; readers should ignore.
	reserved?: string
}
