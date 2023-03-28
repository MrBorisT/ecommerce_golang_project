package sender

type sender struct {
}

func NewStatusSender() *sender {
	return &sender{}
}

func (s *sender) SendStatusChange(orderID int64, status string) {

}
