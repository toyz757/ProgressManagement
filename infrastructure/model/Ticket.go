package model

import (
  "time"
)

type ticket struct {
  id          int    `json:ticket_id`
  title     string    `json:title`
  responsible     string `json:responsible`
  deadline     *time.Time `json:deadline`
}