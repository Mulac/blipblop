package operator

type FSM = FiniteStateMachine

// FiniteStateMachine (or FSM) defines a graph where the 'States' are nodes
// and 'Processors' the out-edges from a node
//
// In the future we can use it however we want, for examle, as a message processing system,
// or build pipeline
//
type FiniteStateMachine interface {
	AddProcessor(State, ...Processor) FSM
	Run() error
}

type State string

type Processor interface {
	Process(m ...Message)
}

type Message interface {
	Evelope() Envelope
	Payload() Payload
}

type Envelope interface {
	Destination() State
	Metadata() Metadata
}

type Metadata map[string]string

type Payload interface {
	Value() interface{}
	Type() PayloadType
}

type PayloadType string
