package vcardservice

import (
	"fmt"

	"./main/services/excelService"
)

func GenerateVcardDataService(vcardArr []excelService.Person) string {

	for _, cardData := range vcardArr {

		if cardData.Email == "" || cardData.Name == "" {
			continue
		} else {
			vcardpath := toVCardDataTransform(cardData)
			return vcardpath
		}
	}
	return ""

}

func toVCardDataTransform(cardData excelService.Person) string {
	vCardData := fmt.Sprintf("BEGIN:VCARD\nVERSION:3.0\nFN:%s\nTEL:%s\nEMAIL:%s\nEND:VCARD",
		cardData.Name, cardData.Cell, cardData.Email)

	return vCardData
	//qrservice.GenerateQRCode(vCardData, filename)

}
