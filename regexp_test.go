// Copyright 2012 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mux

import (
	"regexp"
	"testing"
)

func Test_buildRegexp(t *testing.T) {

	reg, err := regexp.Compile(".*")
	if err != nil {
		t.Error(err)
		return
	}
	cache := map[string]*regexp.Regexp{
		"a": reg,
	}
	reg1 := cache["a"]
	b := reg1.Match([]byte("a"))
	t.Log("1", b)
	delete(cache, "a")
	b = reg1.Match([]byte("a"))
	t.Log("2", b)

}
