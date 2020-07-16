type ticket struct {
  ticketId          int    `json:ticket_id`
  title     string    `json:title`
  responsible     string `json:responsible`
  deadline     date `json:deadline`
}