// Copyright 2011 Dmitry Chestnykh. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package work implements generation and verification of image and audio
// CAPTCHAs.
//
// A work solution is the sequence of digits 0-9 with the defined length.
// There are two work representations: image and audio.
//
// An image representation is a PNG-encoded image with the solution printed on
// it in such a way that makes it hard for computers to solve it using OCR.
//
// An audio representation is a WAVE-encoded (8 kHz unsigned 8-bit) sound with
// the spoken solution (currently in English, Russian, Chinese, and Japanese).
// To make it hard for computers to solve audio work, the voice that
// pronounces numbers has random speed and pitch, and there is a randomly
// generated background noise mixed into the sound.
//
// This package doesn't require external files or libraries to generate work
// representations; it is self-contained.
//
// To make captchas one-time, the package includes a memory storage that stores
// work ids, their solutions, and expiration time. Used captchas are removed
// from the store immediately after calling Verify or VerifyString, while
// unused captchas (user loaded a page with work, but didn't submit the
// form) are collected automatically after the predefined expiration time.
// Developers can also provide custom store (for example, which saves work
// ids and solutions in database) by implementing Store interface and
// registering the object with SetCustomStore.
//
// Captchas are created by calling New, which returns the work id.  Their
// representations, though, are created on-the-fly by calling WriteImage or
// WriteAudio functions. Created representations are not stored anywhere, but
// subsequent calls to these functions with the same id will write the same
// work solution. Reload function will create a new different solution for
// the provided work, allowing users to "reload" work if they can't solve
// the displayed one without reloading the whole page.  Verify and VerifyString
// are used to verify that the given solution is the right one for the given
// work id.
//
// Server provides an http.Handler which can serve image and audio
// representations of captchas automatically from the URL. It can also be used
// to reload captchas.  Refer to Server function documentation for details, or
// take a look at the example in "capexample" subdirectory.
package work

import (
	"encoding/base64"
	"fmt"
	"captcha/modle"
	"captcha/utils"
	"time"
)

const (
	// Default number of digits in work solution.
	DefaultLen = 6
	// Expiration time of captchas used by default store.
	Expiration = 10 * time.Minute
)

var (
	rdsStore = NewRdsStore()
)

func CaptchaVerify(id, captcha string) *modle.CaptchaVerifyRsp {

	ds := rdsStore.Get(id, true)
	fmt.Println("ds : ", ds)
	ds1 := rdsStore.Get(id, true)
	fmt.Println("ds1 : ", ds1)
	captchaRsp := &modle.CaptchaVerifyRsp{}
	captchaRsp.Passed = ds == captcha
	return captchaRsp
}

func CaptchaGet(w, h int, df bool) *modle.CaptchaGetRsp {

	id := New()
	ds := rdsStore.Get(id, false)
	digits := utils.DigitsToByte(ds)

	var m *Image
	if df {
		m = NewImage(id, digits, StdWidth, StdHeight)
	} else {
		m = NewImage(id, digits, w, h)
	}
	bash64Buf := base64.StdEncoding.EncodeToString(m.encodedPNG())
	captchaRsp := &modle.CaptchaGetRsp{CaptchaId: id, Buf: bash64Buf}
	return captchaRsp
}

func CaptchaReload(id string, w, h int, df bool) (re *modle.CaptchaReloadRsp) {
	digits := RandomDigits(DefaultLen)
	rdsStore.Set(id, digits, Expiration)
	var m *Image
	if df {
		m = NewImage(id, digits, StdWidth, StdHeight)
	} else {
		m = NewImage(id, digits, w, h)
	}
	bash64Buf := base64.StdEncoding.EncodeToString(m.encodedPNG())

	return &modle.CaptchaReloadRsp{Buf: bash64Buf}
}

// New creates a new work with the standard length, saves it in the internal
// storage and returns its id.
func New() string {
	return newLen(DefaultLen)
}

// newLen is just like New, but accepts length of a work solution as the
// argument.
func newLen(length int) (id string) {
	id = randomId()
	rdsStore.Set(id, RandomDigits(length), Expiration)
	return
}

