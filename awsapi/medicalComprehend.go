package awsapi

import (
	"fmt"
	"medicalRecognition/medicine"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehendmedical"
)

func DetectEntities(textFromImage string) medicine.Medicine {

	sess := session.Must(session.NewSession())

	medicalComp := comprehendmedical.New(sess, &aws.Config{
		Region: aws.String("us-east-1"),
	})
	entities, err := medicalComp.DetectEntities(&comprehendmedical.DetectEntitiesInput{
		Text: &textFromImage,
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
	return med
}
