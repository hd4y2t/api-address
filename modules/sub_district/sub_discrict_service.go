package sub_district

type SubDistrictService interface{
	GetAll() []SubDistrict
	GetById(id int) SubDistrict
	GetByProvinceID(provinceID int) []SubDistrict
	GetByCityID(cityID int) []SubDistrict
}

type SubDistrictServiceImpl struct{
	subDistrictRepository SubDistrictRepository
}


func NewSubDistrictService(subDistrictRepository SubDistrictRepository) SubDistrictService{
	return &SubDistrictServiceImpl{subDistrictRepository}
}

func (ds *SubDistrictServiceImpl) GetAll() []SubDistrict{
	return ds.subDistrictRepository.FindAll()
}

func (ds *SubDistrictServiceImpl) GetById(id int) SubDistrict{
	return ds.subDistrictRepository.FindById(id)
}

func (ds *SubDistrictServiceImpl) GetByProvinceID(provinceID int) []SubDistrict{
	return ds.subDistrictRepository.FindByProvinceID(provinceID)
}

func (ds *SubDistrictServiceImpl) GetByCityID(cityID int) []SubDistrict{
	return ds.subDistrictRepository.FindByCityID(cityID)
}
