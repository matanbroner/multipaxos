package multipaxos

type Acceptor struct {
	ID            string
	AcceptedValue *string
	PromisedClock *Clock
}

// a promise is returned as a result
// of a proposal.
// a promise either promises a new clock value
// or it returns an accepted value and its clock
type Promise struct {
	Clock  *Clock
	Value  *string
	Reject bool
}

func NewAcceptor(id string) *Acceptor {
	return &Acceptor{
		ID:            id,
		AcceptedValue: nil,
		PromisedClock: nil,
	}
}

// if acceptor has already accepted a value,
// it should return a promise with this value
// and the clock associated with the value
func (a *Acceptor) GetProposal(p *Proposal) *Promise {
	if a.AcceptedValue != nil {
		return &Promise{
			Clock:  a.PromisedClock,
			Value:  a.AcceptedValue,
			Reject: false,
		}
	}
	if a.PromisedClock != nil && a.PromisedClock.IsGreaterThan(p.Clock) {
		return &Promise{
			Reject: true,
		}
	} else {
		a.PromisedClock = p.Clock
		return &Promise{
			Clock:  a.PromisedClock,
			Reject: false,
		}
	}
}

// if an acceptor gets a value to accept with a clock
// that is g.t.o.e to its promised clock, it accepts the value and
// acks it.
// otherwise, it rejects the value
func (a *Acceptor) AcceptValue(c *Clock, v string) (accepted bool) {
	if a.PromisedClock.IsGreaterThan(c) || a.PromisedClock.IsEqualTo(c) {
		a.PromisedClock = c
		a.AcceptedValue = &v
		return true
	} else {
		return false
	}
}
