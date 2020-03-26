package controllers

import (
	"context"

	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/errors"
	usecases "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/contract-registry/contract-registry/use-cases"
	svc "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/types/contract-registry"
)

const component = "contract-registry.controllers.registry"

// ContractRegistryController is a contract registry handler
type ContractRegistryController struct {
	registerUseCase    usecases.RegisterContractUseCase
	getUseCase         usecases.GetContractUseCase
	getMethodsUseCase  usecases.GetMethodsUseCase
	getEventsUseCase   usecases.GetEventsUseCase
	getCatalogUseCase  usecases.GetCatalogUseCase
	getTagsUseCase     usecases.GetTagsUseCase
	setCodeHashUseCase usecases.SetCodeHashUseCase
}

// NewContractRegistryController creates a new contract registry
func NewContractRegistryController(
	registerContractUseCase usecases.RegisterContractUseCase,
	getUseCase usecases.GetContractUseCase,
	getMethodsUseCase usecases.GetMethodsUseCase,
	getEventsUseCase usecases.GetEventsUseCase,
	getCatalogUseCase usecases.GetCatalogUseCase,
	getTagsUseCase usecases.GetTagsUseCase,
	setCodeHashUseCase usecases.SetCodeHashUseCase,
) *ContractRegistryController {
	return &ContractRegistryController{
		registerUseCase:    registerContractUseCase,
		getUseCase:         getUseCase,
		getMethodsUseCase:  getMethodsUseCase,
		getEventsUseCase:   getEventsUseCase,
		getCatalogUseCase:  getCatalogUseCase,
		getTagsUseCase:     getTagsUseCase,
		setCodeHashUseCase: setCodeHashUseCase,
	}
}

// RegisterContract register a contract including ABI, bytecode and deployed bytecode
func (registry *ContractRegistryController) RegisterContract(ctx context.Context, req *svc.RegisterContractRequest) (*svc.RegisterContractResponse, error) {
	err := registry.registerUseCase.Execute(ctx, req.GetContract())
	if err != nil {
		return nil, errors.FromError(err).ExtendComponent(component)
	}

	return &svc.RegisterContractResponse{}, nil
}

// DeregisterContract remove the name + tag association to a contract artifact (abi, bytecode, deployedBytecode). Artifacts are not deleted.
func (registry *ContractRegistryController) DeregisterContract(ctx context.Context, req *svc.DeregisterContractRequest) (*svc.DeregisterContractResponse, error) {
	return nil, errors.FeatureNotSupportedError("DeregisterContract not implemented yet").ExtendComponent(component)
}

// DeleteArtifact remove an artifacts based on its BytecodeHash.
func (registry *ContractRegistryController) DeleteArtifact(ctx context.Context, req *svc.DeleteArtifactRequest) (*svc.DeleteArtifactResponse, error) {
	return nil, errors.FeatureNotSupportedError("DeleteArtifact not implemented yet").ExtendComponent(component)
}

// GetContract loads a contract
func (registry *ContractRegistryController) GetContract(ctx context.Context, req *svc.GetContractRequest) (*svc.GetContractResponse, error) {
	contract, err := registry.getUseCase.Execute(ctx, req.GetContractId())
	if err != nil {
		return nil, errors.FromError(err).ExtendComponent(component)
	}

	return &svc.GetContractResponse{Contract: contract}, nil
}

// GetContractABI loads contract ABI
func (registry *ContractRegistryController) GetContractABI(ctx context.Context, req *svc.GetContractRequest) (*svc.GetContractABIResponse, error) {
	contract, err := registry.getUseCase.Execute(ctx, req.GetContractId())
	if err != nil {
		return nil, errors.FromError(err).ExtendComponent(component)
	}

	return &svc.GetContractABIResponse{Abi: contract.Abi}, nil
}

// GetContractBytecode loads contract bytecode
func (registry *ContractRegistryController) GetContractBytecode(ctx context.Context, req *svc.GetContractRequest) (*svc.GetContractBytecodeResponse, error) {
	contract, err := registry.getUseCase.Execute(ctx, req.GetContractId())
	if err != nil {
		return nil, errors.FromError(err).ExtendComponent(component)
	}

	return &svc.GetContractBytecodeResponse{Bytecode: contract.Bytecode}, nil
}

// GetContractDeployedBytecode loads contract deployed bytecode
func (registry *ContractRegistryController) GetContractDeployedBytecode(ctx context.Context, req *svc.GetContractRequest) (*svc.GetContractDeployedBytecodeResponse, error) {
	contract, err := registry.getUseCase.Execute(ctx, req.GetContractId())
	if err != nil {
		return nil, errors.FromError(err).ExtendComponent(component)
	}

	return &svc.GetContractDeployedBytecodeResponse{DeployedBytecode: contract.DeployedBytecode}, nil
}

// GetMethodsBySelector load method using 4 bytes unique selector and the address of the contract
func (registry *ContractRegistryController) GetMethodsBySelector(ctx context.Context, req *svc.GetMethodsBySelectorRequest) (*svc.GetMethodsBySelectorResponse, error) {
	abi, methodsABI, err := registry.getMethodsUseCase.Execute(ctx, req.GetAccountInstance(), req.GetSelector())

	if err != nil {
		return nil, errors.FromError(err).ExtendComponent(component)
	}

	if abi != "" {
		return &svc.GetMethodsBySelectorResponse{Method: abi}, nil
	}

	return &svc.GetMethodsBySelectorResponse{DefaultMethods: methodsABI}, nil

}

// GetEventsBySigHash load event using event signature hash
func (registry *ContractRegistryController) GetEventsBySigHash(ctx context.Context, req *svc.GetEventsBySigHashRequest) (*svc.GetEventsBySigHashResponse, error) {
	abi, eventsABI, err := registry.getEventsUseCase.Execute(ctx, req.GetAccountInstance(), req.GetSigHash(), req.GetIndexedInputCount())

	if err != nil {
		return nil, errors.FromError(err).ExtendComponent(component)
	}

	if abi != "" {
		return &svc.GetEventsBySigHashResponse{Event: abi}, nil
	}

	return &svc.GetEventsBySigHashResponse{DefaultEvents: eventsABI}, nil
}

// GetCatalog returns a list of all registered contracts.
func (registry *ContractRegistryController) GetCatalog(ctx context.Context, _ *svc.GetCatalogRequest) (*svc.GetCatalogResponse, error) {
	names, err := registry.getCatalogUseCase.Execute(ctx)
	if err != nil {
		return nil, errors.FromError(err).ExtendComponent(component)
	}

	return &svc.GetCatalogResponse{Names: names}, nil
}

// GetTags returns a list of all tags available for a contract name.
func (registry *ContractRegistryController) GetTags(ctx context.Context, req *svc.GetTagsRequest) (*svc.GetTagsResponse, error) {
	names, err := registry.getTagsUseCase.Execute(ctx, req.GetName())
	if err != nil {
		return nil, errors.FromError(err).ExtendComponent(component)
	}

	return &svc.GetTagsResponse{Tags: names}, nil
}

// SetAccountCodeHash set the codehash of a contract address for a given chain
func (registry *ContractRegistryController) SetAccountCodeHash(ctx context.Context, req *svc.SetAccountCodeHashRequest) (*svc.SetAccountCodeHashResponse, error) {
	err := registry.setCodeHashUseCase.Execute(ctx, req.GetAccountInstance(), req.GetCodeHash())
	if err != nil {
		return nil, errors.FromError(err).ExtendComponent(component)
	}

	return &svc.SetAccountCodeHashResponse{}, nil
}
