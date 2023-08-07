package cmd

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var createToken = &cobra.Command{
	Use:   "generate:token",
	Short: "Generate private and public token",
	Run:   generateTokens,
}


func generateTokens(cmd *cobra.Command, args []string) {

	// Generate keys
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := &privateKey.PublicKey

	// Export private key
	privateKeyPEM := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// Export public key
	publicKeyPEM := pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	}

	saveInDotPem(&privateKeyPEM, &publicKeyPEM)

	os.Exit(1)
}

func saveInDotPem(privateKeyPEM, publicKeyPEM *pem.Block) {
	// Private key
	privFile, err := os.Create("private.pem")
	pem.Encode(privFile, privateKeyPEM)

	if err != nil {
		panic(err)
	}

	privFile.Close()
	log.Println("Private key create successfully in: private.pem")

	// Public key
	pubFile, err := os.Create("public.pem")
	pem.Encode(pubFile, publicKeyPEM)

	if err != nil {
		panic(err)
	}

	
	pubFile.Close()
	log.Println("Public key create successfully in: public.pem")
	
}
