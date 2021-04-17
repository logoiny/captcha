// Copyright 2011 Dmitry Chestnykh. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package work

import (
	"fmt"
	"captcha/utils"
	"time"
)

// An object implementing Store interface can be registered with SetCustomStore
// function to handle storage and retrieval of work ids and solutions for
// them, replacing the default memory store.
//
// It is the responsibility of an object to delete expired and used captchas
// when necessary (for example, the default memory store collects them in Set
// method after the certain amount of captchas has been stored.)
type Store interface {
	// Set sets the digits for the work id.
	Set(id string, digits []byte, d time.Duration)

	// Get returns stored digits for the work id. Clear indicates
	// whether the work must be deleted from the store.
	Get(id string, clear bool) string
}

type RdsStore struct{}

func NewRdsStore() *RdsStore {
	return &RdsStore{}
}

func (s *RdsStore) Set(id string, digits []byte, d time.Duration) {

	bs := utils.DigitsToString(digits)
	if d < 0 {
		d = Expiration
	}
	err := rdb.Set(ctx, id, bs, d).Err()
	if err != nil {
		panic(err)
	}
}

func (s *RdsStore) Get(id string, clear bool) string {
	val, err := rdb.Get(ctx, id).Result()
	if err != nil &&  err.Error() != "redis: nil"{
		fmt.Println(err)
		panic(err)
	}
	if clear {
		rdb.Del(ctx, id)
	}
	fmt.Println(val)
	return val
}









