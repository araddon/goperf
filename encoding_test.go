package goperf

import (
	"bytes"
	"code.google.com/p/goprotobuf/proto"
	"encoding/gob"
	"encoding/json"
	"github.com/araddon/goperf/pb"
	"github.com/araddon/goperf/th"
	"github.com/araddon/thrift4go/lib/go/thrift"
	"github.com/ugorji/go-msgpack"
	"labix.org/v2/mgo/bson"
	"net/url"
	"testing"

	"log"
)

/*
go test -bench="coding"
go test -bench="Bson"

*/

type User struct {
	Lang                         string
	Verified                     bool
	Followers_count              int
	Location                     string
	Screen_name                  string
	Following                    bool
	Friends_count                int
	Profile_background_color     string
	Favourites_count             int
	Description                  string
	Notifications                string
	Profile_text_color           string
	Url                          string
	Time_zone                    string
	Statuses_count               int
	Profile_link_color           string
	Geo_enabled                  bool
	Profile_background_image_url string
	Protected                    bool
	Contributors_enabled         bool
	Profile_sidebar_fill_color   string
	Name                         string
	Profile_background_tile      string
	Created_at                   string
	Profile_image_url            string
	Id                           int64
	Utc_offset                   int
	Profile_sidebar_border_color string
}

type Tweet struct {
	Text                    string
	RawBytes                []byte
	Truncated               bool
	Geo                     string
	In_reply_to_screen_name string
	Favorited               bool
	Source                  string
	Contributors            string
	In_reply_to_status_id   string
	In_reply_to_user_id     int64
	Id                      int64
	Id_str                  string
	Created_at              string
	User                    *User
}

