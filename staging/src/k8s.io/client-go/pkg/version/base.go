package version

var (
	gitMajor = "1"
	gitMinor = "35"
	gitVersion   = "v1.35.1-bd1"
	gitCommit    = "a7df1eb91e03ef1e229676274e5f3ae104651f84"
	gitTreeState = "clean"
	buildDate = "2026-02-20T01:49:55Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.35"
)
