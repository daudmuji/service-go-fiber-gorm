package cron

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/robfig/cron/v3"
)

type CronFunc func(ctx *fiber.Ctx) error

type Cron struct {
	robfigCron *cron.Cron
}

func NewCron() *Cron {
	return &Cron{
		robfigCron: cron.New(),
	}
}

// CronHandler wraps cron handler
type CronHandler interface {
	RegisterCron(*Cron)
}

func (c *Cron) AddFunc(name, spec string, cmd CronFunc) {
	_, err := c.robfigCron.AddFunc(spec, c.WrapCronFunc(name, cmd))
	if err != nil {
		log.Println(fmt.Sprintf("[Cron] Error registering cron %s: %+v, skipping...", name, err))
		return
	}

	log.Println(fmt.Sprintf("[Cron] Cron %s is successfully registered, spec: %s", name, spec))
}

func (c *Cron) WrapCronFunc(cronName string, fn CronFunc) func() {
	return func() {
		var ctx *fiber.Ctx
		err := fn(ctx)
		if err != nil {
			log.Println(fmt.Sprintf("[Cron][%s] Cron execution return error: %+v", cronName, err))
			return
		}

		log.Println(fmt.Sprintf("[Cron][%s] Cron is successfully executed", cronName))
	}
}

func (c *Cron) CountEntries() int {
	return len(c.robfigCron.Entries())
}

func (c *Cron) Start() {
	c.robfigCron.Start()
}

func (c *Cron) Stop() {
	c.robfigCron.Stop()
}
