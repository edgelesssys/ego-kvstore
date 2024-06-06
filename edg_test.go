/*
Copyright (c) Edgeless Systems GmbH

SPDX-License-Identifier: AGPL-3.0-only
*/

package kvstore

import "github.com/edgelesssys/ego-kvstore/internal/edg"

func init() {
	edg.TestEnableRandomKey()
}

func (f *memFile) WriteApproved(p []byte) error {
	return f.Write(p)
}
