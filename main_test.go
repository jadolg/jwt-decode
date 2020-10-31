package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseCorrectly(t *testing.T) {
	testToken := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsIng1YyI6WyJNSUlDK1RDQ0FwK2dBd0lCQWdJQkFEQUtCZ2dxaGtqT1BRUURBakJHTVVRd1FnWURWUVFERXpzeVYwNVpPbFZMUzFJNlJFMUVVanBTU1U5Rk9reEhOa0U2UTFWWVZEcE5SbFZNT2tZelNFVTZOVkF5VlRwTFNqTkdPa05CTmxrNlNrbEVVVEFlRncweU1EQXhNRFl5TURVeE1UUmFGdzB5TVRBeE1qVXlNRFV4TVRSYU1FWXhSREJDQmdOVkJBTVRPMVZCVVRjNldGTk9VenBVUjFRek9rRTBXbFU2U1RWSFN6cFNOalJaT2xkRFNFTTZWMVpTU3pwTlNUTlNPa3RZVFRjNlNGZFRNenBDVmxwYU1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBcnh5Mm9uSDBTWHh4a1JCZG9wWDFWc1VuQVovOUpZR3JNSXlrelJuMTRsd1A1SkVmK1hNNUFORW1NLzBYOFJyNUlIN2VTWFV6K1lWaFVucVNKc3lPUi9xd3BTeFdLWUxxVnB1blFOWThIazdmRnlvN0l0bXgxajZ1dnRtVmFibFFPTEZJMUJNVnY2Y3IxVjV3RlZRYWc3SnhkRUFSZWtaR1M5eDlIcnM1NVdxb0lSK29GRGwxVVRjNlBFSjZVWGdwYmhXWHZoU0RPaXBPcUlYdHZkdHJoWFFpd204Y3EyczF0TEQzZzg2WmRYVFg3UDFFZkxFOG1jMEh4anBGNkdiNWxHZFhjdjU5cC9SMzEva0xlL09wRHNnVWJxMEFvd3Bsc1lLb0dlSmdVNDJaZG45SFZGUVFRcEtGTFBNK1pQN0R2ZmVGMWNIWFhGblI1TkpFU1Z1bFRRSURBUUFCbzRHeU1JR3ZNQTRHQTFVZER3RUIvd1FFQXdJSGdEQVBCZ05WSFNVRUNEQUdCZ1JWSFNVQU1FUUdBMVVkRGdROUJEdFZRVkUzT2xoVFRsTTZWRWRVTXpwQk5GcFZPa2sxUjBzNlVqWTBXVHBYUTBoRE9sZFdVa3M2VFVrelVqcExXRTAzT2toWFV6TTZRbFphV2pCR0JnTlZIU01FUHpBOWdEc3lWMDVaT2xWTFMxSTZSRTFFVWpwU1NVOUZPa3hITmtFNlExVllWRHBOUmxWTU9rWXpTRVU2TlZBeVZUcExTak5HT2tOQk5sazZTa2xFVVRBS0JnZ3Foa2pPUFFRREFnTklBREJGQWlFQXl5SEpJU1RZc1p2ZVZyNWE1YzZ4MjhrQ2U5M2w1QndQVGRUTk9SaFB0c0VDSURMR3pYdUxuekRqTCtzcWRkOU5FbkRuMnZ2UFBWVk1NLzhDQW1EaTVudnMiXX0.eyJhY2Nlc3MiOlt7InR5cGUiOiJyZXBvc2l0b3J5IiwibmFtZSI6InJhdGVsaW1pdHByZXZpZXcvdGVzdCIsImFjdGlvbnMiOlsicHVsbCJdLCJwYXJhbWV0ZXJzIjp7InB1bGxfbGltaXQiOiIxMDAiLCJwdWxsX2xpbWl0X2ludGVydmFsIjoiMjE2MDAifX1dLCJhdWQiOiJyZWdpc3RyeS5kb2NrZXIuaW8iLCJleHAiOjE2MDQxNTE3NjMsImlhdCI6MTYwNDE1MTQ2MywiaXNzIjoiYXV0aC5kb2NrZXIuaW8iLCJqdGkiOiJpeVBFdmFEZTV6aDd1MG5tNkI0SCIsIm5iZiI6MTYwNDE1MTE2Mywic3ViIjoiIn0.YkGk_mgH6VEsNhtruGXrIKfKZfEbJoHSy2JzfoLM329df9qmnitj8vRnHUYI2r_DF3AuA7ggA5lgPsJVcSknwokpg1LS0Kt8KN3mkZquZp9HcwJoGisEchMslVFz_noQabssEJwjbfIib5ImM0FwSg1cXnGthoqRavgoUqHNKUlQ91_Mkbk7WWgv8cD8wKKnAzq25UpYcsSLNK5NX5QKHaupvYV6ow7dqvrGhmEN9uQJnxVcICuXu6pFhVJeCXHXuvEcMOEk_hFauW0Xxvdz3YNwM8u8ze5WbAEq5QNSQtLNTzRFvcigTZKYfgBVANz0AWc-Elbw8cTowzzUBBG7ew"
	expectedData := `{
  "access": [
    {
      "actions": [
        "pull"
      ],
      "name": "ratelimitpreview/test",
      "parameters": {
        "pull_limit": "100",
        "pull_limit_interval": "21600"
      },
      "type": "repository"
    }
  ],
  "aud": "registry.docker.io",
  "exp": 1604151763,
  "iat": 1604151463,
  "iss": "auth.docker.io",
  "jti": "iyPEvaDe5zh7u0nm6B4H",
  "nbf": 1604151163,
  "sub": ""
}`
	decodedToken, err := decodeToken(testToken, false)
	assert.NoError(t, err)
	assert.Equal(t, expectedData, decodedToken)
}

func TestCantParse(t *testing.T) {
	testToken := "NotAToken"
	_, err := decodeToken(testToken, false)
	assert.Error(t, err)
}
