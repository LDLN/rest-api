# Welcome to LDLN ReST API

## Intro

   An API for creating and listing LDLN SyncableObjects, built with the revel framework, go, and mongodb

## Configure Application

   There are some extra requirements for the LDLN API in addition to the standard revel parameters. For this API to function, the following values need to be populated in conf/app.conf:

#### revmgo settings: see https://github.com/jgraham909/revmgo

   revmgo.dial = database_ip/database_name 

   revmgo.method = new|clone|copy

#### LDLN settings for mapping models to mongodb collections and rawdek

   landline.db.collection.SyncableObject = SyncableObjects

   landline.db.collection.Schema = Schemas

   landline.rawdek = PasswordForNotVerySensitiveData!

## Start the web server:

   revel run github.com/msgrasser/api

   Run with <tt>--help</tt> for options.

## Go to http://localhost:9100/ and you'll see a list of any SyncableObjects

## Testing the creation of a SyncableObject

  1. Make sure there is a Schema (aka "type") created for the type of object you want to create
  2. POST to the create endpoint:

  `curl -X POST http://localhost:9100/types/GPS/create?rawdata=TESTParam,123,7383`

<hr />

Stuck on something? Check [the wiki](https://github.com/LDLN/core/wiki)!
