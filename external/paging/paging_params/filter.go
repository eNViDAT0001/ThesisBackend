package paging_params

type filterList struct {
	Fields  *map[string]string
	Search  *map[string]string
	Sort    *map[string]string
	Compare *map[string]string
}
type FilterList interface {
	GetFields() *map[string]string
	GetSearch() *map[string]string
	GetSort() *map[string]string
	GetCompare() *map[string]string

	CloneWithSort([]string) FilterList
	CloneWithField([]string) FilterList
}
type FilterBuilder interface {
	WithSorts([]string) FilterBuilder
	WithSearch([]string) FilterBuilder
	WithFields([]string) FilterBuilder
	WithCompare([]string) FilterBuilder
	Build() FilterList
}

func (f *filterList) GetFields() *map[string]string {
	return f.Fields
}
func (f *filterList) GetSearch() *map[string]string {
	return f.Search
}
func (f *filterList) GetSort() *map[string]string {
	return f.Sort
}
func (f *filterList) GetCompare() *map[string]string {
	return f.Compare
}
func (f *filterList) WithFields(input []string) FilterBuilder {
	f.Fields = convertParamsArrayToMap(input)
	return f
}
func (f *filterList) WithSearch(input []string) FilterBuilder {
	f.Search = convertParamsArrayToMap(input)
	return f
}
func (f *filterList) WithSorts(input []string) FilterBuilder {
	f.Sort = convertParamsArrayToMap(input)
	return f
}
func (f *filterList) WithCompare(input []string) FilterBuilder {
	f.Compare = convertParamsArrayToMap(input)
	return f
}
func (f *filterList) Build() FilterList {
	return f
}
func (f *filterList) CloneWithSort(input []string) FilterList {
	return &filterList{
		Search: f.Search,
		Fields: f.Fields,
		Sort:   convertParamsArrayToMap(input),
	}
}
func (f *filterList) CloneWithField(input []string) FilterList {
	return &filterList{
		Search: f.Search,
		Fields: convertParamsArrayToMap(input),
		Sort:   f.Sort,
	}
}
func NewFilterBuilder() FilterBuilder {
	return &filterList{}
}

func convertParamsArrayToMap(input []string) *map[string]string {
	result := make(map[string]string)

	for _, v := range input {
		index := len(v) - 1
		for index >= 0 {
			if v[index] == '_' {
				break
			}
			index--
		}

		if index == -1 {
			continue
		}

		key := v[:index]
		value := v[index+1:]
		result[key] = value
	}
	if len(result) == 0 {
		return nil
	}
	return &result
}
