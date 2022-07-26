package v1

// V1Interface /v1 下的集合
type V1Interface interface {
	VersionGetter
}

// VersionGetter /version 下的集合
type VersionGetter interface {
	Version() VersionInterface
}