var (
	jsons     string = `{"in_reply_to_status_id_str":null,"geo":null,"retweet_count":21,"in_reply_to_status_id":null,"text":"RT @NayMunRai: \"\u0e1d\u0e37\u0e19\"\u0e17\u0e35\u0e48\u0e17\u0e23\u0e21\u0e32\u0e19.. \u0e04\u0e37\u0e2d\u0e01\u0e32\u0e23\u0e1d\u0e37\u0e19\u0e17\u0e31\u0e49\u0e07\u0e46\u0e17\u0e35\u0e48\u0e23\u0e39\u0e49\u0e27\u0e48\u0e32\u0e44\u0e21\u0e48\u0e21\u0e35\u0e17\u0e35\u0e48\"\u0e22\u0e37\u0e19\"","in_reply_to_user_id_str":null,"truncated":false,"entities":{"urls":[],"hashtags":[],"user_mentions":[{"indices":[3,13],"screen_name":"NayMunRai","name":"\u0e15\u0e4b\u0e07 \u0e2b\u0e21\u0e4a\u0e07 \u0e02\u0e37\u0e48\u0e2d (\u30c4)","id":330589922,"id_str":"330589922"}]},"place":null,"retweeted":false,"source":"web","in_reply_to_screen_name":null,"coordinates":null,"retweeted_status":{"in_reply_to_status_id_str":null,"geo":null,"retweet_count":21,"in_reply_to_status_id":null,"text":"\"\u0e1d\u0e37\u0e19\"\u0e17\u0e35\u0e48\u0e17\u0e23\u0e21\u0e32\u0e19.. \u0e04\u0e37\u0e2d\u0e01\u0e32\u0e23\u0e1d\u0e37\u0e19\u0e17\u0e31\u0e49\u0e07\u0e46\u0e17\u0e35\u0e48\u0e23\u0e39\u0e49\u0e27\u0e48\u0e32\u0e44\u0e21\u0e48\u0e21\u0e35\u0e17\u0e35\u0e48\"\u0e22\u0e37\u0e19\"","in_reply_to_user_id_str":null,"truncated":false,"entities":{"urls":[],"hashtags":[],"user_mentions":[]},"place":null,"retweeted":false,"source":"web","in_reply_to_screen_name":null,"coordinates":null,"in_reply_to_user_id":null,"user":{"contributors_enabled":false,"geo_enabled":true,"friends_count":306,"profile_sidebar_fill_color":"cf3863","url":"http:\/\/www.facebook.com\/Kaweekeemao","verified":false,"profile_sidebar_border_color":"25a5cc","show_all_inline_media":false,"listed_count":18,"follow_request_sent":null,"lang":"en","description":"\u0e04\u0e33\u0e04\u0e21\u0e41\u0e19\u0e27\u0e46 \u0e44\u0e21\u0e48\u0e0b\u0e49\u0e33\u0e43\u0e04\u0e23 \u0e42\u0e14\u0e19\u0e43\u0e08 \u0e1a\u0e32\u0e14\u0e08\u0e34\u0e15 \u0e15\u0e49\u0e2d\u0e07\u0e17\u0e35\u0e48\u0e17\u0e27\u0e34\u0e15 - @NayMunRai","location":"\u0e01\u0e27\u0e35\u0e02\u0e35\u0e49\u0e40\u0e21\u0e32","is_translator":false,"profile_use_background_image":true,"default_profile":false,"statuses_count":18305,"notifications":null,"time_zone":"Pacific Time (US & Canada)","profile_text_color":"21ded5","following":null,"profile_background_image_url":"http:\/\/a3.twimg.com\/profile_background_images\/392716142\/Copy_of_283569_144311465650792_130070737074865_280525_3018477_n.jpg","screen_name":"NayMunRai","profile_background_image_url_https":"https:\/\/si0.twimg.com\/profile_background_images\/392716142\/Copy_of_283569_144311465650792_130070737074865_280525_3018477_n.jpg","profile_link_color":"0ac24a","followers_count":3286,"protected":false,"profile_image_url":"http:\/\/a0.twimg.com\/profile_images\/1877092370\/317387_231913520207675_166962110036150_570150_539838096_n_normal.jpg","profile_image_url_https":"https:\/\/si0.twimg.com\/profile_images\/1877092370\/317387_231913520207675_166962110036150_570150_539838096_n_normal.jpg","name":"\u0e15\u0e4b\u0e07 \u0e2b\u0e21\u0e4a\u0e07 \u0e02\u0e37\u0e48\u0e2d (\u30c4)","default_profile_image":false,"profile_background_color":"1f9c9c","id":330589922,"id_str":"330589922","profile_background_tile":true,"utc_offset":-28800,"favourites_count":17,"created_at":"Wed Jul 06 21:09:29 +0000 2011"},"favorited":false,"id":176673544149798912,"id_str":"176673544149798912","contributors":null,"created_at":"Mon Mar 05 14:20:30 +0000 2012"},"in_reply_to_user_id":null,"user":{"contributors_enabled":false,"geo_enabled":true,"friends_count":1322,"profile_sidebar_fill_color":"ffffff","url":"http:\/\/nnanuns.tumblr.com\/","verified":false,"profile_sidebar_border_color":"fefffc","show_all_inline_media":true,"listed_count":1,"follow_request_sent":null,"lang":"en","description":"\u0e0a\u0e37\u0e48\u0e2d (\u0e16\u0e32\u0e21\u0e1c\u0e21\u0e2a\u0e34)' ,\u0e23\u0e31\u0e01\u0e44\u0e01\u0e48 \u0e40\u0e21\u0e19\u0e2b\u0e21\u0e32 \u0e1a\u0e49\u0e32\u0e1b\u0e25\u0e32 \u0e04\u0e23\u0e31\u0e48\u0e07\u0e27\u0e2d\u0e19 , \u0e2a\u0e48\u0e27\u0e19\u0e04\u0e19\u0e2d\u0e37\u0e48\u0e19\u0e41\u0e17\u0e1a\u0e44\u0e21\u0e48\u0e23\u0e39\u0e49\u0e08\u0e31\u0e01 (?) \u0e41\u0e15\u0e48\u0e15\u0e2d\u0e19\u0e19\u0e35\u0e49\u0e1e\u0e22\u0e32\u0e22\u0e32\u0e21(\u0e1e\u0e22\u0e32\u0e22\u0e32\u0e21)\u0e28\u0e36\u0e01\u0e29\u0e32\u0e2d\u0e22\u0e39\u0e48. \u0e1f\u0e25\u0e27.\u0e44\u0e21\u0e48\u0e1f\u0e25\u0e27.\u0e01\u0e47\u0e15\u0e32\u0e21\u0e2a\u0e1a\u0e32\u0e22\u0e04\u0e23\u0e31\u0e1a\u0e1c\u0e21 8?","location":"\u0e17\u0e35\u0e48\u0e41\u0e2b\u0e48\u0e07\u0e19\u0e35\u0e49.","is_translator":false,"profile_use_background_image":true,"default_profile":false,"statuses_count":7563,"notifications":null,"time_zone":"Bangkok","profile_text_color":"fa0505","following":null,"profile_background_image_url":"http:\/\/a0.twimg.com\/profile_background_images\/448225591\/8784.png","screen_name":"Nnanuns","profile_background_image_url_https":"https:\/\/si0.twimg.com\/profile_background_images\/448225591\/8784.png","profile_link_color":"ffcd06","followers_count":323,"protected":false,"profile_image_url":"http:\/\/a0.twimg.com\/profile_images\/1894706962\/174432_776639683_5641056_n_normal.png","profile_image_url_https":"https:\/\/si0.twimg.com\/profile_images\/1894706962\/174432_776639683_5641056_n_normal.png","name":"0813.","default_profile_image":false,"profile_background_color":"ffcd06","id":224607980,"id_str":"224607980","profile_background_tile":true,"utc_offset":25200,"favourites_count":4927,"created_at":"Thu Dec 09 12:16:43 +0000 2010"},"favorited":false,"id":180388160864395264,"id_str":"180388160864395264","contributors":null,"created_at":"Thu Mar 15 20:21:03 +0000 2012"}`
	bsonTweet []byte
	jsonTweet []byte
	protoTw   []byte
	thriftTw  []byte
	msgpackTw []byte
	gobTw     bytes.Buffer
	tw        Tweet
	twl       map[string]interface{}
	pbtw      pb.Tweet
	thtw      th.ThriftTweet
	qs        string = "name1=value1&name2=value2&name3=value3&name4=value4&name5=value5&name6=value6&name7=value7&name8=value8&name9=value9&name10=value10"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	thtw = *th.NewThriftTweet()
	// create structs, maps by using json data to do initial population
	json.Unmarshal([]byte(jsons), &tw)
	json.Unmarshal([]byte(jsons), &pbtw)
	json.Unmarshal([]byte(jsons), &twl)
	json.Unmarshal([]byte(jsons), &thtw)
	//log.Println("pbtw\n", pbtw)
	// create by values per serialization type
	bsonTweet, _ = bson.Marshal(&tw)
	jsonTweet = []byte(jsons)
	protoTw, _ = proto.Marshal(&pbtw)
	//ptw2, err := proto.Marshal(&pbtw)
	//log.Println(err)
	//log.Println(ptw2)
	msgpackTw, _ = msgpack.Marshal(tw)
	enc := gob.NewEncoder(&gobTw)
	err := enc.Encode(tw)
	if err != nil {
		panic("error")
	}
	//log.Println(string(gobTw.Bytes()))

	buf := thrift.NewTMemoryBuffer()
	thbp := thrift.NewTBinaryProtocol(buf, false, true)
	thtw.Write(thbp)
	thriftTw = buf.Bytes()
	tw2 := th.NewThriftTweet()
	tw2.Read(thbp)

}

