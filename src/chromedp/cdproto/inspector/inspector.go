// Package inspector provides the Chrome DevTools Protocol
// commands, types, and events for the Inspector domain.
//
// Generated by the cdproto-gen command.
package inspector

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"
	"merkaba/chromedp/cdproto/cdp"
)

// DisableParams disables inspector domain notifications.
type DisableParams struct{}

// Disable disables inspector domain notifications.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Inspector#method-disable
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes Inspector.disable against the provided context.
func (p *DisableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandDisable, nil, nil)
}

// EnableParams enables inspector domain notifications.
type EnableParams struct{}

// Enable enables inspector domain notifications.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Inspector#method-enable
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes Inspector.enable against the provided context.
func (p *EnableParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandEnable, nil, nil)
}

// Command names.
const (
	CommandDisable = "Inspector.disable"
	CommandEnable  = "Inspector.enable"
)
