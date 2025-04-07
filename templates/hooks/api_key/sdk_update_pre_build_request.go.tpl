	// Handle tag updates separately through TagResource/UntagResource APIs
	if delta.DifferentAt("Spec.Tags") {
		err = rm.syncTags(ctx, desired, latest)
		if err != nil {
			return nil, err
		}
	}

	// If only tags were different, return early
	if !delta.DifferentExcept("Spec.Tags") {
		return desired, nil
	}
