package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"html/template"
	"io"
	"math/big"

	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sohag1990/alc_shared/models"
)

func InvIDGenerator(order models.Order) string {

	invID := GetYearFromTime(order.CreatedAt) + "-" + Uint64ConvertToString(order.InvID)
	return invID
}
func StringToInt(str string) int {
	intVal, _ := strconv.Atoi(str)
	return intVal
}
func Uint64ConvertToString(value uint64) string {
	return strconv.FormatUint(value, 10)
}

func GetYearFromTime(t *time.Time) string {
	if t != nil {
		return t.Format("2006")
	}
	return ""
}
func GetGetFormattedDate(t *time.Time) string {
	if t != nil {
		return t.Format("1/2/2006")
	}
	return ""

}
func GetGetFormattedInvDate(order models.Order) string {
	if order.CreatedAt != nil {
		return order.CreatedAt.Format("1/2/2006")
	}
	return ""
}
func Noescape(str string) template.HTML {
	return template.HTML(str)
}

func TokenGenerator() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// GeneratePassword generates a random password of the specified length and complexity level.
// Complexity levels are 1, 2, and 3, with 1 being the simplest and 3 being the most complex.
func GeneratePassword(length, complexity int) (string, error) {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var numberRunes = []rune("0123456789")
	var symbolRunes = []rune("!@#$%^&*()_+{}[]|\\;:'\"<>,.?/")

	var password []rune
	for i := 0; i < length; i++ {
		var err error
		var r rune
		switch complexity {
		case 1:
			r, err = randomRune(letterRunes)
		case 2:
			if n, err := cryptoRandInt(2); err != nil {
				return "", err
			} else if n == 0 {
				r, err = randomRune(letterRunes)
			} else {
				r, err = randomRune(numberRunes)
			}
		case 3:
			if n, err := cryptoRandInt(3); err != nil {
				return "", err
			} else if n == 0 {
				r, err = randomRune(letterRunes)
			} else if n == 1 {
				r, err = randomRune(numberRunes)
			} else {
				r, err = randomRune(symbolRunes)
			}
		}
		if err != nil {
			return "", err
		}
		password = append(password, r)
	}

	return string(password), nil
}
func randomRune(runes []rune) (rune, error) {
	n, err := cryptoRandInt(len(runes))
	if err != nil {
		return 0, err
	}
	return runes[n], nil
}

func cryptoRandInt(max int) (int, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return int(nBig.Int64()), nil
}

var (
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// IsValidEmailFormat Email invalid formate checker
func IsValidEmailFormat(email string) bool {
	if !emailRegexp.MatchString(email) {
		return false
	}
	return true
}

func UrlSchema(c *gin.Context) string {
	isHTTPS := c.Request.TLS != nil

	// Determine the URL schema based on the HTTPS flag
	scheme := "http"
	if isHTTPS {
		scheme = "https"
	} else {
		scheme = "http"
	}

	if strings.Split(c.Request.Host, ":")[0] == "localhost" {
		scheme = "http"
	} else {
		scheme = "https"
	}

	return scheme
}
func GetHosturl(c *gin.Context) string {
	return UrlSchema(c) + "://" + c.Request.Host
}
func GetHostWithoutProtocol(c *gin.Context) string {
	return c.Request.Host
}

// Pluralize any word also you can add irregular plural words preset.
func Pluralize(word string) string {
	// Define some irregular plural forms
	irregularPlurals := map[string]string{
		"branch": "branches",
		// Add more irregular plurals if needed
	}

	// Check if the word is in the irregular plurals map
	if plural, ok := irregularPlurals[strings.ToLower(word)]; ok {
		return plural
	}

	// Apply some basic rules for pluralization
	if strings.HasSuffix(word, "s") || strings.HasSuffix(word, "x") || strings.HasSuffix(word, "z") {
		return word + "es"
	} else if strings.HasSuffix(word, "y") {
		// Change "y" to "ies" if the word ends with a consonant
		if strings.ContainsAny(string(word[len(word)-2]), "aeiou") {
			return word + "s"
		} else {
			return word[:len(word)-1] + "ies"
		}
	} else {
		return word + "s"
	}
}

// Declare the key as a global variable
var key = []byte("R&>+DLHp+Rp!XP}9C%;X<yw_YJbTH}Z-$S~$4R!WY$p}9`u$pW/(*G-a8umz -+$")

// Encrypt encrypts plaintext using the given key with AES-GCM.
func Encrypt(plaintext []byte) (string, error) {
	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Use Galois Counter Mode (GCM) for encryption
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Create a nonce for the encryption
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Encrypt the plaintext using AES-GCM
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)

	// Return the encrypted text as a hexadecimal string
	return hex.EncodeToString(ciphertext), nil
}

// Decrypt decrypts ciphertext using the given key with AES-GCM.
func Decrypt(ciphertextHex string) (string, error) {
	// Decode the hexadecimal string into bytes
	ciphertext, err := hex.DecodeString(ciphertextHex)
	if err != nil {
		return "", err
	}

	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Use Galois Counter Mode (GCM) for decryption
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Extract the nonce from the ciphertext
	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Decrypt the ciphertext using AES-GCM
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	// Return the original plaintext as a string
	return string(plaintext), nil
}
