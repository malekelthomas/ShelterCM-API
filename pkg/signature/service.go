package signature

type SignatureRepository interface {
	Create(signature *Signature) (*Signature, error)
}

type SignatureService interface {
	Create(signature *Signature) (*Signature, error)
}

type service struct {
	sr SignatureRepository
}

func NewSignatureService(sr SignatureRepository) SignatureService {
	return &service{sr: sr}
}

func (s *service) Create(signature *Signature) (*Signature, error) {
	return s.sr.Create(signature)
}
