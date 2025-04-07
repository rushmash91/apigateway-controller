
	// fetch tags
	if r.ko.Status.ID != nil  {
		resourceARN := string(*r.ko.Status.ACKResourceMetadata.ARN)
		tags, err := rm.fetchCurrentTags(ctx, &resourceARN)
		if err != nil {
			return nil, err
		}
		r.ko.Spec.Tags = aws.StringMap(tags)
	}
