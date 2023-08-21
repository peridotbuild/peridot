package starlark

type Stage string

const (
	StageDevelopment Stage = "development"
	StageStaging     Stage = "staging"
	StageProduction  Stage = "production"
)

var validStages = []Stage{
	StageDevelopment,
	StageStaging,
	StageProduction,
}

func GetValidStages() []Stage {
	return validStages
}
