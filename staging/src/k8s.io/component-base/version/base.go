package version

var (
	gitMajor = "1"
	gitMinor = "35"
	gitVersion   = "v1.35.0-rc.1k3s1"
	gitCommit    = "b88e78a71013b6171663333c6ab2c56c9c2d7529"
	gitTreeState = "clean"
	buildDate = "2025-12-11T22:25:53Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.35"
)
