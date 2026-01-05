package libs

import "golang.org/x/crypto/bcrypt"

func HashPass(password string) string {
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashPass)
}

func CheckPass(totalPass, passReq string) bool {
	return bcrypt.CompareHashAndPassword([]byte(totalPass), []byte(passReq)) == nil
}
