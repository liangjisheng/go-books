package pattern

// 饿汉模式, 初始化时就会创建

type singleton struct {

}

var instance *singleton

func init()  {
	instance = new(singleton)
}

func GetInstance()  *singleton{
	return instance
}
