package pdf

import (
	"log"

	"github.com/signintech/gopdf"
)

type service struct {
}

func NewPDFService() service {
	return service{}
}

type PDF interface {
	GeneratePDF() error
}

func (s service) GeneratePDF() error {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4}) //595.28, 841.89 = A4
	pdf.AddPage()

	if err := pdf.AddTTFFont("THSarabunNew", "TH Sarabun New Regular.ttf"); err != nil {
		log.Print(err.Error())
		return err
	}

	if err := pdf.SetFont("THSarabunNew", "", 20); err != nil {
		log.Print(err.Error())
		return err
	}

	pdf.SetX(10)
	pdf.SetY(10)
	pdf.SetTextColor(156, 197, 140)
	//pdf.SetTextColorCMYK(0, 6, 14, 0)

	//write image
	err := pdf.Image("image.jpeg", 10, 50, nil)
	if err != nil {
		return err
	}

	//write text
	err = pdf.Cell(nil, "picture")
	if err != nil {
		return err
	}
	//write text
	//err = pdf.Text("Hello world!!!")
	//if err != nil {
	//	log.Print(err.Error())
	//	return err
	//}

	//write pdf
	err = pdf.WritePdf("hello.pdf")
	if err != nil {
		log.Print(err.Error())
		return err
	}
	return nil
}
