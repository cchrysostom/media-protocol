package messages

import (
	"encoding/json"
	"fmt"
	//"reflect"
	"testing"
)

func TestDecodeOIP041(t *testing.T) {
	// ToDo: a better test/errors
	//oip041a, err := DecodeOIP041(oip041_music_example)
	//if err != nil {
	//	t.Error("error 1")
	//}
	//if !reflect.DeepEqual(oip041a, oip041_music_example_obj) {
	//	t.Error("not equal 1")
	//}
	//
	//oip041e, err := DecodeOIP041(oip041_edit_example)
	//if err != nil {
	//	t.Error("error 2")
	//}
	//if !reflect.DeepEqual(oip041e, oip041_edit_example_obj) {
	//	t.Error("not equal 2")
	//}

	fmt.Println("Testing OIP041 oip_mp_example")
	oip041t, err := DecodeOIP041(oip_mp_example)
	fmt.Println("DONE Testing OIP041 oip_mp_example")
	fmt.Println(oip041t)
	fmt.Println(err)

	err = StoreOIP041Artifact(oip041t, "txid", 6, nil)
	fmt.Println(err)
}

//var oip_mp_example = `{
//  "oip-041": {
//    "artifact": {
//      "type": "Research-Tomogram",
//      "info": {
//        "title": "Test Tomogram",
//        "description": "Description for the Test Tomogram",
//        "year": 2018
//      },
//      "storage": {
//        "network": "ipfs",
//        "location": "QmRWCQQXVHuu8jKALsGaXayecwxN5q4i3sJo738sAifAvq",
//        "files": [
//          {
//            "fname": "tomogram-Converted_1.mp4",
//            "fsize": 8371978
//          }
//        ]
//      },
//      "timestamp": 1521499921,
//      "publisher": "FLZXRaHzVPxJJfaoM32CWT4GZHuj2rx63k"
//    },
//    "signature": "H0IJSKIqSlxk81cO3w7fExdNc9X0WOMPHHm3gDtEGkyPYaIZJAjq7gkAwsgAOUsi3oWT+vgfhwyaDy6ycPTEBpw="
//  }
//}`

//var oip_mp_example = `{"oip-041":{"artifact":{"type":"Video-Basic","info":{"extraInfo":{"artist":"Doggo","genre":"Pets & Animals","tags":["Doggo"]},"title":"Doggo Doggo","year":2017,"description":"Doggo"},"storage":{"network":"IPFS","files":[]},"payment":{"fiat":"USD","scale":"1000:1","disPer":0.3,"sugTip":[],"tokens":{}},"timestamp":1503538609,"publisher":"F7XXR9G2ePrmqur3gJWWYsvMUZqioarPAr"},"signature":"IB6vOzVimzipryyxettJIGrZcPAFC10o0OIWfMGjszG+AO5JLkoLCixn1v6rvVoQLEifx9AYYAEW3emdM7KrS7g="}}`

var oip_mp_example = /*tokenly_example*/ `{"oip-041":{"artifact":{"publisher":"FQV7FCrnqSZH98zBQLaSWa6NVyBFzh7jtx","timestamp":1483483636,"type":"music","info":{"title":"Sad Powers","description":"This is music generated by an artificial intelligence algorithm.","year":2017,"extraInfo":{"artist":"Mr. Chips","composers":"Artificial Intelligence","company":"The Internet","copyright":"No Copyright","usageProhibitions":"No prohibitions.","usageRights":"Can be used anywhere.","genre":"Instrumental","tags":"AI, Instrumental, MIDI"}},"storage":{"network":"IPFS","location":"QmZXViTYh7QqYGvSRQqu4G6nuSs3KxkuPPTsDjC2MCB4dD","files":[{"dname":"Sad Powers","fname":"8fc38057537a015f48d182ad85a03eb3043b198a.mp3","fsize":1969262,"minPlay":100,"sugPlay":200,"type":"music","tokenlyID":"ac976fa8-bca0-4569-976b-12d902a037bc"},{"dname":"Cover Art","fname":"204f7089b1ed78e2253aef34ae18b3a2e5791439.jpg","fsize":152480,"type":"coverArt"}]},"payment":{"fiat":"USD","scale":"100:1","sugTip":200,"tokens":{"SOUP":1,"DEVON":1},"addresses":[{"token":"BTC","address":"1JztLWos5K7LsqW5E78EASgiVBaCe6f7cD"}]}},"signature":"IDalLHPvzxmWOiWs8eUGtfrn6gFwZe/Vxrs0Joh85DA9Bnw9gpKVD5h+WeWPZna+9XeaYrI5ggzRI0O+kcU6foU="}}`

