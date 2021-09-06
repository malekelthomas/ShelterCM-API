package postgres

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/malekelthomas/ShelterCM-API/pkg/config"
	"github.com/malekelthomas/ShelterCM-API/pkg/signature"
	"github.com/skip2/go-qrcode"
	"gorm.io/gorm"
)

type SignatureStore struct {
	db *gorm.DB
}

func NewSignatureStore() *SignatureStore {
	var o sync.Once
	o.Do(func() {
		if err := gdb.AutoMigrate(&signature.Signature{}); err != nil {
			panic(err)
		}
	})
	return &SignatureStore{
		db: gdb,
	}
}

func (ss *SignatureStore) Create(signature *signature.Signature) (*signature.Signature, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	uploader := s3manager.NewUploader(sess)
	png, err := qrcode.Encode(signature.UserID.String(), qrcode.Highest, 256)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(png)
	s3res, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: &config.Config.AWS.SignatureBucket,
		Key:    aws.String(fmt.Sprintf("signatures/%s-%v.png", signature.UserID.String(), signature.SignatureType)),
		Body:   r,
	})
	if err != nil {
		return nil, err
	}
	signature.Source = s3res.Location

	result := ss.db.Create(&signature)
	if result.Error != nil {
		return nil, result.Error
	}

	return signature, nil
}
