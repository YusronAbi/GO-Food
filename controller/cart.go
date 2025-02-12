package controller

import (
	"fmt"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	repository  entity.CartRepository
	menuService entity.MenuService
	cartService entity.CartService
}

func NewCartHandler(repository entity.CartRepository, menuService entity.MenuService, cartService entity.CartService) *CartHandler {
	return &CartHandler{repository, menuService, cartService}
}

func (h *CartHandler) GetMany(ctx *gin.Context) {
	userID, err := helper.GetUserIDFromCookie(ctx)
	if err != nil {
		helper.UnauthorizedResponse(ctx, err)
		return
	}

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	carts, totalItems, err := h.repository.GetMany(ctx, userID, page, limit)
	if err != nil {
		helper.InternalServerError(ctx, "Failed to fetch data")
		return
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))
	if page > totalPages {
		helper.NotFoundResponse(ctx, "Data not found")
		return
	}

	helper.PaginationResponse(ctx, "Fetch data successfully", page, limit, totalPages, totalItems, carts)
}

func (h *CartHandler) CalculateTotalPrice(ctx *gin.Context) {
	userID, err := helper.GetUserIDFromCookie(ctx)
	if err != nil {
		helper.UnauthorizedResponse(ctx, err)
		return
	}

	calculation, err := h.cartService.CalculatePrice(ctx, uint(userID), "pending")
	if err != nil {
		helper.InternalServerError(ctx, fmt.Sprintf("Failed to calculate total price: %v", err))
		return
	}

	helper.SuccessResponse(ctx, "Total price calculated successfully", calculation)
}

func (h *CartHandler) CreateOne(ctx *gin.Context) {
	var cart entity.Cart
	if err := ctx.ShouldBindJSON(&cart); err != nil {
		helper.BadRequestResponse(ctx, "Invalid request data")
		return
	}

	userID, err := helper.GetUserIDFromCookie(ctx)
	if err != nil {
		helper.UnauthorizedResponse(ctx, err)
		return
	}

	cart.UserID = userID
	existingCart, _ := h.repository.FindByUserAndMenuAndStatus(ctx, userID, cart.MenuID, "pending")
	if existingCart != nil {
		existingCart.Quantity += cart.Quantity
		existingCart.Subtotal, err = h.menuService.CalculateSubTotal(ctx, cart.MenuID, existingCart.Quantity)
	} else {
		cart.Subtotal, err = h.menuService.CalculateSubTotal(ctx, cart.MenuID, cart.Quantity)
	}

	if err != nil || h.menuService.DecreaseMenu(ctx, cart.MenuID, cart.Quantity) != nil {
		helper.BadRequestResponse(ctx, "Please check quantity")
		return
	}

	if existingCart != nil {
		updatedCart, err := h.repository.UpdateOne(ctx, existingCart)
		if err == nil {
			helper.SuccessResponse(ctx, "Cart updated successfully", updatedCart)
		}
	} else {
		createCart, err := h.repository.CreateOne(ctx, &cart)
		if err == nil {
			helper.SuccessResponse(ctx, "Create data successfully", createCart)
		}
	}
	if err != nil {
		helper.InternalServerError(ctx, "Failed to process cart")
	}
}

func (h *CartHandler) UpdateQuantity(ctx *gin.Context, increase bool) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.BadRequestResponse(ctx, "Invalid ID")
		return
	}

	qty := 1
	var updateErr error
	if increase {
		updateErr = h.cartService.IncreaseCart(ctx, uint(id), qty)
	} else {
		updateErr = h.cartService.DecreaseCart(ctx, uint(id), qty)
	}

	if updateErr != nil {
		helper.InternalServerError(ctx, "Failed to update quantity")
		return
	}

	msg := "Increase quantity successfully"
	if !increase {
		msg = "Decrease quantity successfully"
	}
	helper.SuccessResponse(ctx, msg, nil)
}

func (h *CartHandler) DeleteOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.BadRequestResponse(ctx, "Invalid ID")
		return
	}

	if err := h.repository.DeleteOne(ctx, uint(id)); err != nil {
		helper.InternalServerError(ctx, "Failed to delete data")
		return
	}

	helper.SuccessResponse(ctx, "Delete data successfully", nil)
}

func (h *CartHandler) GetOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helper.BadRequestResponse(ctx, "Invalid ID")
		return
	}

	cart, err := h.repository.GetOne(ctx, uint(id))
	if err != nil {
		helper.InternalServerError(ctx, "Failed to fetch data")
		return
	}

	helper.SuccessResponse(ctx, "Fetch data successfully", cart)
}
