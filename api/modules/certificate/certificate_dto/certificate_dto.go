package certificate_dto

// CertificateRequest is the create/update payload for a certificate.
type CertificateRequest struct {
	Name    string `json:"name"    binding:"required"`
	Issuer  string `json:"issuer"  binding:"required"`
	Period  string `json:"period"`
	Link    string `json:"link"`
	OrderNo int    `json:"orderNo"`
}
