package publisher

type Publisher interface {
	Init(map[string]string) error
	Name() string
	Version() string
	Publish(newRelease string) error
}
