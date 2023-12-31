// Copyright (c) 2015 Uber Technologies, Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package relay contains relaying interfaces for external use.
//
// These interfaces are currently unstable, and aren't covered by the API
// backwards-compatibility guarantee.
package relay

import (
	"context"
	"time"

	"github.com/temporalio/tchannel-go/thrift/arg2"
)

// KeyVal is a key, val pair in arg2
type KeyVal struct {
	Key []byte
	Val []byte
}

// CallFrame is an interface that abstracts access to the call req frame.
type CallFrame interface {
	// TTL is the TTL of the underlying frame
	TTL() time.Duration
	// Caller is the name of the originating service.
	Caller() []byte
	// Service is the name of the destination service.
	Service() []byte
	// Method is the name of the method being called.
	Method() []byte
	// RoutingDelegate is the name of the routing delegate, if any.
	RoutingDelegate() []byte
	// RoutingKey may refer to an alternate traffic group instead of the
	// traffic group identified by the service name.
	RoutingKey() []byte
	// Arg2StartOffset returns the offset from start of payload to the
	// beginning of Arg2 in bytes.
	Arg2StartOffset() int
	// Arg2EndOffset returns the offset from start of payload to the end of
	// Arg2 in bytes, and hasMore to indicate if there are more frames and
	// Arg3 has not started (i.e. Arg2 is fragmented).
	Arg2EndOffset() (_ int, hasMore bool)
	// Arg2Iterator returns the iterator for reading Arg2 key value pair
	// of TChannel-Thrift Arg Scheme. If no iterator is available, return
	// io.EOF.
	Arg2Iterator() (arg2.KeyValIterator, error)
	// Arg2Append appends a key/val pair to arg2
	Arg2Append(key, val []byte)
}

// RespFrame is an interface that abstracts access to the CallRes frame
type RespFrame interface {
	// OK indicates whether the call was successful
	OK() bool

	// ArgScheme returns the scheme of the arg
	ArgScheme() []byte

	// Arg2IsFragmented indicates whether arg2 runs over the first frame
	Arg2IsFragmented() bool

	// Arg2 returns the raw arg2 payload
	Arg2() []byte
}

// Conn contains information about the underlying connection.
type Conn struct {
	// RemoteAddr is the remote address of the underlying TCP connection.
	RemoteAddr string

	// RemoteProcessName is the process name sent in the TChannel handshake.
	RemoteProcessName string

	// IsOutbound returns whether this connection is an outbound connection
	// initiated via the relay.
	IsOutbound bool

	// Context contains connection-specific context which can be accessed via
	// RelayHost.Start()
	Context context.Context
}

// RateLimitDropError is the error that should be returned from
// RelayHosts.Get if the request should be dropped silently.
// This is bit of a hack, because rate limiting of this nature isn't part of
// the actual TChannel protocol.
// The relayer will record that it has dropped the packet, but *won't* notify
// the client.
type RateLimitDropError struct{}

func (e RateLimitDropError) Error() string {
	return "frame dropped silently due to rate limiting"
}
