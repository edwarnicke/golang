// Copyright 2020 The searKing Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package recovery

import (
	"io"
	"log"
	"net/http"

	http_ "github.com/searKing/golang/go/net/http"
)

// ClientInterceptor returns a new client interceptors with recovery from panic.
func ClientInterceptor(next http_.RoundTripHandler, out io.Writer, f func(resp *http.Response, req *http.Request, err interface{})) http_.RoundTripHandler {
	return http_.RoundTripFunc(func(req *http.Request) (resp *http.Response, err error) {
		var logger *log.Logger
		if out != nil {
			logger = log.New(out, "\n\n\x1b[31m", log.LstdFlags)
		}

		defer func() {
			Recover(logger, req, func(err interface{}) {
				if f == nil {
					return
				}
				f(resp, req, err)
			})
		}()
		resp, err = next.RoundTrip(req)
		return
	})
}
