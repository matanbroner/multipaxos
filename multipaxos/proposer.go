package multipaxos

type Proposal struct {
	Value string
	Clock *Clock
}

type Proposer struct {
	ID       string
	Proposal *Proposal
	Promises int
	Clock    *Clock
}

func NewProposer(id string) *Proposer {
	return &Proposer{
		ID:       id,
		Proposal: nil,
		Promises: 0,
		Clock:    NewClock(id),
	}
}

func (p *Proposer) ProposeValue(v string, acceptors []*Acceptor) {
	p.Clock.Incerement()
	p.Proposal = &Proposal{
		Clock: p.Clock,
		Value: v,
	}
	for _, a := range acceptors {
		promise := a.GetProposal(p.Proposal)
		if promise.Reject {
			continue
		}

	}
}
