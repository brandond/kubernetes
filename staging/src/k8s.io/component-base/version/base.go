package version

var (
	gitMajor = "1"
	gitMinor = "35"
	gitVersion   = "v1.35.1-bd2"
	gitCommit    = "58f1388c3549674f77503b2c5b170ec150c88ca3"
	gitTreeState = "clean"
	buildDate = "2026-02-20T05:22:52Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.35"
)
