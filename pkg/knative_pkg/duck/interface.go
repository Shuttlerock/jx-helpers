package duck

import "github.com/shuttlerock/jx-helpers/v3/pkg/knative_pkg/kmeta"

// OneOfOurs is the union of our Accessor interface and the OwnerRefable interface
// that is implemented by our resources that implement the kmeta.Accessor.
type OneOfOurs interface {
	kmeta.Accessor
	kmeta.OwnerRefable
}
