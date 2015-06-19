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
	"github.com/msgrasser/api/app"
	"github.com/revel/revel"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"reflect"
	"io"
	"encoding/base64"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

type Model struct {
	Id				bson.ObjectId 	`bson:"_id,omitempty"`
}

func Collection(m interface{}, s *mgo.Session) *mgo.Collection {
	typ := reflect.TypeOf(m).Elem()
	n := typ.Name()

	var found bool
	var c string
	if c, found = revel.Config.String("landline.db.collection." + n); !found {
		c = n
	}
	return s.DB(app.DB).C(c)
}

func Encrypt(key, text []byte) []byte {
        block, err := aes.NewCipher(key)
        if err != nil {
                panic(err)
        }
        b := EncodeBase64(text)
        ciphertext := make([]byte, aes.BlockSize+len(b))
        iv := ciphertext[:aes.BlockSize]
        if _, err := io.ReadFull(rand.Reader, iv); err != nil {
                panic(err)
        }
        cfb := cipher.NewCFBEncrypter(block, iv)
        cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
        return ciphertext
}

func Decrypt(key, text []byte) []byte {
        block, err := aes.NewCipher(key)
        if err != nil {
                panic(err)
        }
        if len(text) < aes.BlockSize {
                panic("ciphertext too short")
        }
        iv := text[:aes.BlockSize]
        text = text[aes.BlockSize:]
        cfb := cipher.NewCFBDecrypter(block, iv)
        cfb.XORKeyStream(text, text)
        return DecodeBase64(string(text))
}

func EncodeBase64(b []byte) string {
        return base64.StdEncoding.EncodeToString(b)
}

func DecodeBase64(s string) []byte {
        data, err := base64.StdEncoding.DecodeString(s)
        if err != nil {
                panic(err)
        }
        return data
}