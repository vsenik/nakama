// Copyright 2017 The Nakama Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"github.com/satori/go.uuid"

	"go.uber.org/zap"
)

type session interface {
	Logger() *zap.Logger
	ID() uuid.UUID
	UserID() uuid.UUID

	Handle() string
	SetHandle(string)

	Lang() string
	Expiry() int64

	Consume(func(logger *zap.Logger, session session, envelope *Envelope))
	Unregister()

	Send(envelope *Envelope) error
	SendBytes(payload []byte) error

	Close()
}
