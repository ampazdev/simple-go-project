package bridge

type ConfigReader interface {
	ReadFile(dest interface{}, path ...string) error
}


