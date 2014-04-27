// Copyright © 2014 Terry Mao, LiuDing All rights reserved.
// This file is part of gopush-cluster.

// gopush-cluster is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// gopush-cluster is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with gopush-cluster.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"encoding/json"
)

const (
	StorageTypeRedis = "redis"
	StorageTypeMysql = "mysql"
)

// The Message struct
type Message struct {
	Msg json.RawMessage // message content 
	MsgId int64 // message id
    GroupId int // group id
}

// Struct for delele message
type DelMessageInfo struct {
	Key  string
	Msgs []interface{}
}

var UseStorage Storage

// Stored messages interface
type Storage interface {
	// Save message
	Save(key string, msg json.RawMessage, mid int64, gid int, expire uint) error
	// Get messages
	Get(key string, mid int64) ([]json.RawMessage, error)
	// Delete key
	DelKey(key string) error
	// Delete multiple messages
	DelMulti(info *DelMessageInfo) error
}

// InitStorage init the storage type(mysql or redis).
func InitStorage() error {
	if Conf.StorageType == StorageTypeRedis {
		UseStorage = NewRedis()
	} else if Conf.StorageType == StorageTypeMysql {
		UseStorage = NewMYSQL()
	} else {
        glog.Errorf("unknown storage type: \"%s\"", Conf.StorageType)
        return ErrStorageType
    }
	return nil
}
