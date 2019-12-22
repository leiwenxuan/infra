package infra

import (
	"fmt"
	"reflect"

	"github.com/tietang/props/kvs"
)

// 常数
const KeyProps = "_conf"

// 基础资源上下结构体
type StarterContext map[string]interface{}

func (s StarterContext) Props() kvs.ConfigSource {
	p := s[KeyProps]
	if p == nil {
		panic("配置还没有初始化")
	}
	return p.(kvs.ConfigSource)
}
func (s StarterContext) SetProps(conf kvs.ConfigSource) {
	s[KeyProps] = conf
}

//基础资源启动器接口

type Starter interface {
	// 1. 系统启动， 初始化一些基础资源
	Init(StarterContext)
	// 2. 系统资源的安装
	Setup(StarterContext)
	// 启动基础资源
	Start(StarterContext)
	// 3. 启动器是否可阻塞
	StartBlocking() bool
	// 4. 基础资源的停止和销毁
	Stop(StarterContext)
}

// 基础空启动器的实现， 为了方便资源启动器的代码实现
type BaseStarter struct {
}

// 语法实现
var _ Starter = new(BaseStarter)

func (b *BaseStarter) Init(ctx StarterContext) {

}
func (b BaseStarter) Setup(StarterContext) {
	panic("implement me")
}

func (b BaseStarter) Start(StarterContext) {
	panic("implement me")
}

func (b BaseStarter) StartBlocking() bool {
	panic("implement me")
}

func (b BaseStarter) Stop(StarterContext) {
	panic("implement me")
}

// 服务启动注册器
type starterRegister struct {
	nonBlockingStarters []Starter
	blockingStarters    []Starter
}

// 启动器注册
func (r *starterRegister) Register(s Starter) {
	if s.StartBlocking() {
		r.blockingStarters = append(r.blockingStarters, s)
	} else {
		r.nonBlockingStarters = append(r.nonBlockingStarters, s)
	}
	typ := reflect.TypeOf(s)
	fmt.Println(typ.String())
}

func (r *starterRegister) AllStarters() []Starter {
	starters := make([]Starter, 0)
	starters = append(starters, r.nonBlockingStarters...)
	starters = append(starters, r.blockingStarters...)

	return starters
}

var StarterRegister *starterRegister = new(starterRegister)

// 注册starter
func Register(s Starter) {
	StarterRegister.Register(s)
}

// 获取所有注册的starter
func GetStarters() []Starter {
	return StarterRegister.AllStarters()
}
