package medicine

func Pharmacy() []string {
	PharmacyKnowledge := []string{"ibuprofeno", "loperamida"}
	return PharmacyKnowledge
}

type Medicine struct {
	Text        string
	Description string
	LinkRef     string
	Host        string
}

func NewIbuprofeno() Medicine {
	ibuprofeno := Medicine{}
	ibuprofeno.Text = "ibuprofeno"
	ibuprofeno.Description = "ibuprofeno text"
	ibuprofeno.LinkRef = ""
	ibuprofeno.Host = "/ibuprofeno"
	return ibuprofeno
}

func NewLoperamida() Medicine {
	loperamida := Medicine{}
	loperamida.Text = "loperamida"
	loperamida.Description = "loperamida text"
	loperamida.LinkRef = ""
	loperamida.Host = "/loperamida"
	return loperamida
}
