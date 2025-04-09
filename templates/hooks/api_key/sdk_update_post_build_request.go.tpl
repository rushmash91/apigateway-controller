    updateApiKeyInput(desired, input, delta)

	// Handle tag updates separately through TagResource/UntagResource APIs
	if delta.DifferentAt("Spec.Tags") {
		if err := updateTags(ctx, rm, desired, latest); err != nil {
			return nil, err
		}
	}

	if !delta.DifferentExcept("Spec.Tags") {
		return desired, nil
	}
