package playwright

import (
	"errors"

	"github.com/playwright-community/playwright-go"
	"go.k6.io/k6/js/modules"
)

type locatorWrapper struct {
	Locator playwright.Locator
	vu      modules.VU
}

func newLocatorWrapper(locator playwright.Locator, vu modules.VU) *locatorWrapper {
	return &locatorWrapper{
		Locator: locator,
		vu:      vu,
	}
}

func (loc *locatorWrapper) Click(options ...playwright.PageClickOptions) {
	err := loc.Locator.Click(options...)
	if err != nil {
		Throw(loc.vu.Runtime(), "Error clicking element", err)
	}
}

func (loc *locatorWrapper) ScrollIntoViewIfNeeded(options ...playwright.LocatorScrollIntoViewIfNeededOptions) {
	err := loc.Locator.ScrollIntoViewIfNeeded(options...)
	if err != nil {
		Throw(loc.vu.Runtime(), "Error scrolling element into view", err)
	}
}

func (loc *locatorWrapper) Count() int {
	num, err := loc.Locator.Count()
	if err != nil {
		Throw(loc.vu.Runtime(), "Error counting elements", err)
	}
	return num
}

func (loc *locatorWrapper) Fill(value string, options ...playwright.FrameFillOptions) {
	err := loc.Locator.Fill(value, options...)
	if err != nil {
		Throw(loc.vu.Runtime(), "Error filling input", err)
	}
}

func (loc *locatorWrapper) First() *locatorWrapper {
	firstLocator := loc.Locator.First()
	if firstLocator == nil {
		msg := "Error getting first element"
		Throw(loc.vu.Runtime(), msg, errors.New(msg))
	}
	return newLocatorWrapper(firstLocator, loc.vu)
}

func (loc *locatorWrapper) IsDisabled(options ...playwright.FrameIsDisabledOptions) bool {
	isDisabled, err := loc.Locator.IsDisabled(options...)
	if err != nil {
		Throw(loc.vu.Runtime(), "Error checking if element is disabled", err)
	}
	return isDisabled
}

func (loc *locatorWrapper) IsVisible(options ...playwright.FrameIsVisibleOptions) bool {
	isVisible, err := loc.Locator.IsVisible(options...)
	if err != nil {
		Throw(loc.vu.Runtime(), "Error checking if element is visible", err)
	}
	return isVisible
}

func (loc *locatorWrapper) IsChecked(options ...playwright.FrameIsCheckedOptions) bool {
	isChecked, err := loc.Locator.IsChecked(options...)
	if err != nil {
		Throw(loc.vu.Runtime(), "Error checking if element is checked", err)
	}
	return isChecked
}

func (loc *locatorWrapper) IsEnabled(options ...playwright.FrameIsEnabledOptions) bool {
	isEnabled, err := loc.Locator.IsEnabled(options...)
	if err != nil {
		Throw(loc.vu.Runtime(), "Error checking if element is enabled", err)
	}
	return isEnabled
}

func (loc *locatorWrapper) Last() *locatorWrapper {
	lastLocator := loc.Locator.Last()
	if lastLocator == nil {
		msg := "Error getting last element"
		Throw(loc.vu.Runtime(), msg, errors.New(msg))
	}
	return newLocatorWrapper(lastLocator, loc.vu)
}

func (loc *locatorWrapper) Type(text string, options ...playwright.PageTypeOptions) {
	err := loc.Locator.Type(text, options...)
	if err != nil {
		Throw(loc.vu.Runtime(), "Error getting element type", err)
	}
}

func (loc *locatorWrapper) Nth(index int) *locatorWrapper {
	nthLocator := loc.Locator.Nth(index)
	if nthLocator == nil {
		msg := "Error getting the nth element"
		Throw(loc.vu.Runtime(), msg, errors.New(msg))
	}
	return newLocatorWrapper(nthLocator, loc.vu)
}

func (loc *locatorWrapper) Check() {
	err := loc.Locator.Check()
	if err != nil {
		Throw(loc.vu.Runtime(), "Error checking input field", err)
	}
}
