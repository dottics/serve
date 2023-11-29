package auth

import (
	"testing"
)

//func TestGenerateJWT(t *testing.T) {
//	s, err := generateJWT()
//	if err != nil {
//		t.Errorf("expected no error got: %v", err)
//	}
//	log.Print("s", s)
//	t.Error("force fail")
//}

//func TestVerifyJWT2(t *testing.T) {
//	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJpc3MiOiJ0ZXN0IiwiYXVkIjoic2luZ2xlIn0.QAWg1vGvnqRuCFTMcPkjZljXHh8U3L_qUjszOtQbeaA"
//	verifyJWT2(tokenString)
//}

func TestParseJWT(t *testing.T) {
	// we should just be able to parse it
	tokenString, _ := createJWT()
	token, err := ParseJWT(tokenString)
	if token == nil {
		t.Errorf("expected token got nil")
	}
	_, err = token.Claims.GetAudience()
	if err != nil {
		t.Errorf("expected error nil got %v", err)
	}
	//log.Print(auds)
	//t.Error("force")
}

func TestValidateToken(t *testing.T) {
	tokenString, _ := createJWT()
	//tokenString := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJnUl9OS1BqRVZDRFB4V3NYeDBEbTFmOE90eWNHLUtFZW5zSGhMUTVWM3VNIn0.eyJleHAiOjE2OTk1MzA5MjMsImlhdCI6MTY5OTUzMDYyMywianRpIjoiNjY4NDVmNTEtNzY0OC00N2U2LWI1MDMtMmQ5NTAyNjZiNDY1IiwiaXNzIjoiaHR0cHM6Ly9pZHAtbm9uLXByb2QuaW50LmNhcGluZXQvYXV0aC9yZWFsbXMvREVWIiwiYXVkIjoiUGxhdGZvcm0tREVWIiwic3ViIjoiNWQxNzEwYjEtNmJiNC00NGIzLWJkNGItZjllZGQ1MGIxYzEwIiwidHlwIjoiQmVhcmVyIiwiYXpwIjoiUGxhdGZvcm0tREVWIiwic2Vzc2lvbl9zdGF0ZSI6IjA3NDk3NTE2LTIwNGItNDM2YS1hYzY0LTkyNDMyNTg4MzNlOSIsImFjciI6IjEiLCJhbGxvd2VkLW9yaWdpbnMiOlsiKiIsImh0dHA6Ly9sb2NhbGhvc3QiXSwic2NvcGUiOiJvcGVuaWQgZW1haWwgcHJvZmlsZSIsInNpZCI6IjA3NDk3NTE2LTIwNGItNDM2YS1hYzY0LTkyNDMyNTg4MzNlOSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwicm9sZXMiOlsiY2hhdC1pbnRlcm5hbCIsInByaXZhdGUiLCJicm9rZXIiXSwiY2xpZW50X2lwIjoiMTAuMjI0LjcuMzAiLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJzX2NoYXRfbXNfaW50X2RldiIsImVtYWlsIjoic19jaGF0X21zX2ludF9kZXZAY2FwaXRlY2JhbmsuY28uemEifQ.aImkWVT5K2e0WgCErromZii1beZnrYLlWdnMJpZttLA53W8DF5qXppwrpel4ZRR0lFrWSM3NW22r2PRBO4k-HDig4fvaNg6DqL3fW1-SZZ0_XxG9TKyanLsrH0Vp7LNW7BK4NvWHLAJPfxYvUY5wrVgaugHym1YFibzcrgN85VBd9Wo10ah5P5Y-81dz8UynH2IJZF8qXujtRQys1Bvl0U_8KIq7xUi-Q1WNGk104XgXZEfH14ILP5usctOYaD5Xc859WsdC5Y0-1GDRrgQvYcIchv91N919gHBr-nns6UNWReImg4CV0Flu-HWicWtFEsAwXf0pDNWbShWsYE0Geg"
	token, err := ParseJWT(tokenString)
	if token == nil {
		t.Errorf("expected token got nil")
	}

	err = ValidateToken(token)

	if err != nil {
		t.Errorf("token error: %v", err)
	}
}
