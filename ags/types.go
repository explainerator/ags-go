package ags

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Codes struct {
	Upc string `bson:"upc,omitempty"`
	Sku string `bson:"sku,omitempty"`
}

type ExternalId struct {
	Poynt string `bson:"Poynt,omitempty"`
	Extra bson.M `bson:",inline"`
}

type Item struct {
	Id         string     `bson:"id,omitempty"`
	Name       string     `bson:"name,omitempty"`
	Codes      Codes      `bson:"codes,omitempty"`
	Price      int32      `bson:"price,omitempty"`
	ExternalId ExternalId `bson:"externalId,omitempty"`
	Extra      bson.M     `bson:",inline"`
	DbName     string
}

// {
// 	"Extra": {
// 	  "__v": 10,
// 	  "_id": "60f9026eae6871278457d9e4",
// 	  "brand": "Grav",
// 	  "categoryId": "qZtTwC2Hcc6LgzHLketRV2",
// 	  "codes": {
// 		"upc": "665570227150"
// 	  },
// 	  "cost": 1261,
// 	  "description": "A little bit of science can go a long way. The GRAV\u00ae Beaker Spoon echoes the iconic silhouette of our beaker water pipes. But this little beaker doesn't need any water or moving parts to stand out. In fact it stands up on its own, making it easy to put down between hits. The shape may be whimsical but this pipe functions like a dream, with an ergonomic carb and a flush mouthpiece designed for comfort. ",
// 	  "discontinued": false,
// 	  "distributorUrl": "",
// 	  "externalId": {
// 		"Square": "JW5EM2SSRXZZUFXSHXXOG53L"
// 	  },
// 	  "flags": {
// 		"discontinued": false,
// 		"hidden": false,
// 		"special": false
// 	  },
// 	  "group": {
// 		"enabled": true,
// 		"groupId": "7TUMnaviKdDCpE4BaBF25q",
// 		"variantName": "Green"
// 	  },
// 	  "hidden": false,
// 	  "homeUrls": [],
// 	  "id": "eGZTW6o4tvqxYVaYRtqubL",
// 	  "imageId": "eGZTW6o4tvqxYVaYRtqubL",
// 	  "inventory": {
// 		"bridgewater": {
// 		  "count": 1,
// 		  "history": [
// 			{
// 			  "action": "reset",
// 			  "count": 1,
// 			  "resetAt": "2023-05-06T05:43:59.000Z",
// 			  "updatedAt": "2023-05-06T05:43:59.000Z"
// 			},
// 			{
// 			  "action": "reset",
// 			  "count": 2,
// 			  "resetAt": "2023-05-16T06:55:30.005Z",
// 			  "updatedAt": "2023-05-16T06:55:30.005Z"
// 			},
// 			{
// 			  "action": "reset",
// 			  "count": 1,
// 			  "resetAt": "2023-05-17T20:32:08.709-03:00",
// 			  "updatedAt": "2023-05-17T20:32:08.709-03:00"
// 			}
// 		  ],
// 		  "resetAt": "2023-05-17T20:32:08.709-03:00",
// 		  "updatedAt": "2023-05-17T20:32:08.709-03:00"
// 		},
// 		"liverpool": {
// 		  "count": 0,
// 		  "history": [
// 			{
// 			  "action": "reset",
// 			  "count": 0,
// 			  "resetAt": "1969-12-31T20:00:00-04:00",
// 			  "updatedAt": "2023-05-06T21:52:30.837-03:00"
// 			},
// 			{
// 			  "action": "zeroed",
// 			  "count": 0,
// 			  "resetAt": "1969-12-31T20:00:00-04:00",
// 			  "updatedAt": "2023-05-06T21:52:30.837-03:00"
// 			}
// 		  ],
// 		  "resetAt": "1969-12-31T20:00:00-04:00",
// 		  "updatedAt": "2023-05-17T16:38:31.661-03:00"
// 		},
// 		"warehouse": {
// 		  "count": 0,
// 		  "history": [
// 			{
// 			  "action": "reset",
// 			  "count": 0,
// 			  "resetAt": "1969-12-31T20:00:00-04:00",
// 			  "updatedAt": "1969-12-31T20:00:00-04:00"
// 			},
// 			{
// 			  "action": "zeroed",
// 			  "count": 0,
// 			  "resetAt": "1969-12-31T20:00:00-04:00",
// 			  "updatedAt": "1969-12-31T20:00:00-04:00"
// 			}
// 		  ],
// 		  "resetAt": "1969-12-31T20:00:00-04:00",
// 		  "updatedAt": "2023-05-17T16:38:31.56-03:00"
// 		}
// 	  },
// 	  "name": "Grav > 4\" Beaker Spoon Green",
// 	  "notes": "",
// 	  "price": 2495,
// 	  "stock": [
// 		1,
// 		0,
// 		0
// 	  ],
// 	  "supplierUrls": [
// 		"https://bobhq.com/grav-4-beaker-spoon/"
// 	  ],
// 	  "updatedAt": "2021-07-10T01:50:07.08-03:00"
// 	}
//   }
