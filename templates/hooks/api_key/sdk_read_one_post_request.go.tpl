	
	if resp.StageKeys != nil {
		r.ko.Spec.StageKeys = getStageKeysFromStrings(resp.StageKeys)
	} else {
		r.ko.Spec.StageKeys = nil
	} 
