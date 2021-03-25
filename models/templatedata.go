package models

//Template data holds data sent from handlers to templates.
type TemplateData struct {
	StringMap map[string]string
	IntMap map[string]int32
	FloatMap map[string]float32
	Data map[string]interface{}
	CSRFToken string
	Flash string
	Warning string
	Error string
}

var TD *TemplateData

func Init(){
	TD = &TemplateData{}
}

func SetTeamplateDataStringMap(keyName, valueName string) {
	if (TD.StringMap == nil) {
		stringMap := make(map[string]string)
		stringMap[keyName] = valueName
		TD.StringMap = stringMap
	} else {
		TD.StringMap[keyName] = valueName
	}

}


func SetTeamplateDataIntMap(keyName string, valueName int32) {
	intMap := make(map[string]int32)
	intMap[keyName] = valueName
	TD.IntMap = intMap
}


func SetTeamplateDataFloatMap(keyName string, valueName float32) {
	floatMap := make(map[string]float32)
	floatMap[keyName] = valueName
	TD.FloatMap = floatMap
}


func SetTeamplateDataCSRFToken(data map[string]interface{}) {
	TD.Data = data
}

func SetTeamplateDataFlash(flashMessage string) {
	TD.Flash = flashMessage
}
func SetTeamplateDataWarning(warningMessage string) {
	TD.Warning = warningMessage
}
func SetTeamplateDataError(errorMessage string) {

	TD.Error = errorMessage
}