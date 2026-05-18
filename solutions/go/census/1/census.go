package census

type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{name, age, address}
}

func (r *Resident) HasRequiredInfo() bool {
	return r.Name != "" && r.Address != nil && r.Address["street"] != ""
}

func (r *Resident) Delete() {
	if r != nil {
		*r=Resident{}
	}
}
func Count(residents []*Resident) int {
	var count int
	for _, resident := range residents {
		if resident.HasRequiredInfo() {
			count++
		}
	}
	return count
}
