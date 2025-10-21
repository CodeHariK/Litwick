package nnfs

type _Neuron struct {
	weights []float32
	bias    float32
}

type _Layer struct {
	input_len int
	neurons   []_Neuron
}

func NewLayer(weights [][]float32, bias []float32) (*_Layer, error) {
	if len(weights) != len(bias) {
		return nil, ERROR_WRONG_PARAMETERS
	}
	if len(bias) == 0 {
		return nil, ERROR_EMPTY_PARAMETERS
	}

	var output _Layer
	output.input_len = len(weights[0])
	output.neurons = make([]_Neuron, len(bias))

	for i := 0; i < len(bias); i++ {
		if len(weights[i]) != output.input_len {
			return nil, ERROR_WRONG_PARAMETERS
		}
		output.neurons[i] = _Neuron{
			weights[i],
			bias[i],
		}
	}

	return &output, nil
}

func (n *_Neuron) calc(inputs []float32) float32 {
	output := n.bias
	for i := 0; i < len(inputs); i++ {
		output += inputs[i] * n.weights[i]
	}

	return output
}

func (l *_Layer) Calc(inputs []float32) ([]float32, error) {
	if len(inputs) != l.input_len {
		return nil, ERROR_WRONG_PARAMETERS
	}

	var output []float32 = make([]float32, len(l.neurons))
	for i := 0; i < len(l.neurons); i++ {
		output[i] = l.neurons[i].calc(inputs)
	}
	return output, nil
}
