package example

import (
	"context"
	"encoding/json"

	"github.com/mkawserm/abesh/iface"
	"github.com/mkawserm/abesh/model"
)

func ExPostServe(
	ctx context.Context,
	input *model.Event,
	kvStore iface.IKVStore,
	contractId string) (*model.Event, error) {

	type Result struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}

	type Output struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	var key = input.Metadata.Query["key"]
	var result = &Result{}
	var output = &Output{}

	err := json.Unmarshal(input.Value, result)
	if err != nil {
		output.Message = err.Error()
		marshal, err := json.Marshal(output)
		if err != nil {
			return nil, err
		}
		return model.GenerateOutputEvent(
			input.Metadata,
			contractId,
			"NOT OK",
			400,
			"application/json",
			marshal,
		), nil
	}

	err = kvStore.Set(ctx, key, result, 0)
	if err != nil {
		output.Message = err.Error()
		marshal, err := json.Marshal(output)
		if err != nil {
			return nil, err
		}
		return model.GenerateOutputEvent(
			input.Metadata,
			contractId,
			"NOT OK",
			400,
			"application/json",
			marshal,
		), nil
	}

	output.Message = "set value success"
	output.Data = result
	marshal, err := json.Marshal(output)
	if err != nil {
		return nil, err
	}

	return model.GenerateOutputEvent(
		input.Metadata,
		contractId,
		"OK",
		200,
		"application/json",
		marshal,
	), nil
}

func ExServeGet(ctx context.Context, input *model.Event, kvStore iface.IKVStore, contractId string) (*model.Event, error) {
	type Result struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}

	type Output struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	var key = input.Metadata.Query["key"]
	var result = &Result{}
	var output = &Output{}

	var err = kvStore.Get(ctx, key, result)
	if err != nil {
		output.Message = err.Error()
		output.Data = make(map[string]interface{})
		marshal, err := json.Marshal(output)
		if err != nil {
			return nil, err
		}
		return model.GenerateOutputEvent(input.Metadata, contractId, "NOT OK", 400, "application/json", marshal), nil
	}

	output.Message = "get value success"
	output.Data = result
	marshal, err := json.Marshal(output)
	if err != nil {
		return nil, err
	}

	return model.GenerateOutputEvent(input.Metadata, contractId, "OK", 200, "application/json", marshal), nil
}

func ExServeDelete(
	ctx context.Context,
	input *model.Event,
	kvStore iface.IKVStore,
	contractId string) (*model.Event, error) {

	type Output struct {
		Message string `json:"message"`
	}

	var key = input.Metadata.Query["key"]
	var output = &Output{}

	err := kvStore.Delete(ctx, key)
	if err != nil {
		output.Message = err.Error()
		marshal, err := json.Marshal(output)
		if err != nil {
			return nil, err
		}
		return model.GenerateOutputEvent(
			input.Metadata,
			contractId,
			"NOT OK",
			400,
			"application/json",
			marshal,
		), nil
	}

	output.Message = "delete value success"
	marshal, err := json.Marshal(output)
	if err != nil {
		return nil, err
	}

	return model.GenerateOutputEvent(
		input.Metadata,
		contractId,
		"OK",
		200,
		"application/json",
		marshal,
	), nil
}
