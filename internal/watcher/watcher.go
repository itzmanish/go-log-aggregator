package watcher

type Watcher interface {
	Watch()
	Result() chan Result
	Close()
	String() string
}
