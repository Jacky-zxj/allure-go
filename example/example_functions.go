package example

import (
	"github.com/Jacky-zxj/allure-go"
)

func doSomething() {
	allure.Step(allure.Description("Something"), allure.Action(func() {
		doSomethingNested()
	}))
}

func doSomethingNested() {
	allure.Step(allure.Description("Because we can!"), allure.Action(func() {}))
}
