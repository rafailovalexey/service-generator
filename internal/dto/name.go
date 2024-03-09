package dto

type NameDto struct {
	LowerCaseFirstLetter   string
	CamelCaseSingular      string
	CamelCasePlural        string
	LowerCamelCaseSingular string
	LowerCamelCasePlural   string
	SnakeCaseSingular      string
	SnakeCasePlural        string
}

func NewNameDto(
	lowerCaseFirstLetter string,
	camelCaseSingular string,
	camelCasePlural string,
	lowerCamelCaseSingular string,
	lowerCamelCasePlural string,
	snakeCaseSingular string,
	snakeCasePlural string,
) *NameDto {
	return &NameDto{
		LowerCaseFirstLetter:   lowerCaseFirstLetter,
		CamelCaseSingular:      camelCaseSingular,
		CamelCasePlural:        camelCasePlural,
		LowerCamelCaseSingular: lowerCamelCaseSingular,
		LowerCamelCasePlural:   lowerCamelCasePlural,
		SnakeCaseSingular:      snakeCaseSingular,
		SnakeCasePlural:        snakeCasePlural,
	}
}
