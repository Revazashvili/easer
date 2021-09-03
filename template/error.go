package template

import "errors"

var ErrTemplateNotFound = errors.New("template not found")
var ErrTemplatesNotFound = errors.New("can't find any template")
var ErrTemplateNotCreated = errors.New("can't create template")
var ErrTemplateNotUpdated = errors.New("can't update template")
