package handler

import (
	"golang/campaign"
	"golang/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context){
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error tp get campaigns", http.StatusBadRequest, "error", nil )
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Campaigns List", http.StatusOK, "success", campaigns )
	c.JSON(http.StatusOK, response)
		
}