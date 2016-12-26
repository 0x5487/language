package language

var (
	ls *LanguageService
)

func init() {
	ls = NewLanguageService()
}

func AddResource(key string, res Resource) {
	ls.AddResource(key, res)
}

func Find(key string, locale string) string {
	return ls.Find(key, locale)
}
