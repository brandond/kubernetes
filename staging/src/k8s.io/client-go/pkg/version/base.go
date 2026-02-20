package version

var (
	gitMajor = "1"
	gitMinor = "35"
	gitVersion   = "v1.35.1-bd4"
	gitCommit    = "8561d9cd8b3f14f228746fa220c9fbb88a2714b1"
	gitTreeState = "clean"
	buildDate = "2026-02-20T08:22:38Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.35"
)