func BenchmarkEncodingJsonTweet(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(&twl)
	}
}

//1.02 BenchmarkEncodingJsonTweet	    5000	    539056 ns/op = 1855/sec
//1.10 BenchmarkEncodingJsonTweet	   10000	    137992 ns/op

func BenchmarkEncodingJsonTweetStruct(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(&tw)
		//log.Println(data)
	}
}

//1.02 BenchmarkEncodingJsonTweetStruct	   50000	     69187 ns/op = 14,454/sec
//1.10 BenchmarkEncodingJsonTweetStruct	  100000	     26531 ns/op = 37,692/sec

func BenchmarkDecodingJsonTweet(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var tw map[string]interface{}
		json.Unmarshal(jsonTweet, &tw)
	}
}

// 1.02 BenchmarkEncodingJsonTweet	    5000	    510799 ns/op  = 1,946/sec
// 1.10 BenchmarkDecodingJsonTweet	   10000	    199803 ns/op

func BenchmarkDecodingJsonTweetStruct(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tw := Tweet{}
		json.Unmarshal([]byte(jsons), &tw)
	}
}

// 1.02 BenchmarkDecodingJsonTweetStruct	    2000	   1137723 ns/op  = 873/sec
// 1.10 BenchmarkDecodingJsonTweetStruct	   10000	    194380 ns/op

func BenchmarkEncodingGobTweetStruct(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var bb bytes.Buffer
		enc := gob.NewEncoder(&bb)
		_ = enc.Encode(tw)
		_ = bb.Bytes()
	}
}

//1.02 BenchmarkEncodingGobTweetStruct	   20000	     84273 ns/op = 11,866/sec
//1.10 BenchmarkEncodingGobTweetStruct	   50000	     32141 ns/op = 31,113/sec