var oip041_edit_example_obj Oip041 = Oip041{
	Edit: Oip041Edit{
		Timestamp: 1234567890,
		TxID:      "$artifactID",
		Patch:     json.RawMessage(oip041_edit_example_patch_rawmsg),
	},
	Signature: "$txid-$MD5HashOfPatch-$timestamp",
}

var oip041_edit_example_patch_rawmsg = []byte(`{
                "add":[
                    {
                        "path":"/payment/tokens/mtcproducer",
                        "value":""
                    }
                ],
                "replace":[
                    {
                        "path":"/storage/files/3/fname",
                        "value":"birthdayepFirst.jpg"
                    },
                    {
                        "path":"/storage/files/3/dname",
                        "value":"Cover Art 2"
                    },
                    {
                        "path":"/info/title",
                        "value":"Happy Birthday"
                    },
                    {
                        "path":"/timestamp",
                        "value":1481420001
                    }
                ],
                "remove":[
                    {
                        "path":"/payment/tokens/mtmproducer"
                    },
                    {
                        "path":"/storage/files/0/sugBuy"
                    }
                ]
            }`)

var oip041_edit_example = `{
    "oip-041":{
        "editArtifact":{
            "txid":"$artifactID",
            "timestamp":1234567890,
            "patch":{
                "add":[
                    {
                        "path":"/payment/tokens/mtcproducer",
                        "value":""
                    }
                ],
                "replace":[
                    {
                        "path":"/storage/files/3/fname",
                        "value":"birthdayepFirst.jpg"
                    },
                    {
                        "path":"/storage/files/3/dname",
                        "value":"Cover Art 2"
                    },
                    {
                        "path":"/info/title",
                        "value":"Happy Birthday"
                    },
                    {
                        "path":"/timestamp",
                        "value":1481420001
                    }
                ],
                "remove":[
                    {
                        "path":"/payment/tokens/mtmproducer"
                    },
                    {
                        "path":"/storage/files/0/sugBuy"
                    }
                ]
            }
        },
    	"signature":"$txid-$MD5HashOfPatch-$timestamp"
    }
}`

var oip041_music_example = `{
  "oip-041": {
    "artifact": {
      "publisher": "F97Tp8LYnw94CpXmAhqACXWTT36jyvLCWx",
      "timestamp": 1470269387,
      "type": "music",
      "storage":{
		"files": [
		  {
		    "dname": "Skipping Stones",
		    "fname": "1 - Skipping Stones.mp3",
		    "fsize": 6515667,
		    "type": "album track",
		    "duration": 1533.603293,
		    "sugPlay": 100,
		    "minPlay": null,
		    "sugBuy": 750,
		    "minBuy": 500,
		    "promo": 10,
		    "retail": 15,
		    "ptpFT": 10,
		    "ptpDT": 20,
		    "ptpDA": 50
		  },
		  {
		    "dname": "Lessons",
		    "fname": "2 - Lessons with intro.mp3",
		    "fsize": 6515667,
		    "type": "album track",
		    "duration": 1231.155243,
		    "disallowPlay": 1,
		    "sugBuy": 750,
		    "minBuy": 500,
		    "promo": 10,
		    "retail": 15,
		    "ptpFT": 10,
		    "ptpDT": 20,
		    "ptpDA": 50
		  },
		  {
		    "dname": "Born to Roam",
		    "fname": "3 - Born to Roam.mp3",
		    "fsize": 6515667,
		    "type": "album track",
		    "duration": 2374.550714,
		    "sugPlay": 100,
		    "minPlay": 50,
		    "disallowBuy": 1,
		    "promo": 10,
		    "retail": 15,
		    "ptpFT": 10,
		    "ptpDT": 20,
		    "ptpDA": 50
		  },
		  {
		    "dname": "Cover Art",
		    "fname": "birthdayepFINAL.jpg",
		    "type": "coverArt",
		    "disallowBuy": 1
		  }
	    ],
	    "network": "IPFS",
        "location": "QmPukCZKeJD4KZFtstpvrguLaq94rsWfBxLU1QoZxvgRxA"
      },
      "info": {
        "title": "Happy Birthday EP",
        "description": "this is the second organically grown, gluten free album released by Adam B. Levine - contact adam@tokenly.com with questions or comments or discuss collaborations.",
        "year": 2016,
        "extra-info": {
          "artist": "Adam B. Levine",
          "company": "",
          "composers": [
            "Adam B. Levine"
          ],
          "copyright": "",
          "tokenly_ID": "",
          "usageProhibitions": "",
          "usageRights": "",
          "tags": []
        }
      },
      "payment": {
        "fiat": "USD",
        "scale": "1000:1",
        "sug_tip": [
          5,
          50,
          100
        ],
        "tokens": {
          "MTMCOLLECTOR": "",
          "MTMPRODUCER": "",
          "HAPPYBDAYEP": "",
          "EARLY": "",
          "LTBCOIN": "",
          "BTC": "1GMMg2J5iUKnDf5PbRr9TcKV3R6KfUiB55"
        }
      }
    },
    "signature": "H27r7UxUb8BozjEvV0v++nCyRI7S6yyroeKCJQpgU5NO3CP6FpXWs5kCxy8vhmMhbtpj/FMj+8s3+updw7g+bmE="
  }
}`

