package GoAPIClient

import "fmt"

type BrawlersResponse struct {
	Brawlers []Brawler `json:"items"`
}

type Brawler struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	StarPowers []StarPower `json:"starPowers"`
	Gadgets    []Gadget    `json:"gadget"`
}

type StarPower struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Gadget struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (b Brawler) Info() string {
	return fmt.Sprintf("[ID] %d | [NAME] %s", b.ID, b.Name)
}
