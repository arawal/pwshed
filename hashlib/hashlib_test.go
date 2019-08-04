package hashlib

import (
	"encoding/base64"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashSHA256(t *testing.T) {
	expected := "8qZk9CXChWG3k63kB2L3Iwl8vPXgpK99lgvebvOXxfyoT1J9SCnPzBxUorEYZsAe+vqArWdOAMChEZR3ng6jOw=="

	actual, err := hashSHA256("securestring")
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Errorf("Expected %s but received %s", expected, actual)
	}
}

func TestHashSHA512(t *testing.T) {
	expected := "8qZk9CXChWG3k63kB2L3Iwl8vPXgpK99lgvebvOXxfyoT1J9SCnPzBxUorEYZsAe+vqArWdOAMChEZR3ng6jOw=="

	actual, err := hashSHA512("securestring")
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Errorf("Expected %s but received %s", expected, actual)
	}
}

func TestHashMD5(t *testing.T) {
	expected := "Bec6C38G3VZ04CbVNIUieQ=="

	actual, err := hashMD5("securestring")
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Errorf("Expected %s but received %s", expected, actual)
	}
}

func TestHashSHA3(t *testing.T) {
	expected := "chBQTKUoivgDzB3H9zDrIjYsVJvFhwGZ1ZwI1ZsQecttcTcoOWk07K1SyPfhfzsNf6XmBys0stnbQhHGku8qgw=="

	actual, err := hashSHA3("securestring")
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Errorf("Expected %s but received %s", expected, actual)
	}
}

func TestHashBcrypt(t *testing.T) {
	result, err := hashBcrypt("securestring")
	if err != nil {
		t.Fatal(err)
	}

	actual, err := base64.StdEncoding.DecodeString(result)
	if err != nil {
		t.Fatal(err)
	}

	err = bcrypt.CompareHashAndPassword(actual, []byte("securestring"))

	if err != nil {
		t.Errorf("Expected no error but received %s", err.Error())
	}
}

func TestHash(t *testing.T) {
	t.Run("valid_alg", func(*testing.T) {
		expected := "Bec6C38G3VZ04CbVNIUieQ=="

		actual, err := Hash("securestring", "MD5")
		if err != nil {
			t.Fatal(err)
		}

		if actual != expected {
			t.Errorf("Expected %s but received %s", expected, actual)
		}
	})
	t.Run("invalid_alg", func(*testing.T) {
		result, err := Hash("securestring", "SOMERANDOMALG1")
		if err != nil {
			t.Fatal(err)
		}

		actual, err := base64.StdEncoding.DecodeString(result)
		if err != nil {
			t.Fatal(err)
		}

		err = bcrypt.CompareHashAndPassword(actual, []byte("securestring"))

		if err != nil {
			t.Errorf("Expected no error but received %s", err.Error())
		}
	})
	t.Run("missing_alg", func(*testing.T) {
		result, err := Hash("securestring", "")
		if err != nil {
			t.Fatal(err)
		}

		actual, err := base64.StdEncoding.DecodeString(result)
		if err != nil {
			t.Fatal(err)
		}

		err = bcrypt.CompareHashAndPassword(actual, []byte("securestring"))

		if err != nil {
			t.Errorf("Expected no error but received %s", err.Error())
		}
	})
	t.Run("missing_password", func(*testing.T) {
		_, err := Hash("", "")
		if err == nil {
			t.Errorf("Expected error but got none")
		}

		if err.Error() != "no password provided in request" {
			t.Errorf("Unexpected error: %s", err.Error())
		}
	})
}
