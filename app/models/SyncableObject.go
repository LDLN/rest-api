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
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type SyncableObject struct {
	Model          				`bson:",inline"`
	Uuid						string		`json:"uuid" bson:"uuid"`
	ObjectType 					string		`json:"object_type" bson:"object_type"`
	KeyValuePairs				string		`json:"key_value_pairs" bson:"key_value_pairs"`
	TimeModifiedSinceCreation	int			`json:"time_modified_since_creation" bson:"time_modified_since_creation"`
}

func GetAllSyncableObjects(s *mgo.Session) []SyncableObject {
	syncableObjects := []SyncableObject{}
	syncableObject := new(SyncableObject)
	query := Collection(syncableObject, s).Find(bson.M{})
	query.All(&syncableObjects)
	return syncableObjects
}