package main

import "fmt"

// 组件 定义接口
type Beverage interface {
	Serve() string
	Cost() float32
}

// 具体组件 实现组件接口，具有基本功能
type Coffee struct {
}

func (c *Coffee) Serve() string {
	return "simple coffee"
}
func (c *Coffee) Cost() float32 {
	return 5.0
}

// 装饰器 持有一个组件的引用
type CoffeeDecorator struct {
	beverage Beverage
}

func (m *CoffeeDecorator) Serve() string {
	return m.beverage.Serve()
}
func (m *CoffeeDecorator) Cost() float32 {
	return m.beverage.Cost()
}

type MilkDecorator struct {
	CoffeeDecorator
}

func (m *MilkDecorator) Serve() string {
	return m.CoffeeDecorator.Serve() + " with milk"
}
func (m *MilkDecorator) Cost() float32 {
	return m.CoffeeDecorator.Cost() + 2.0
}

type SugarDecorator struct {
	CoffeeDecorator
}

func (s *SugarDecorator) Serve() string {
	return s.CoffeeDecorator.Serve() + "with sugar"
}
func (s *SugarDecorator) Cost() float32 {
	return s.CoffeeDecorator.Cost() + 1.0
}

func main() {
	coffee := &Coffee{}
	milkCoffee := &MilkDecorator{CoffeeDecorator{coffee}}
	sugarCoffee := &SugarDecorator{CoffeeDecorator{coffee}}
	sugarAndMilkCoffee := &SugarDecorator{CoffeeDecorator{milkCoffee}}
	fmt.Println(milkCoffee.Serve(), milkCoffee.Cost())
	fmt.Println(sugarCoffee.Serve(), sugarCoffee.Cost())
	fmt.Println(sugarAndMilkCoffee.Serve(), sugarAndMilkCoffee.Cost())

}
