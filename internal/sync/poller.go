package sync

import (
	"log"
	"technical-test/internal/service"
	"technical-test/pkg/tzkt"
	"time"
)

// Poller periodically fetches delegations from the Tzkt API
type Poller struct {
	client *tzkt.Client
	svc    *service.DelegationService
}

// NewPoller creates a new Poller instance
func NewPoller(client *tzkt.Client, svc *service.DelegationService) *Poller {
	return &Poller{client: client, svc: svc}
}

// StartPolling begins the polling process at regular intervals
func (p *Poller) StartPolling() {
	ticker := time.NewTicker(1 * time.Minute)
	for {
		select {
		case <-ticker.C:
			if err := p.poll(); err != nil {
				log.Printf("Error during polling: %v", err)
			}
		}
	}
}

// poll fetches new delegations and saves them
func (p *Poller) poll() error {
	delegations, err := p.client.GetDelegations()
	if err != nil {
		return err
	}

	for _, d := range delegations {
		err := p.svc.SaveDelegation(d.Sender.Address, d.Timestamp, d.Amount, d.Level)
		if err != nil {
			return err
		}
	}
	return nil
}