func BenchmarkDecodingGobTweet(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tw := Tweet{}
		dec := gob.NewDecoder(&gobTw)
		err := dec.Decode(&tw)
		if err != nil {
			b.Fatalf("Error unmarshaling json: %v", err)
		}
	}
}

// BenchmarkDecodingGobTweet	   500000	     3879 ns/op  = 257,798/sec

func BenchmarkEncodingBsonTweetStruct(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = bson.Marshal(&tw)
	}
}

// 1.02 BenchmarkEncodingBsonTweetStruct	   50000	     103165 ns/op =  9,693/sec
// 1.10 BenchmarkEncodingBsonTweetStruct	  100000	      23573 ns/op = 42,421/sec

func BenchmarkDecodingBsonTweet(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var tw map[string]interface{}
		bson.Unmarshal(bsonTweet, &tw)
	}
}

// 1.02 BenchmarkDecodingBsonTweet	   10000	     120168 ns/op  = 8264/sec
// 1.10 BenchmarkDecodingBsonTweet	   50000	     43841 ns/op

func BenchmarkDecodingBsonTweetStruct(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tw := Tweet{}
		bson.Unmarshal(bsonTweet, &tw)
	}
}

// 1.02 BenchmarkDecodingBsonTweetStruct	   20000	     81859 ns/op  = 12,216/sec
// 1.10 BenchmarkDecodingBsonTweetStruct	  100000	     23506 ns/op  = 42,542/sec

func BenchmarkDecodingQueryNV(b *testing.B) {
	b.StartTimer()
	// ParseQuery(query string) (m Values, err error)
	for i := 0; i < b.N; i++ {
		_, _ = url.ParseQuery(qs)
	}
}

// 1.02 BenchmarkDecodingQueryNV	  100000	     16101 ns/op = 62108/sec
// 1.10 BenchmarkDecodingQueryNV	  500000	      6084 ns/op =

func BenchmarkEncodingPBTweetStruct(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = proto.Marshal(&pbtw)
	}
}

// 1.02 BenchmarkEncodingPBTweetStruct	  200000	      13181 ns/op =  75,867/sec
// 1.10 BenchmarkEncodingPBTweetStruct	  500000	      4730 ns/op  = 164,366/sec

func BenchmarkDecodingPBTweetStruct(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tw := pb.Tweet{}
		proto.Unmarshal(protoTw, &tw)
	}
}

// 1.02 BenchmarkDecodingPBTweetStruct	  100000	      18960 ns/op =  52,743/sec
// 1.10 BenchmarkDecodingPBTweetStruct	  500000	       6969 ns/op = 143,493/sec

func BenchmarkEncodingMPTweetStruct(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = msgpack.Marshal(&tw)
	}
}

// 1.02 BenchmarkEncodingMPTweetStruct	   10000	     120501 ns/op = 8,330/sec
// 1.10 BenchmarkEncodingMPTweetStruct	  100000	      24138 ns/op

func BenchmarkDecodingMPTweetStruct(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mtw := Tweet{}
		msgpack.Unmarshal(msgpackTw, &mtw, nil)
	}
}

// 1.02 BenchmarkDecodingMPTweetStruct	   10000	     102020 ns/op = 9,802/sec
// 1.10 BenchmarkDecodingMPTweetStruct	   50000	     33395 ns/op

func BenchmarkEncodingThriftTweetStruct(b *testing.B) {
	b.StartTimer()
	// presumably there is a faster/better way to do this?
	tmem := thrift.NewTMemoryBuffer()
	thbp := thrift.NewTBinaryProtocol(tmem, false, true)
	for i := 0; i < b.N; i++ {
		thtw.Write(thbp)
		_ = tmem.Bytes()
		tmem.Reset()
	}
}

// 1.02 BenchmarkEncodingThriftTweetStruct	  100000	      26294 ns/op = 38,031/sec
// 1.10 BenchmarkEncodingThriftTweetStruct	  100000	      20001 ns/op = 49,998/sec

func BenchmarkDecodingThriftTweetStruct(b *testing.B) {

	b.StartTimer()
	tmem := thrift.NewTMemoryBuffer()
	thbp := thrift.NewTBinaryProtocol(tmem, false, true)
	for i := 0; i < b.N; i++ {
		tmem.Write(thriftTw)
		tw := th.NewThriftTweet()
		tw.Read(thbp)
	}
}

// 1.02 BenchmarkDecodingThriftTweetStruct	   20000	     74837 ns/op = 13,362/sec
// 1.10 BenchmarkDecodingThriftTweetStruct	   50000	     45650 ns/op = 21,906
