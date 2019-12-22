package infra

import (
	"fmt"
	"reflect"

	"github.com/tietang/props/kvs"
)

// 应用程序
type BootApplication struct {
	IsTest     bool
	conf       kvs.ConfigSource
	starterCtx StarterContext
}

// 构造函数
func New(conf kvs.ConfigSource) *BootApplication {
	e := &BootApplication{conf: conf, starterCtx: StarterContext{}}
	e.starterCtx.SetProps(conf)
	return e
}

func (b *BootApplication) Start() {
	// 1. 初始化所有starter
	b.init()
	// 2. 安装starter
	b.setup()
	// 3. 启动starter
	b.start()
}

// 程序初始化
func (b *BootApplication) init() {
	for _, value := range GetStarters() {
		typ := reflect.TypeOf(value)
		fmt.Println(typ.String())
		value.Init(b.starterCtx)
	}
}

// 程序安装
func (b *BootApplication) setup() {
	fmt.Println("setup　starters", GetStarters())
	for _, v := range GetStarters() {
		typ := reflect.TypeOf(v)
		fmt.Println("setup", typ.String())
		v.Setup(b.starterCtx)
	}
}

// 程序开始运行, 开始调用
func (b *BootApplication) start() {
	fmt.Println("Starting starters ...", GetStarters())
	for i, v := range GetStarters() {
		typ := reflect.TypeOf(v)
		fmt.Println("starting: ", typ.String())
		if b.starterCtx.Props().GetBoolDefault("testing", false) {
			go v.Start(b.starterCtx)
			continue
		}
		fmt.Println("v.StartBlocking()", v.StartBlocking())
		if v.StartBlocking() {
			// 如果是最后一个阻塞, 直接启动
			if i+1 == len(GetStarters()) {
				v.Start(b.starterCtx)
			} else {
				// 阻塞的异步启动
				go v.Start(b.starterCtx)
			}
		} else {
			v.Start(b.starterCtx)
		}
	}

}
