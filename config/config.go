package config

import (
  "sort"
)

type LangStruct struct {
	Languages    []string
	LanguagesMap map[string]string
}

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

func Languages() *LangStruct {
	langs := make(map[string]string)

	langs["en"] = "English"
	langs["de"] = "Deutch"
	langs["lt"] = "Lietuvių"

	var lang_keys []string
	for k := range langs {
		lang_keys = append(lang_keys, k)
	}
	sort.Strings(lang_keys)

	var langIndex = &LangStruct{Languages: lang_keys, LanguagesMap: langs}

	return langIndex
}

func Social() *ProviderIndex {
	providers := make(map[string]string)
	providers["facebook"] = "Facebook"
	providers["gplus"] = "Google"
	providers["linkedin"] = "Linkedin"
	providers["twitter"] = "Twitter"

	var provider_keys []string
	for k := range providers {
		provider_keys = append(provider_keys, k)
	}
	sort.Strings(provider_keys)

	var providerIndex = &ProviderIndex{Providers: provider_keys, ProvidersMap: providers}

  return providerIndex
}
