package enigma

import (
	"strings"
	"testing"
	"time"

	"github.com/ory-am/fosite/enigma/jwthelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMerge(t *testing.T) {
	for k, c := range [][]map[string]interface{}{
		{
			{"foo": "bar"},
			{"baz": "bar"},
			{"foo": "bar", "baz": "bar"},
		},
		{
			{"foo": "bar"},
			{"foo": "baz"},
			{"foo": "bar"},
		},
		{
			{},
			{"foo": "baz"},
			{"foo": "baz"},
		},
		{
			{"foo": "bar"},
			{"foo": "baz", "bar": "baz"},
			{"foo": "bar", "bar": "baz"},
		},
	} {
		assert.EqualValues(t, c[2], merge(c[0], c[1]), "Case %d", k)
	}
}

func TestLoadCertificate(t *testing.T) {
	for _, c := range TestCertificates {
		out, err := LoadCertificate(c[0])
		assert.Nil(t, err)
		assert.Equal(t, c[1], string(out))
	}
	_, err := LoadCertificate("")
	assert.NotNil(t, err)
	_, err = LoadCertificate("foobar")
	assert.NotNil(t, err)
}

func TestRejectsAlgAndTypHeader(t *testing.T) {
	for _, headers := range []map[string]interface{}{
		{"alg": "foo"},
		{"typ": "foo"},
		{"typ": "foo", "alg": "foo"},
	} {
		claims, _ := jwthelper.NewClaimsContext("fosite", "peter", "group0", "",
			time.Now().Add(time.Hour), time.Now(), time.Now(), make(map[string]interface{}))

		j := JWTEnigma{
			PrivateKey: []byte(TestCertificates[0][1]),
			PublicKey:  []byte(TestCertificates[1][1]),
		}
		_, _, err := j.Generate(claims, headers)
		assert.NotNil(t, err)
	}
}

func TestGenerateJWT(t *testing.T) {
	claims, err := jwthelper.NewClaimsContext("fosite", "peter", "group0", "",
		time.Now().Add(time.Hour), time.Now(), time.Now(), make(map[string]interface{}))

	j := JWTEnigma{
		PrivateKey: []byte(TestCertificates[0][1]),
		PublicKey:  []byte(TestCertificates[1][1]),
	}

	token, sig, err := j.Generate(claims, make(map[string]interface{}))
	require.Nil(t, err, "%s", err)
	require.NotNil(t, token)

	sig, err = j.Validate(token)
	require.Nil(t, err, "%s", err)

	sig, err = j.Validate(token + "." + "0123456789")
	require.NotNil(t, err, "%s", err)

	partToken := strings.Split(token, ".")[2]

	sig, err = j.Validate(partToken)
	require.NotNil(t, err, "%s", err)

	// Lets change the public certificate to a different public one...
	j.PublicKey = []byte("new")

	_, err = j.Validate(token)
	require.NotNil(t, err, "%s", err)

	// Reset public key
	j.PublicKey = []byte(TestCertificates[1][1])

	// Lets change the private certificate to a different one...
	j.PrivateKey = []byte("new")
	_, _, err = j.Generate(claims, make(map[string]interface{}))
	require.NotNil(t, err, "%s", err)

	// Reset private key
	j.PrivateKey = []byte(TestCertificates[0][1])

	// Lets validate the exp claim
	claims, err = jwthelper.NewClaimsContext("fosite", "peter", "group0", "",
		time.Now().Add(-time.Hour), time.Now(), time.Now(), make(map[string]interface{}))

	token, sig, err = j.Generate(claims, make(map[string]interface{}))
	require.Nil(t, err, "%s", err)
	require.NotNil(t, token)
	t.Logf("%s.%s", token, sig)

	sig, err = j.Validate(token)
	require.NotNil(t, err, "%s", err)

	// Lets validate the nbf claim
	claims, err = jwthelper.NewClaimsContext("fosite", "peter", "group0", "",
		time.Now().Add(time.Hour), time.Now().Add(time.Hour), time.Now(), make(map[string]interface{}))

	token, sig, err = j.Generate(claims, make(map[string]interface{}))
	require.Nil(t, err, "%s", err)
	require.NotNil(t, token)
	t.Logf("%s.%s", token, sig)

	sig, err = j.Validate(token)
	require.NotNil(t, err, "%s", err)

}

func TestValidateSignatureRejectsJWT(t *testing.T) {
	var err error
	j := JWTEnigma{
		PrivateKey: []byte(TestCertificates[0][1]),
		PublicKey:  []byte(TestCertificates[1][1]),
	}

	for k, c := range []string{
		"",
		" ",
		"foo.bar",
		"foo.",
		".foo",
	} {
		_, err = j.Validate(c)
		assert.NotNil(t, err, "%s", err)
		t.Logf("Passed test case %d", k)
	}
}
