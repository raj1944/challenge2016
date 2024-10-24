package distributor

import "strings"

type Distributor struct {
	Name           string
	Includes       []string
	Excludes       []string
	SubDistributor []*Distributor
}

func NewDistributor(name string, includes, excludes []string) *Distributor {
	return &Distributor{
		Name:     name,
		Includes: includes,
		Excludes: excludes,
	}
}

func (d *Distributor) AddSubDistributor(sub *Distributor) {
	d.SubDistributor = append(d.SubDistributor, sub)
}

func (d *Distributor) HasPermission(region string) bool {
	region = strings.ToUpper(region)

	for _, exclude := range d.Excludes {
		if strings.HasSuffix(region, exclude) {
			return false
		}
	}

	for _, include := range d.Includes {
		if strings.HasSuffix(region, include) {
			return true
		}
	}

	return false
}
