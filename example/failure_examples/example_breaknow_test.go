package failure_examples

import (
	"github.com/Jacky-zxj/allure-go"
	"github.com/pkg/errors"
	"testing"
)

func TestBreakNowAllure(t *testing.T) {
	allure.Test(t, allure.Description("Test with Allure error in it"), allure.Action(func() {
		allure.BreakNow(errors.New("A more serious error"))
		allure.Step(allure.Description("Step you're not supposed to see"), allure.Action(func() {}))
	}))
}
