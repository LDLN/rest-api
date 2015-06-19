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
package models

import (
)

type Schema struct {
	ObjectKey		string			`json:"object_key" bson:"object_key"`
	ObjectLabel 	string			`json:"object_label" bson:"object_label"`
	Weight			int				`json:"weight" bson:"weight"`
	Schema			[]SchemaField	`json:"schema" bson:"schema"`
}

type SchemaField struct {
	Key				string 			`json:"key" bson:"key"`
	Label			string 			`json:"label" bson:"label"`
	Type			string 			`json:"type" bson:"type"`
	Weight			int 			`json:"weight" bson:"weight"`
	Required		bool 			`json:"required" bson:"required"`
	Default			string 			`json:"default" bson:"default"`
}