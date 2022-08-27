package v1

// V1Interface /v1 下的集合
type V1Interface interface {
	VersionGetter
	PostsGetter
}

// VersionGetter /version 下的集合
type VersionGetter interface {
	Version() VersionInterface
}

// PostsGetter /posts 下的结合
type PostsGetter interface {
	Posts() PostsInterface
}
