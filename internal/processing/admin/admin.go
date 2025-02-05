// GoToSocial
// Copyright (C) GoToSocial Authors admin@gotosocial.org
// SPDX-License-Identifier: AGPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package admin

import (
	"github.com/superseriousbusiness/gotosocial/internal/cleaner"
	"github.com/superseriousbusiness/gotosocial/internal/email"
	"github.com/superseriousbusiness/gotosocial/internal/gtsmodel"
	"github.com/superseriousbusiness/gotosocial/internal/media"
	"github.com/superseriousbusiness/gotosocial/internal/state"
	"github.com/superseriousbusiness/gotosocial/internal/transport"
	"github.com/superseriousbusiness/gotosocial/internal/typeutils"
)

type Processor struct {
	state               *state.State
	cleaner             *cleaner.Cleaner
	converter           *typeutils.Converter
	mediaManager        *media.Manager
	transportController transport.Controller
	emailSender         email.Sender

	// admin Actions currently
	// undergoing processing
	actions *Actions
}

func (p *Processor) Actions() *Actions {
	return p.actions
}

// New returns a new admin processor.
func New(state *state.State, converter *typeutils.Converter, mediaManager *media.Manager, transportController transport.Controller, emailSender email.Sender) Processor {
	return Processor{
		state:               state,
		cleaner:             cleaner.New(state),
		converter:           converter,
		mediaManager:        mediaManager,
		transportController: transportController,
		emailSender:         emailSender,

		actions: &Actions{
			r:     make(map[string]*gtsmodel.AdminAction),
			state: state,
		},
	}
}
