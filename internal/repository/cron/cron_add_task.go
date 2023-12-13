package cron

import (
	"github.com/ccHR0305/webstack-go/internal/repository/mysql/model"
	"strings"

	"github.com/spf13/cast"
)

func (s *server) AddTask(task *model.CronTask) {
	spec := "0 " + strings.TrimSpace(task.Spec)
	name := cast.ToString(task.ID)

	s.cron.AddFunc(spec, s.AddJob(task), name)
}
