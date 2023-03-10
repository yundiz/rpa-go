package heapprofiler

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"merkaba/chromedp/cdproto/runtime"
)

// HeapSnapshotObjectID heap snapshot object id.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler#type-HeapSnapshotObjectId
type HeapSnapshotObjectID string

// String returns the HeapSnapshotObjectID as string value.
func (t HeapSnapshotObjectID) String() string {
	return string(t)
}

// SamplingHeapProfileNode sampling Heap Profile node. Holds callsite
// information, allocation statistics and child nodes.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler#type-SamplingHeapProfileNode
type SamplingHeapProfileNode struct {
	CallFrame *runtime.CallFrame         `json:"callFrame"` // Function location.
	SelfSize  float64                    `json:"selfSize"`  // Allocations size in bytes for the node excluding children.
	ID        int64                      `json:"id"`        // Node id. Ids are unique across all profiles collected between startSampling and stopSampling.
	Children  []*SamplingHeapProfileNode `json:"children"`  // Child nodes.
}

// SamplingHeapProfileSample a single sample from a sampling profile.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler#type-SamplingHeapProfileSample
type SamplingHeapProfileSample struct {
	Size    float64 `json:"size"`    // Allocation size in bytes attributed to the sample.
	NodeID  int64   `json:"nodeId"`  // Id of the corresponding profile tree node.
	Ordinal float64 `json:"ordinal"` // Time-ordered sample ordinal number. It is unique across all profiles retrieved between startSampling and stopSampling.
}

// SamplingHeapProfile sampling profile.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler#type-SamplingHeapProfile
type SamplingHeapProfile struct {
	Head    *SamplingHeapProfileNode     `json:"head"`
	Samples []*SamplingHeapProfileSample `json:"samples"`
}
