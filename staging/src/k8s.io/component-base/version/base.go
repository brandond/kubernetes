package version

var (
	gitMajor = "1"
	gitMinor = "35"
	gitVersion   = "v1.35.1-bd3"
	gitCommit    = "88844eaceebdff9c7fe728097a4334e13733c876"
	gitTreeState = "clean"
	buildDate = "2026-02-20T08:03:45Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.35"
)
