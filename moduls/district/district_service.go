package district

type DistrictService interface {
	GetAll() []District
	GetById(id int) District
	GetByProvinceID(provinceID int) []District
	GetByCityID(cityID int) []District
}

type DistrictServiceImpl struct{
	districtRepository DistrictRepository
}

func NewDistrictService(districtRepository DistrictRepository) DistrictService{
	return &DistrictServiceImpl{districtRepository}
}

func (ds *DistrictServiceImpl) GetAll() []District{
	return ds.districtRepository.FindAll()
}

func (ds *DistrictServiceImpl) GetById(id int) District{
	return ds.districtRepository.FindById(id)
}

func (ds *DistrictServiceImpl) GetByProvinceID(provinceID int) []District{
	return ds.districtRepository.FindByProvinceID(provinceID)
}

func (ds *DistrictServiceImpl) GetByCityID(cityID int) []District{
	return ds.districtRepository.FindByCityID(cityID)
}
