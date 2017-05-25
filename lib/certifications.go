package lib

import (
	"github.com/arreyder/compliance-masonry/lib/certifications"
	"github.com/arreyder/compliance-masonry/lib/common"
)

// LoadCertification struct loads certifications into a Certification struct
// and add it to the main object.
func (ws *localWorkspace) LoadCertification(certificationFile string) error {
	cert, err := certifications.Load(certificationFile)
	if err != nil {
		return err
	}
	ws.certification = cert
	return nil
}

func (ws *localWorkspace) GetCertification() common.Certification {
	return ws.certification
}
