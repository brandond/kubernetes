package version

var (
	gitMajor = "1"
	gitMinor = "34"
	gitVersion   = "v1.34.5-bd1"
	gitCommit    = "1f21cdcbad4e72a351349aa80e3f70c742788dab"
	gitTreeState = "clean"
	buildDate = "2026-03-04T23:10:40Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.34"
)
