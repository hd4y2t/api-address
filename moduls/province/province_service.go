package province

type ProvinceService interface{
	GetAll() []Province
	GetById(id int) Province
}

type ProvinceServiceImpl struct{
	provinceRepository ProvinceRepository
}


func NewProvinceService(provinceRepository ProvinceRepository) ProvinceService{
	return &ProvinceServiceImpl{provinceRepository}
}

func (ps *ProvinceServiceImpl) GetAll() []Province{
	return ps.provinceRepository.FindAll()
}

func (ps *ProvinceServiceImpl) GetById(id int) Province{
	return ps.provinceRepository.FindById(id)
}