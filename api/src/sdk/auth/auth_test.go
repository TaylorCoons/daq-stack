package auth

import (
	"encoding/base64"
	"fmt"
	"testing"
)

// TODO: Move these into config
var correctUser string = "admin"
var correctPass string = "pass"

var incorrectUser string = "incorrect"
var incorrectPass string = "incorrect"

func TestBasicAuth(t *testing.T) {
	value := EncodeBasicAuth(correctUser, correctPass)
	if basicAuth(value) != true {
		t.Error("admin auth did not authenticate actual value")
	}

	value = EncodeBasicAuth(incorrectUser, incorrectPass)
	if basicAuth(value) == true {
		t.Error("admin auth authenticated incorrect value")
	}
}

func TestGenerateToken(t *testing.T) {
	const testSize = 32
	res, err := GenerateToken(testSize)
	if err != nil {
		t.Error("generate token returned an error")
	}
	decoded, err := base64.URLEncoding.DecodeString(res.Token)
	if err != nil {
		t.Error("generate token returned non url encoded string")
	}
	if len(decoded) != testSize {
		t.Error("generate token returned wrong sized token")
	}
}

// func TestCreateToken(t *testing.T) {
// 	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
// 	defer mt.Close()

// 	mt.Run("success", func(mt *mtest.T) {

// 		encoding := EncodeBasicAuth(correctUser, correctPass)

// 		token, err := CreateToken(client, encoding)
// 		if err != nil {
// 			t.Error("create token returned error")
// 		}

// 		token := models.Token{
// 			Token:     "Test",
// 			CreatedAt: time.Now().UTC(),
// 			ExpiresOn: time.Now().UTC(),
// 		}

// 		coll := mt.CreateCollection(mtest.Collection{
// 			Name: "test",
// 		}, true)
// 		coll.InsertOne(context.Background(), bson.D{{"X", 1}})

// 		res := coll.FindOne(context.Background(), bson.D{{"X", 1}})
// 		type temp struct {
// 			X int32 `bson:"x"`
// 		}
// 		v := temp{}
// 		res.Decode(&v)
// 		fmt.Println(v)
// 		t.Error("check")
// 	})

// 	encoding := EncodeBasicAuth(incorrectUser, incorrectPass)
// 	_, err := CreateToken(encoding)
// 	if _, ok := err.(*NotAuthorized); !ok {
// 	    t.Error("create token did not throw an authorization error")
// 	}
// 	encoding = EncodeBasicAuth(correctUser, correctPass)
// 	token, err := CreateToken(encoding)
// 	if err != nil {
// 		t.Error("create token returned unexpected error")
// 	}
// 	if len(token) == 0 {
// 		t.Error("create token returned empty token")
// 	}
// 	TODO: Test token is inserted into DB

// }

// func TestSupportedAuthTypes(t *testing.T) {
// 	ret := SupportedAuthTypes()
// 	for i := range supportedAuth {
// 		if supportedAuth[i] != ret[i] {
// 			t.Error("supported auth types not returned correctly")
// 		}
// 	}
// }

// func TestRevokeToken(t *testing.T) {
// 	encoding := EncodeBasicAuth(incorrectUser, incorrectPass)
// 	err := RevokeToken(encoding)
// 	if _, ok := err.(*NotAuthorized); !ok {
// 		t.Error("revoke token did not throw an authorization error")
// 	}
// 	encoding = EncodeBasicAuth(correctUser, correctPass)
// 	err = RevokeToken(encoding)
// 	if err != nil {
// 		t.Error("revoke token returned unexpected error")
// 	}
// 	// TODO: Test token is removed from DB
// }

func EncodeBasicAuth(user string, pass string) string {
	return base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user, pass)))
}
