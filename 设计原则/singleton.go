package 设计原则

var(
	SingletonObject *singleton
)

type singleton struct {}

func init() {
	SingletonObject = new(singleton)
}

func GetSingleton() *singleton {
	return SingletonObject
}