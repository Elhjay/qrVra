package vcardhandler

import (
	"context"
	"fmt"
	"log"

	"./main/database"

	vcardentity "./main/entities/vcardEntity"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func VCFHandler(c *gin.Context) {
	client := database.GetClient()
	collection := client.Database("qrCode").Collection("employees")
	id := c.Param("id")

	var finalUser vcardentity.FinalUsers
	filter := bson.M{"pid": id}
	err := collection.FindOne(context.Background(), filter).Decode(&finalUser)
	if err != nil {
		log.Fatal(err)
	}
	vcfData := fmt.Sprintf("BEGIN:VCARD\nVERSION:3.0\nFN:%s\nTEL:%s\nEMAIL:%s\nEND:VCARD",
		finalUser.Name, finalUser.Cell, finalUser.Email)

	c.Header("Content-Type", "text/vcard")
	c.Header("Content-Disposition", "attachment; filename="+id+".vcf")

	c.String(200, vcfData)
}

type UpdateStruct struct {
	FieldName  string      `json:"fieldname"`
	FieldValue interface{} `json:"fieldvalue"`
}

func UpdateVCardInfoHandler(c *gin.Context) {
	client := database.GetClient()
	collection := client.Database("qrCode").Collection("employees")

	id := c.Param("id")

	var updateFromReq UpdateStruct
	err := c.BindJSON(&updateFromReq)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(updateFromReq)

	filter := bson.M{"pid": id}
	updater := bson.M{"$set": bson.M{updateFromReq.FieldName: updateFromReq.FieldValue}}

	var finalUser vcardentity.FinalUsers

	err = collection.FindOneAndUpdate(context.Background(), filter, updater).Decode(&finalUser)

	if err != nil {
		log.Fatal(err)
	}

	vcfData := fmt.Sprintf("BEGIN:VCARD\nVERSION:3.0\nFN:%s\nTEL:%s\nEMAIL:%s\nEND:VCARD",
		finalUser.Name, finalUser.Cell, finalUser.Email)

	c.Header("Content-Type", "text/vcard")
	c.Header("Content-Disposition", "attachment; filename="+id+".vcf")

	c.String(200, vcfData)

}
