package helper

import "testing"

// nama func harus diawali dengan Test
func TestHelloWorld(t *testing.T) {
	result := HelloWorld("ardin")

	if result != "Hello "+"ardin" {
		panic("Result is not Hello ardin")
	}
}

func TestHello2(t *testing.T) {
	result := HelloWorld("ardin")

	if result != "Hello "+"ardin" {
		panic("Result is not Hello ardin")
	}
}

// cara menjalankan file testing
// apabila ingin me run test dengan perintah : go test -v, atau go test ./... apabila ada di root folder
// dan tidak mau masuk ke directory testnya
// apabila ingin menjalankan test pada function tertentu bisa menggunakan go test -v run TestNamaFunction
