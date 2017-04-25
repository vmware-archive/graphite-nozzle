package metrics

type StatsdUDPClient struct {
	StatsdSender
}

func NewStatsdUDPClient(conf map[string]string) (sender StatsdClient, err error) {
	statsd_sender, err := newStatsdClient(conf)

	if err != nil {
		return nil, err
	}

	return &StatsdUDPClient{StatsdSender: statsd_sender}, err
}

func (sender *StatsdUDPClient) Connect() (err error) {
	err = sender.CreateSocket()
	return
}

func (sender *StatsdUDPClient) Reconnect() (err error) {
	return sender.Connect()
}
