package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ozinshe/pkg/entity"
	"strconv"
)

func (h *Handler) SignUp(c *gin.Context) {
	var user entity.User
	if err := c.BindJSON(&user); err != nil {
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}
	err := h.svc.PasswordValidator(user.Password)
	if err != nil {
		h.WriteHTTPResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.svc.ConfirmPasswordValidator(user.Password, user.ConfirmPassword)
	if err != nil {
		h.WriteHTTPResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.svc.SignUp(&user)
	if err != nil {
		if err.Error() == entity.AlreadyExist {
			h.WriteHTTPResponse(c, http.StatusBadRequest, "User with this email already exists")
			return
		} else if err.Error() == entity.InvalidEmail {
			h.WriteHTTPResponse(c, http.StatusBadRequest, "invalid email")
			return
		} else {
			h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	h.WriteHTTPResponse(c, http.StatusOK, "user registered with id "+strconv.Itoa(user.Id)+",check your email to verify your account")

}

func (h *Handler) VerifyAccount(c *gin.Context) {
	secretCode := c.Query("link")
	err := h.svc.VerifyAccount(secretCode)
	if err != nil {
		if err.Error() == entity.DidNotFind {
			h.WriteHTTPResponse(c, http.StatusBadRequest, "invalid  verification link")
			return
		} else if err.Error() == entity.ExpiredLink {
			h.WriteHTTPResponse(c, http.StatusBadRequest, "verification link expired")
			return
		}
		h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
	}
	h.WriteHTTPResponse(c, http.StatusOK, "account confirmed!")
}

func (h *Handler) SignIn(c *gin.Context) {
	var credentials entity.Credentials
	if err := c.BindJSON(&credentials); err != nil {
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}
	user, err := h.svc.SigIn(&credentials)
	if err != nil {
		if err.Error() == entity.DidNotFind {
			h.WriteHTTPResponse(c, http.StatusNotFound, "Invalid email: "+credentials.Email)
			return
		} else if err.Error() == entity.NotVerifiedEmail {
			h.WriteHTTPResponse(c, http.StatusBadRequest, "this email is not verified: "+credentials.Email)
			return
		} else if err.Error() == entity.InvalidPassword {
			h.WriteHTTPResponse(c, http.StatusBadRequest, "Invalid password"+credentials.Email)
			return
		}
		h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	role := "user"
	if user.Role == "admin" {
		role = "admin"
	}
	token, err := h.svc.TokenGenerator(user.Id, user.Email, role)
	if err != nil {
		h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) ChangePassword(c *gin.Context) {
	var passwords map[string]interface{}
	if err := c.BindJSON(&passwords); err != nil {
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}
	oldPassword, ok := passwords["oldPassword"].(string)
	if !ok {
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}
	newPassword, ok := passwords["newPassword"].(string)
	if !ok {
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}
	err := h.svc.PasswordValidator(newPassword)
	if err != nil {
		h.WriteHTTPResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	newPasswordConfirm, ok := passwords["newPasswordConfirm"].(string)
	if !ok {
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}
	err = h.svc.ConfirmPasswordValidator(newPassword, newPasswordConfirm)
	if err != nil {
		h.WriteHTTPResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId := c.Value("decodedClaims").(*entity.Claims).Sub
	fmt.Println(oldPassword, newPassword)
	err = h.svc.ChangePasswordByUserId(userId, oldPassword, newPassword)
	if err != nil {
		if err.Error() == entity.InvalidPassword {
			h.WriteHTTPResponse(c, http.StatusBadRequest, "invalid password")
			return
		}
		log.Println(err)
		h.WriteHTTPResponse(c, http.StatusInternalServerError, "")
		return
	}
	h.WriteHTTPResponse(c, http.StatusOK, "password changed")
}

func (h *Handler) PasswordRecover(c *gin.Context) {
	var emailMap map[string]interface{}
	if err := c.BindJSON(&emailMap); err != nil {
		h.WriteHTTPResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}
	email := emailMap["email"].(string)
	err := h.svc.PasswordRecover(email)
	if err != nil {
		h.WriteHTTPResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	h.WriteHTTPResponse(c, http.StatusOK, "email with new password has sent")
}

func (h *Handler) HomePageHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"Message": "Hello,World!",
	})
}
