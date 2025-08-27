package city

type CityService interface {
	GetAll() []City
	GetById(id int) City
	GetByProvinceID(provinceID int) []City
}

type CityServiceImpl struct {
	cityRepository CityRepository
}

func NewCityService(cityRepository CityRepository) CityService {
	return &CityServiceImpl{cityRepository}
}

func (cs *CityServiceImpl) GetAll() []City {
	return cs.cityRepository.FindAll()
}

func (cs *CityServiceImpl) GetById(id int) City {
	return cs.cityRepository.FindById(id)
}

func (cs *CityServiceImpl) GetByProvinceID(provinceID int) []City {
	return cs.cityRepository.FindByProvinceID(provinceID)
}
