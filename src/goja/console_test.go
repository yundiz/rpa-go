package goja

import (
	"testing"
)

func TestConsole(t *testing.T) {
	vm := New()

	InitConsole(vm, nil)

	if c := vm.Get("console"); c == nil {
		t.Fatal("console not found")
	}

	if _, err := vm.RunString("console.log('')"); err != nil {
		t.Fatal("console.log() error", err)
	}

	if _, err := vm.RunString("console.error('')"); err != nil {
		t.Fatal("console.error() error", err)
	}

	if _, err := vm.RunString("console.warn('')"); err != nil {
		t.Fatal("console.warn() error", err)
	}
}

func TestConsoleWithPrinter(t *testing.T) {
	vm := New()

	var lastPrint string
	printer := PrinterFunc(func(leve string, s string) {
		lastPrint = s
	})

	InitConsole(vm, printer)

	if c := vm.Get("console"); c == nil {
		t.Fatal("console not found")
	}

	if _, err := vm.RunString("console.log('log')"); err != nil {
		t.Fatal("console.log() error", err)
	}
	if lastPrint != "log" {
		t.Fatal("lastPrint not 'log'", lastPrint)
	}

	if _, err := vm.RunString("console.error('error')"); err != nil {
		t.Fatal("console.error() error", err)
	}
	if lastPrint != "error" {
		t.Fatal("lastPrint not 'error'", lastPrint)
	}

	if _, err := vm.RunString("console.warn('warn')"); err != nil {
		t.Fatal("console.warn() error", err)
	}
	if lastPrint != "warn" {
		t.Fatal("lastPrint not 'warn'", lastPrint)
	}
}
