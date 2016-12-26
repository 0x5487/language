package language

type Resource struct {
	Locale string
	Name   string
}

type Dict map[string]string

type LanguageService struct {
	dict map[string]map[string]string
}

func NewLanguageService() *LanguageService {
	ls := &LanguageService{}
	ls.dict = map[string]map[string]string{}

	return ls
}

func (ls *LanguageService) AddResource(key string, res Resource) {
	if v, ok := ls.dict[key]; ok {
		v[res.Locale] = res.Name
	} else {
		ls.dict[key] = map[string]string{res.Locale: res.Name}
	}
}

func (ls *LanguageService) Find(key string, locale string) string {
	if val, ok := ls.dict[key]; ok {
		if vv, ok := val[locale]; ok {
			return vv
		}
	}
	return ""
}
