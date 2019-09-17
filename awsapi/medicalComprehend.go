package awsapi

import (
	"fmt"
	"medicalRecognition/medicine"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehendmedical"
)

func DetectEntities(textFromImage string) medicine.Medicine {

	sess := session.Must(session.NewSession())

	medicalComp := comprehendmedical.New(sess, &aws.Config{
		Region: aws.String("us-east-1"),
	})
	textFromImageProcessed := strings.ToLower(textFromImage)

	entities, err := medicalComp.DetectEntities(&comprehendmedical.DetectEntitiesInput{
		Text: &textFromImageProcessed,
	})
	if err != nil {
		fmt.Println(err)
	} else {
		for _, entitie := range entities.Entities {
			if medicine.Existance(*entitie.Text) {
				info := medicine.GetInfo(*entitie.Text)
				return info
			}
		}
	}
	var med medicine.Medicine
	med.Host = "/notfound"
	return med
}
