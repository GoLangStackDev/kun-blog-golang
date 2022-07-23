package server

type IClass interface {
	// Build 控制器必须实现的方法
	Build(core *KCore)
}
