/*
 *  Copyright 2014-2015 LDLN
 *
 *  This file is part of LDLN Base Station.
 *
 *  LDLN Base Station is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  any later version.
 *
 *  LDLN Base Station is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License
 *  along with LDLN Base Station.  If not, see <http://www.gnu.org/licenses/>.
 */
package app

import (
	"github.com/revel/revel"
	"log"
)

var (
	DB string
	RAW_DEK string
)

func AppInit() {
	RegisterDB()
	RegisterRawDek()
}

func RegisterDB() {
	var found bool
	if DB, found = revel.Config.String("landline.db"); !found {
		DB = "landline"
	}
}

func RegisterRawDek() {
	var found bool
	if RAW_DEK, found = revel.Config.String("landline.rawdek"); !found {
		log.Println("landline.raw_dek missing in app.conf!")
	}
}