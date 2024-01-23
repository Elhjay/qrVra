package qrservice

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

func GenerateQRCode(data string, filename string) {
	// vCard (VCF) data as a string
	// 	vCardData := `BEGIN:VCARD
	// VERSION:3.0
	// FN:John Doe
	// TEL:+123456789
	// EMAIL:john.doe@example.com
	// END:VCARD`

	// Generate QR code
	qrCode, err := qrcode.New(data, qrcode.Medium)
	if err != nil {
		fmt.Println("Error generating QR code:", err)
		return
	}

	// Save the QR code to a file
	err = qrCode.WriteFile(256, filename)
	if err != nil {
		fmt.Println("Error saving QR code:", err)
		return
	}

	fmt.Println("QR code with vCard data generated successfully and saved to qrcode_vcard.png!")
}
