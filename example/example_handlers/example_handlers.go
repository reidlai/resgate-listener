package example

import (
	rmh "github.com/reidlai/resgate-listener/pkg/resgate_message_handler"
)

type ExampleHandler struct {
	*rmh.BaseResgateMessageHandler
}


func (h *ExampleHandler) AccessResourceHandler() (interface{}, error) {
	return struct{}{}, nil
}

func (h *ExampleHandler) GetResourcesHandler() (interface{}, error) {
	return struct{}{}, nil
}

func (h *ExampleHandler) GetResourceHandler() (interface{}, error) {
	return struct{}{}, nil
}

func (h *ExampleHandler) AddResourceHandler() (interface{}, error) {
	return struct{}{}, nil
}

func (h *ExampleHandler) ChangeResourceHandler() (interface{}, error) {
	return struct{}{}, nil
}

func (h *ExampleHandler) RemoveResourceHandler() (interface{}, error) {
	return struct{}{}, nil
}
