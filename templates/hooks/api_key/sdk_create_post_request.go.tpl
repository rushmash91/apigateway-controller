	if resp.StageKeys != nil {
		stageKeys := make([]*svcapitypes.StageKey, 0, len(resp.StageKeys))
		for _, stageKeyStr := range resp.StageKeys {
			parts := strings.Split(stageKeyStr, "/")
			if len(parts) == 2 {
				restAPIID := parts[0]
				stageName := parts[1]
				stageKeys = append(stageKeys, &svcapitypes.StageKey{
					RestAPIID: &restAPIID,
					StageName: &stageName,
				})
			}
		}
		desired.ko.Spec.StageKeys = stageKeys
	} else {
		desired.ko.Spec.StageKeys = nil
	} 