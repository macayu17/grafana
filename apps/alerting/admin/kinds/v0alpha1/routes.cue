package v0alpha1

// routes declares the custom subroutes exposed by the alerting-admin app
// alongside its kinds. The `/summary` aggregate returns observation data
// across the status kinds in this group as a single composite payload,
// grouped by area of concern to mirror the structure of Config.spec.
// Useful for clients that want everything in one shot.
//
// Per-kind listing and GET (e.g. /externalalertmanagersyncs) remain the
// standard k8s API and don't need to be declared here. The aggregate is
// deliberately not at `/status` to avoid shadowing the k8s status-subresource
// pattern (`/<plural>/<name>/status`), which has different semantics.
//
// When adding a new status kind, place its observation payload under the
// matching area in the response (the same key path used by Config.spec)
// and update the handler in pkg/app/app.go to populate it.
routes: {
	namespaced: {
		"/summary": {
			"GET": {
				name: "getSummary"
				response: {
					body: [string]: _
				}
				responseMetadata: typeMeta: false
			}
		}
	}
}
