// Copyright 2011,2012 Johann Höchtl. All rights reserved.
// Use of this source code is governed by a Modified BSD License
// that can be found in the LICENSE file.

// RESTFul interface for coordinate transformations.

// +build !appengine
package main

import "net/http"

func main() {

	binding := binding()

	http.ListenAndServe(binding, nil)
}
