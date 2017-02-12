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
package controllers

import (
	"github.com/revel/revel"
	"github.com/msgrasser/api/app"
	"github.com/msgrasser/api/app/models"
	"github.com/jgraham909/revmgo"
	"labix.org/v2/mgo/bson"
	"encoding/json"
	"encoding/hex"
	"code.google.com/p/go-uuid/uuid"
)

type Application struct {
	*revel.Controller
	revmgo.MongoController
}

func (c Application) ListSyncableObjects() revel.Result {
	var test string
	c.Params.Bind(&test, "test")
	//return c.RenderText(test)
	
	syncableObjects := models.GetAllSyncableObjects(c.MongoSession)
	
	return c.RenderJson(syncableObjects)
}

func (c Application) CreateSyncableObject(objType string) revel.Result {
	
	key_values := make(map[string]interface{})

	// make sure a schema with that object type exists, or throw error
	var schema interface{}
	err := c.MongoSession.DB("landline").C("Schemas").Find(bson.M{"object_key":objType}).One(&schema)
	if err != nil {
		return c.RenderText("SyncableObject type not found.")
	}

	// parse params and put into map
	for k, v := range c.Params.Values {
		if k != "objType" {
			// TODO: validate against fields of Schema, e.g.:
			// if schema.GetSchemaFields().contains(k)
			key_values[k] = v[0]
		}
	}
	
	// convert map to json to string of json
	key_values_map, err := json.Marshal(key_values)
	if err != nil {
		panic(err)
	}
	key_values_string := string(key_values_map[:])
	
	// find the deployment
	dbd := c.MongoSession.DB("landline").C("Deployments")
 	var deployment map[string]string
 	err = dbd.Find(bson.M{}).One(&deployment)

 	if deployment["enc_is_on"] == "True" {
		key := app.RAW_DEK
	    key_values_string = hex.EncodeToString(models.Encrypt([]byte(key), []byte(key_values_string)))
	    revel.INFO.Println(key_values_string)
	}

	syncableObject := models.SyncableObject{models.Model{}, uuid.New(), objType, key_values_string, 0}
	
	// Save to db
	dbc := c.MongoSession.DB("landline").C("SyncableObjects")
	err = dbc.Insert(syncableObject)
	if err != nil {
		panic(err)
	}
	
	return c.RenderJson(syncableObject)
}
