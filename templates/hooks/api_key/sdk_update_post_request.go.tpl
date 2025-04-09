	
	if resp.StageKeys != nil {
		desired.ko.Spec.StageKeys = getStageKeysFromStrings(resp.StageKeys)
	} else {
		desired.ko.Spec.StageKeys = nil
	} 
	