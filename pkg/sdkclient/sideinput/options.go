/*
Copyright 2022 The Numaproj Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sideinput

import "time"

type options struct {
	sockAddr         string
	maxMessageSize   int
	sideInputTimeout time.Duration
}

// Option is the interface to apply options.
type Option func(*options)

// WithSockAddr start the client with the given sock addr. This is mainly used for testing purpose.
func WithSockAddr(addr string) Option {
	return func(opts *options) {
		opts.sockAddr = addr
	}
}

// WithMaxMessageSize sets the max message size to the given size.
func WithMaxMessageSize(size int) Option {
	return func(o *options) {
		o.maxMessageSize = size
	}
}

// WithSideInputTimeout sets the side input timeout to the given timeout.
func WithSideInputTimeout(t time.Duration) Option {
	return func(o *options) {
		o.sideInputTimeout = t
	}
}
