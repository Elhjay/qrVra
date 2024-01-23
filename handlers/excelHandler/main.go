package excelhandler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"./database"
	"./services/excelService"
	"./services/qrservice"

	vcardentity "./entities/vcardEntity"

	"github.com/gin-gonic/gin"
)

// func toVCardDataTransform(cardData excelService.Person) string {
// 	vCardData := fmt.Sprintf("BEGIN:VCARD\nVERSION:3.0\nFN:%s\nTEL:%s\nEMAIL:%s\nEND:VCARD",
// 		cardData.Name, cardData.Cell, cardData.Email)

// 	return vCardData
// 	//qrservice.GenerateQRCode(vCardData, filename)

// }

func ExcelFileHandler(c *gin.Context) {
	client := database.GetClient()
	collection := client.Database("qrCode").Collection("employees")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.SaveUploadedFile(file, "./uploads/"+file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	users := excelService.DataFromExcel("./uploads/" + file.Filename)
	var finalUsersArr []vcardentity.FinalUsers

	for _, user := range users {

		if user.Email == "" || user.Name == "" {
			continue
		} else {
			finalUser := vcardentity.FinalUsers{
				Pid:        user.Pid,
				Name:       user.Name,
				Cell:       user.Cell,
				Work:       user.Work,
				Email:      user.Email,
				Department: user.Department,
				Section:    user.Section,
				Link:       fmt.Sprintf("http://localhost:8000/vcf/%v", user.Pid),
				QrLink:     "",
			}

			finalUsersArr = append(finalUsersArr, finalUser)

		}

	}

	if len(finalUsersArr) > 0 {
		for _, link := range finalUsersArr {
			_, err := collection.InsertOne(context.Background(), link)
			if err != nil {
				log.Fatal("error adding users")
			}

			qrservice.GenerateQRCode(link.Link, "./uploads/qrcodes"+link.Pid+".png")
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})

}
