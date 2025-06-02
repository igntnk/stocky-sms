package models

type SupplyState string

const (
	Created   SupplyState = "created"
	InWork    SupplyState = "in_work"
	Served    SupplyState = "served"
	OnTheRoad SupplyState = "on_the_road"
	Shipped   SupplyState = "shipped"
	Done      SupplyState = "done"
)

type Supply struct {
	Uuid            string      `json:"uuid" yaml:"uuid"`
	Comment         string      `json:"comment" yaml:"comment"`
	CreationDate    string      `json:"creationDate" yaml:"creationDate"`
	DesiredDate     string      `json:"desiredDate" yaml:"desiredDate"`
	Status          SupplyState `json:"status" yaml:"status"`
	ResponsibleUser string      `json:"responsibleUser" yaml:"responsibleUser"`
	Edited          bool        `json:"edited" yaml:"edited"`
	EditedDate      string      `json:"editedDate" yaml:"editedDate"`
	Cost            float64     `json:"cost" yaml:"cost"`
}

type SupplyWithProducts struct {
	Supply   `yaml:",inline" json:",inline"`
	Products []SupplyProduct `json:"products" yaml:"products"`
}

type SupplyProduct struct {
	Product `yaml:",inline" json:",inline" mapstructure:",squash"`
	Amount  float64 `json:"amount" yaml:"amount"`
}
