	if err := setResourceIDAnnotation(ko); err != nil {
		return nil, ackerr.NewTerminalError(err)
	}
