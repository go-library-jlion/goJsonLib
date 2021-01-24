package goJsonLib
import (
"encoding/json"
)

//TrySerialization
func TrySerialization(v interface{}) string {
	result,error:=json.Marshal(v)
	if(error!=nil){
		return ""
	}
	return string(result)
}
