package service

import (
	"github.com/labstack/echo/v4"
	"github.com/sail-host/cloud/internal/app/dto"
	"github.com/sail-host/cloud/internal/app/model"
)

type DomainService struct {
}

type IDomainService interface {
	CreateDomain(c echo.Context, request dto.CreateDomainRequest) (*dto.BaseResponse, *dto.BaseError)
	UpdateDomain(c echo.Context, id uint, request dto.UpdateDomainRequest) (*dto.BaseResponse, *dto.BaseError)
	DeleteDomain(id uint) *dto.BaseError
	GetListDomain() (*dto.DomainListResponse, *dto.BaseError)
	GetDomainByID(id uint) (*dto.BaseResponse, *dto.BaseError)
}

func NewIDomainService() IDomainService {
	return &DomainService{}
}

func (s *DomainService) CreateDomain(c echo.Context, request dto.CreateDomainRequest) (*dto.BaseResponse, *dto.BaseError) {
	var baseError dto.BaseError
	var baseResponse dto.BaseResponse

	err := c.Validate(request)
	if err != nil {
		baseError.Status = "error"
		baseError.Message = err.Error()
		return &baseResponse, &baseError
	}

	if request.DNSProvider == "cloudflare" && (request.CloudflareZoneID == "" || request.CloudflareAPIKey == "") {
		baseError.Status = "error"
		baseError.Message = "Cloudflare zone id and api key is required"
		return &baseResponse, &baseError
	}

	domain := model.Domain{
		Domain:           request.Domain,
		DNSProvider:      request.DNSProvider,
		CloudflareZoneID: request.CloudflareZoneID,
		CloudflareAPIKey: request.CloudflareAPIKey,
	}

	err = domainRepo.CreateDomain(&domain)
	if err != nil {
		baseError.Status = "error"
		baseError.Message = err.Error()
		return &baseResponse, &baseError
	}

	baseResponse.Status = "success"
	baseResponse.Message = "Domain created successfully"
	return &baseResponse, &baseError
}

func (s *DomainService) UpdateDomain(c echo.Context, id uint, request dto.UpdateDomainRequest) (*dto.BaseResponse, *dto.BaseError) {
	var baseError dto.BaseError
	var baseResponse dto.BaseResponse

	err := c.Validate(request)
	if err != nil {
		baseError.Status = "error"
		baseError.Message = err.Error()
		return &baseResponse, &baseError
	}

	if request.DNSProvider == "cloudflare" && (request.CloudflareZoneID == "" || request.CloudflareAPIKey == "") {
		baseError.Status = "error"
		baseError.Message = "Cloudflare zone id and api key is required"
		return &baseResponse, &baseError
	}

	domain, err := domainRepo.GetDomainByID(id)
	if err != nil {
		baseError.Status = "error"
		baseError.Message = err.Error()
		return &baseResponse, &baseError
	}

	domain.Domain = request.Domain
	domain.DNSProvider = request.DNSProvider
	domain.CloudflareZoneID = request.CloudflareZoneID
	domain.CloudflareAPIKey = request.CloudflareAPIKey

	err = domainRepo.UpdateDomain(domain)
	if err != nil {
		baseError.Status = "error"
		baseError.Message = err.Error()
		return &baseResponse, &baseError
	}

	baseResponse.Status = "success"
	baseResponse.Message = "Domain updated successfully"
	return &baseResponse, &baseError
}

func (s *DomainService) DeleteDomain(id uint) *dto.BaseError {
	var baseError dto.BaseError

	domain, err := domainRepo.GetDomainByID(id)
	if err != nil {
		baseError.Status = "error"
		baseError.Message = err.Error()
		return &baseError
	}

	err = domainRepo.DeleteDomain(domain.ID)
	if err != nil {
		baseError.Status = "error"
		baseError.Message = err.Error()
		return &baseError
	}

	return &baseError
}

func (s *DomainService) GetListDomain() (*dto.DomainListResponse, *dto.BaseError) {
	var baseError dto.BaseError
	var baseResponse dto.DomainListResponse

	domains, err := domainRepo.GetListDomain()
	if err != nil {
		baseError.Status = "error"
		baseError.Message = err.Error()
		return &baseResponse, &baseError
	}

	baseResponse.Status = "success"
	baseResponse.Message = "Domain retrieved successfully"

	for _, domain := range domains {
		baseResponse.Data = append(baseResponse.Data, dto.Domain{
			ID:          domain.ID,
			Domain:      domain.Domain,
			DNSProvider: domain.DNSProvider,
			CreatedAt:   domain.CreatedAt,
			UpdatedAt:   domain.UpdatedAt,
		})
	}

	return &baseResponse, &baseError
}

func (s *DomainService) GetDomainByID(id uint) (*dto.BaseResponse, *dto.BaseError) {
	var baseError dto.BaseError
	var baseResponse dto.BaseResponse

	domain, err := domainRepo.GetDomainByID(id)
	if err != nil {
		baseError.Status = "error"
		baseError.Message = err.Error()
		return &baseResponse, &baseError
	}

	baseResponse.Status = "success"
	baseResponse.Message = "Domain retrieved successfully"
	baseResponse.Data = domain

	return &baseResponse, &baseError
}
