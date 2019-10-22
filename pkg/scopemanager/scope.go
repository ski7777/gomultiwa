package scopemanager

type Scope struct {
	Name     string `json:"Name"`
	Value    string `json:"Value"`
	Approved bool   `json:"Approved"`
}

func (s *Scope) EqualsTo(o *Scope) bool {
	if o != nil {
		if s.Name == o.Name {
			return true
		}
		if s.Value == o.Value {
			return true
		}
	}
	return false
}
