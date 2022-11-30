package playwright

import (
	"log"
	"fmt"
	"github.com/playwright-community/playwright-go"
)

type dialogWrapper struct {
	Dialog playwright.Dialog
}

func (d* dialogWrapper) Accept(texts ...string) {
	fmt.Print("Into the accept function")
	err := d.Dialog.Accept(texts...)
	if err != nil {
		fmt.Print("got error")
		log.Fatalf("error with accepting dialog: %v", err)
	}
}

func newDialogWrapper(dialog playwright.Dialog) *dialogWrapper {
	return &dialogWrapper {
		Dialog: dialog,
	}
}