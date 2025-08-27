package village

type VillageService interface {
	GetAll() []Village
	GetById(id int) Village
	GetByProvinceID(provinceID int) []Village
	GetByCityID(cityID int) []Village
	GetBySubDistrictID(subDistrictID int) []Village
}

type VillageServiceImpl struct {
	villageRepository VillageRepository
}

func NewVillageService(villageRepository VillageRepository) VillageService {
	return &VillageServiceImpl{villageRepository}
}

func (vs *VillageServiceImpl) GetAll() []Village {
	return vs.villageRepository.FindAll()
}

func (vs *VillageServiceImpl) GetById(id int) Village {
	return vs.villageRepository.FindById(id)
}

func (vs *VillageServiceImpl) GetByProvinceID(provinceID int) []Village {
	return vs.villageRepository.FindByProvinceID(provinceID)
}

func (vs *VillageServiceImpl) GetByCityID(cityID int) []Village {
	return vs.villageRepository.FindByCityID(cityID)
}

func (vs *VillageServiceImpl) GetBySubDistrictID(subDistrictID int) []Village {
	return vs.villageRepository.FindBySubDistrictID(subDistrictID)
}
