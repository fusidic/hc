package characteristic

type Number struct {
    *Characteristic
}

func NewNumber(value, min, max, step interface{}, format string) *Number {    
    n := Number{NewCharacteristic(value, format, CharTypeUnknown, nil)}
    n.MinValue = min
    n.MaxValue = max
    n.MinStep = step
    
    return &n
}

func (n *Number) SetValue(value interface{}) {
    n.Value = value
}

func (n *Number) SetMinValue(value interface{}) {
    n.MinValue = value
}

func (n *Number) SetMaxValue(value interface{}) {
    n.MaxValue = value
}

func (n *Number) SetMinStepValue(value interface{}) {
    n.MinStep = value
}

func (n *Number) GetValue() interface{} {
    return n.Value
}

func (n *Number) GetMinValue() interface{} {
    return n.MinValue
}

func (n *Number) GetMaxValue() interface{} {
    return n.MaxValue
}

func (n *Number) GetMinStepValue() interface{} {
    return n.MinStep
}