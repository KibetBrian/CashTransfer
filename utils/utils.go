package utils
import(
	"golang.org/x/crypto/bcrypt"
)

//Hash user plain text password
func HashPassword(plainText string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainText), 10)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

//Function to compare hashed password
func CompareHash(plainTextPassword, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainTextPassword))
	return err == nil
}
