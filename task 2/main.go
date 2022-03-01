package main

import "crypto"

//func GeneratePassword(){
//	rand.Seed(time.Now().UnixNano())
//	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
//		"abcdefghijklmnopqrstuvwxyzåäö")
//	length := 8
//	var b strings.Builder
//	for i := 0; i < length; i++ {
//		b.WriteRune(chars[rand.Intn(len(chars))])
//	}
//	str := b.String()
//	fmt.Println(str)
//
//	Hash(str)
//}
//
//func Hash(str string){
//	hash := sha256.New()
//	hash.Write([]byte(str))
//	fmt.Printf("%x", hash.Sum(nil))
//}

func main() {
	//GeneratePassword()
	crypto.RegisterHash()
}
