package singleton


//单例模式的要点有三个：
//一是某个类只能有一个实例；
//二是它必须自行创建这个实例；
//三是它必须自行向整个系统提供这个实例。
//
//单例模式是一种对象创建型模式。
//单例模式又名单件模式或单态模式。

var(
	SingletonObject *singleton
)

type singleton struct {
	// what at all
}

func init() {
	SingletonObject = new(singleton)
}

func GetSingleton() *singleton {
	return SingletonObject
}