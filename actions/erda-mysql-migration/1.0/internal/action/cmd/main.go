// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"os"

	"github.com/erda-project/erda/pkg/database/sqlparser/migrator"
	"github.com/sirupsen/logrus"

	"github.com/erda-project/erda-actions/actions/erda-mysql-migration/1.0/internal/action/migration"
	"github.com/erda-project/erda-actions/actions/erda-mysql-migration/1.0/internal/common"
	"github.com/erda-project/erda-actions/pkg/metawriter"
)

func main() {
	logrus.Infoln("Erda MySQL Migration start working")
	logrus.Infof("Configuration: %+v", *migration.Configuration())

	go common.FatalError(common.StartSandbox)

	mig, err := migrator.New(migration.Configuration())
	if err != nil {
		_ = metawriter.Write(map[string]interface{}{"success": false, "err": err})
		logrus.Fatalf("failed to start Erda MySQL Migration: %v", err)
	}
	if err = mig.Run(); err != nil {
		_ = metawriter.Write(map[string]interface{}{"success": false, "err": err})
		logrus.Fatalf("failed to migrate: %v", err)
	}

	logrus.Infoln("migrate complete !")
	_ = metawriter.Write(map[string]interface{}{"success": false})

	os.Exit(0)
}