var ei, _ = json.Marshal(Oip041MusicExtraInfo{
	Artist:  "Adam B. Levine",
	Company: "",
	Composers: []string{
		"Adam B. Levine",
	},
	Copyright:         "",
	UsageProhibitions: "",
	UsageRights:       "",
	Tags:              []string{},
})

var oip041_music_example_obj Oip041 = Oip041{
	Artifact: Oip041Artifact{
		Publisher: "F97Tp8LYnw94CpXmAhqACXWTT36jyvLCWx",
		Timestamp: 1470269387,
		Type:      "music",
		Info: Oip041Info{
			Title:           "Happy Birthday EP",
			Description:     "this is the second organically grown, gluten free album released by Adam B. Levine - contact adam@tokenly.com with questions or comments or discuss collaborations.",
			Year:            2016,
			ExtraInfo:       ei,
			ExtraInfoString: "",
		},
		Storage: Oip041Storage{
			Network:  "IPFS",
			Location: "QmPukCZKeJD4KZFtstpvrguLaq94rsWfBxLU1QoZxvgRxA",
			Files: []Oip041Files{
				{
					DisallowBuy:  false,
					Dname:        "Skipping Stones",
					Duration:     1533.603293,
					Fname:        "1 - Skipping Stones.mp3",
					Fsize:        6515667,
					MinPlay:      0,
					SugPlay:      100,
					Promo:        10,
					Retail:       15,
					PtpFT:        10,
					PtpDT:        20,
					PtpDA:        50,
					Type:         "album track",
					TokenlyID:    "",
					DisallowPlay: false,
					MinBuy:       500,
					SugBuy:       750,
					//Storage:      Oip041Storage{},
				},
				{
					DisallowBuy:  false,
					Dname:        "Lessons",
					Duration:     1231.155243,
					Fname:        "2 - Lessons with intro.mp3",
					Fsize:        6515667,
					MinPlay:      0,
					SugPlay:      0,
					Promo:        10,
					Retail:       15,
					PtpFT:        10,
					PtpDT:        20,
					PtpDA:        50,
					Type:         "album track",
					TokenlyID:    "",
					DisallowPlay: true,
					MinBuy:       500,
					SugBuy:       750,
					//Storage:      Oip041Storage{},
				},
				{
					DisallowBuy:  true,
					Dname:        "Born to Roam",
					Duration:     2374.550714,
					Fname:        "3 - Born to Roam.mp3",
					Fsize:        6515667,
					MinPlay:      50,
					SugPlay:      100,
					Promo:        10,
					Retail:       15,
					PtpFT:        10,
					PtpDT:        20,
					PtpDA:        50,
					Type:         "album track",
					TokenlyID:    "",
					DisallowPlay: false,
					MinBuy:       0,
					SugBuy:       0,
					//Storage:      Oip041Storage{},
				},
				{
					DisallowBuy:  true,
					Dname:        "Cover Art",
					Duration:     0,
					Fname:        "birthdayepFINAL.jpg",
					Fsize:        0,
					MinPlay:      0,
					SugPlay:      0,
					Promo:        0,
					Retail:       0,
					PtpFT:        0,
					PtpDT:        0,
					PtpDA:        0,
					Type:         "coverArt",
					TokenlyID:    "",
					DisallowPlay: false,
					MinBuy:       0,
					SugBuy:       0,
					//Storage:      Oip041Storage{},
				},
			},
		},
		Payment: Oip041Payment{
			Fiat:  "USD",
			Scale: "1000:1",
			SugTip: []int{
				5,
				50,
				100,
			},
			Tokens: map[string]int{
				"MTMCOLLECTOR": 1,
				"MTMPRODUCER":  1,
				"HAPPYBDAYEP":  1,
				"EARLY":        1,
				"LTBCOIN":      1},
			Addresses: []Oip041Address{
				{
					Token:   "BTC",
					Address: "1GMMg2J5iUKnDf5PbRr9TcKV3R6KfUiB55",
				},
			},
		},
	},
	Signature: "H27r7UxUb8BozjEvV0v++nCyRI7S6yyroeKCJQpgU5NO3CP6FpXWs5kCxy8vhmMhbtpj/FMj+8s3+updw7g+bmE=",
}
