package v0alpha1

// routes declares the custom subroutes exposed by the alerting-admin app
// alongside its kinds. The `/status` aggregate returns the singleton status
// objects across the kinds in this group as a single composite payload —
// useful for clients that want everything in one shot. Per-kind listing and
// GET (e.g. /externalalertmanagersyncs) remain the standard k8s API and
// don't need to be declared here.
//
// When adding a new status kind, extend the response with an optional field
// whose key is the camelCased kind name and update the handler in
// pkg/app/app.go to populate it.
routes: {
	namespaced: {
		"/status": {
			"GET": {
				name: "getStatus"
				response: {
					body: [string]: _
				}
				responseMetadata: typeMeta: false
			}
		}
	}
}
