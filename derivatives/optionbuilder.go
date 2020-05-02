package derivatives

// OptionBuilder is the scaffold for the builder pattern code
type OptionBuilder struct {
	option *Contract
}

// NewOptionBuilder begins the sequence to instantiate a new Call or Put
func NewOptionBuilder() *OptionBuilder {
	option := &Contract{}
	b := &OptionBuilder{option: option}
	return b
}

// WithUnderlying sets the underlying asset for the option
func (b *OptionBuilder) WithUnderlying(u Underlying) *OptionBuilder {
	b.option.Underlying = u
	return b
}

// Strike is the strike price of the option
func (b *OptionBuilder) Strike(k float64) *OptionBuilder {
	b.option.k = k
	return b
}

// TTE is the time to expiration in years
func (b *OptionBuilder) TTE(t float64) *OptionBuilder {
	b.option.t = t
	return b
}

// Rate is the risk free rate
func (b *OptionBuilder) Rate(r float64) *OptionBuilder {
	b.option.r = r
	return b
}

// Call completes the OptionBuilder sequence and returns the
// option with option type set to Call
func (b *OptionBuilder) Call() *Contract {
	b.option.ContractType = CALL
	return b.option
}

// Put completes the OptionBuilder sequence and returns the
// option with type set to Put
func (b *OptionBuilder) Put() *Contract {
	b.option.ContractType = PUT
	return b.option
}
