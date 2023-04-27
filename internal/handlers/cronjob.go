package handlers

import (
	"log"

	"github.com/robfig/cron/v3"
)

func (h Handlers) RunCronJobs() {
	c := cron.New()

	// every day when 00.00
	c.AddFunc("0 0 * * *", func() {
		err := h.tempImageService.DeleteTempImagesIsTmpMoreOneDay()
		if err != nil {
			log.Print(err.Error())
		}
	})

	// every day when 01.00
	c.AddFunc("0 1 * * *", func() {
		err := h.tempImageService.DeleteTempImagesIsDeletedMoreOneDay()
		if err != nil {
			log.Print(err.Error())
		}
	})

	// every day when 02.00
	c.AddFunc("0 2 * * *", func() {
		err := h.tempImageService.DeleteTempImagesIsNotTmpNotUsedMoreOneDay()
		if err != nil {
			log.Print(err.Error())
		}
	})

	c.Start()

}
