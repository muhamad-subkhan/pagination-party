package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	productdto "party/dto/product"
	dto "party/dto/result"
	"party/models"
	"party/pkg/middleware"
	"party/repositories"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type handlerProduct struct {
	ProductRepositories repositories.ProductRepositories
}

func HandlerProduct(ProductRepositories repositories.ProductRepositories) *handlerProduct {
	return &handlerProduct{ProductRepositories}
}


func (h *handlerProduct) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var product models.Product

	product, err := h.ProductRepositories.GetProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: product}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProduct) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile")
	filepath := ""
	if dataContex != nil {
		filepath = dataContex.(string)
	}

	price, _ := strconv.Atoi(r.FormValue("price"))
	qty, _ := strconv.Atoi(r.FormValue("qty"))

	request := productdto.ProductRequest{
		Name:  r.FormValue("name"),
		Qty:   qty,
		Price: price,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	var ctx = context.Background()
	// var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	// var API_KEY = os.Getenv("API_KEY")
	// var API_SECRET = os.Getenv("API_SECRET")
	var CLOUD_NAME = "dss399smz"
	var API_KEY = "147652571825927"
	var API_SECRET = "eJoljtlnNkgIlNTsR_bM7Ji2RwA"

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "party"})

	if err != nil {
		fmt.Println(err.Error())
	}

	product := models.Product{
		Name:  request.Name,
		Qty:   request.Qty,
		Price: request.Price,
		Image: resp.SecureURL,
	}

	product, err = h.ProductRepositories.CreateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	product, _ = h.ProductRepositories.GetProduct(product.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: product}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProduct) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile")
	filepath := dataContex.(string)

	price, _ := strconv.Atoi(r.FormValue("price"))
	qty, _ := strconv.Atoi(r.FormValue("qty"))

	request := productdto.UpdateProductRequest{
		Name:  r.FormValue("name"),
		Qty:   qty,
		Price: price,
		Image: filepath,
	}
	fmt.Println(request)
	var ctx = context.Background()
	// var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	// var API_KEY = os.Getenv("API_KEY")
	// var API_SECRET = os.Getenv("API_SECRET")
	var CLOUD_NAME = "dss399smz"
	var API_KEY = "147652571825927"
	var API_SECRET = "eJoljtlnNkgIlNTsR_bM7Ji2RwA"

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "party"})

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("error gaeys")
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	product, err := h.ProductRepositories.GetProduct(id)
	fmt.Println("ajajaj", id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// product.Price = request.Price
	// product.Qty = request.Qty
	if request.Name != "" {
		product.Name = request.Name
	}

	if request.Price != 0 || request.Price <= 0 {
		product.Price = request.Price
	}
	if request.Qty != 0 || request.Qty <= 0 {
		product.Qty = request.Qty
	}

	if filepath != "false" {
		product.Image = resp.SecureURL
	}

	product, err = h.ProductRepositories.UpdateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: product}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerProduct) FindProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	// products, err := h.ProductRepositories.FindProduct()
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	// w.WriteHeader(http.StatusOK)
	// response := dto.SuccessResult{Code: "Succes", Data: products}
	// json.NewEncoder(w).Encode(response)
	
	limit := r.Context().Value(middleware.LimitKey).(int)
	page := r.Context().Value(middleware.PageKey).(int)
	offset := (page - 1) * limit
	
	ProductList, err := h.ProductRepositories.FindProduct(limit, offset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	pageInfo := productdto.PageInfo{
		Size: limit,
		Current: page,
		Total: models.GetSize(),
	}

	pageInfo.TotalPages = int(math.Ceil(float64(pageInfo.Total) / float64(pageInfo.Size)))

	if pageInfo.Current > pageInfo.TotalPages {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: ProductList}
	json.NewEncoder(w).Encode(response)
}
