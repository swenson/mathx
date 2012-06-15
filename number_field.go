package ntag

type NumberField struct {
  polynomial *IntPolynomial
}

func MakeNumberField(poly *IntPolynomial) *NumberField {
  k := new(NumberField)
  k.polynomial = poly
  return k
}

func (k *NumberField) Degree() int {
  return len(k.polynomial.coeffs) - 1
}

func (k *NumberField) ClassNumber() int {
  if k.Degree() == 2 {
    return classNumberImagQuad(k)
  }
  return -1
}
