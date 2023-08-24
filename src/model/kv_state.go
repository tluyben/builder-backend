// Copyright 2022 The ILLA Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"time"

	"github.com/google/uuid"
)

type KVState struct {
	ID        int       `json:"id" 		   gorm:"column:id;type:bigserial"`
	UID       uuid.UUID `json:"uid" 	   gorm:"column:uid;type:uuid;not null"`
	TeamID    int       `json:"teamID"    gorm:"column:team_id;type:bigserial"`
	StateType int       `json:"state_type" gorm:"column:state_type;type:bigint"`
	AppRefID  int       `json:"app_ref_id" gorm:"column:app_ref_id;type:bigint"`
	Version   int       `json:"version"    gorm:"column:version;type:bigint"`
	Key       string    `json:"key" 	   gorm:"column:key;type:text"`
	Value     string    `json:"value" 	   gorm:"column:value;type:jsonb"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp"`
	CreatedBy int       `json:"created_by" gorm:"column:created_by;type:bigint"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp"`
	UpdatedBy int       `json:"updated_by" gorm:"column:updated_by;type:bigint"`
}

func (kvstate *KVState) CleanID() {
	kvstate.ID = 0
}

func (kvstate *KVState) InitUID() {
	kvstate.UID = uuid.New()
}

func (kvstate *KVState) InitCreatedAt() {
	kvstate.CreatedAt = time.Now().UTC()
}

func (kvstate *KVState) InitUpdatedAt() {
	kvstate.UpdatedAt = time.Now().UTC()
}

func (kvstate *KVState) AppendNewVersion(newVersion int) {
	kvstate.CleanID()
	kvstate.InitUID()
	kvstate.Version = newVersion
}

func (kvstate *KVState) ExportID() int {
	return kvstate.ID
}
