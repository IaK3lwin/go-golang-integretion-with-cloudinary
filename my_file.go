package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func credentials() (*cloudinary.Cloudinary, context.Context) {
	//a url de acessso ao cloudinary está nas variaáveis de ambientes do sistema
	
	//inicializo o client do cloudinary
	cld, err := cloudinary.New()
	if err != nil {
		fmt.Println(err)
	}

	//configuro a url segura | https
	cld.Config.URL.Secure = true

	//crio o contexto
	ctx := context.Background()

	return cld, ctx
}


func uploadImage(fileName string, id string) string {
	//carrego o client e o context do cloudinarty
	cdl, ctx := credentials()

	//tentar abrir a imagem
	arq, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return " "
	}

	//parametros
	params := uploader.UploadParams{
		PublicID: id,
		
	}

	//fazer o upload
	result, err := cdl.Upload.Upload(ctx, arq, params)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return result.SecureURL // manda a ur

}