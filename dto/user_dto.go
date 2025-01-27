package dto

import (
	"errors"
	"mime/multipart"

	"github.com/Caknoooo/go-gin-clean-starter/entity"
)

const (
	// Failed
	MESSAGE_FAILED_GET_DATA_FROM_BODY      = "failed get data from body"
	MESSAGE_FAILED_REGISTER_USER           = "failed create user"
	MESSAGE_FAILED_GET_LIST_USER           = "failed get list user"
	MESSAGE_FAILED_GET_USER_TOKEN          = "failed get user token"
	MESSAGE_FAILED_TOKEN_NOT_VALID         = "token not valid"
	MESSAGE_FAILED_TOKEN_NOT_FOUND         = "token not found"
	MESSAGE_FAILED_GET_USER                = "failed get user"
	MESSAGE_FAILED_LOGIN                   = "failed login"
	MESSAGE_FAILED_WRONG_EMAIL_OR_PASSWORD = "wrong email or password"
	MESSAGE_FAILED_UPDATE_USER             = "failed update user"
	MESSAGE_FAILED_DELETE_USER             = "failed delete user"
	MESSAGE_FAILED_PROSES_REQUEST          = "failed proses request"
	MESSAGE_FAILED_DENIED_ACCESS           = "denied access"
	MESSAGE_FAILED_VERIFY_EMAIL            = "failed verify email"
	MESSAGE_FAILED_RESET_PASSWORD          = "failed reset password"
	MESSAGE_FAILED_EMAIL_NOT_FOUND         = "email not found"
	MESSAGE_FAILED_TOKEN_EXPIRED           = "token expired"
	MESSAGE_FAILED_FORGET_PASSWORD         = "failed handle forget password"
	MESSAGE_FAILED_TOKEN_NOT_ALLOWED       = "token not allowed"

	// Success
	MESSAGE_SUCCESS_REGISTER_USER           = "success create user"
	MESSAGE_SUCCESS_GET_LIST_USER           = "success get list user"
	MESSAGE_SUCCESS_GET_USER                = "success get user"
	MESSAGE_SUCCESS_LOGIN                   = "success login"
	MESSAGE_SUCCESS_UPDATE_USER             = "success update user"
	MESSAGE_SUCCESS_DELETE_USER             = "success delete user"
	MESSAGE_SEND_VERIFICATION_EMAIL_SUCCESS = "success send verification email"
	MESSAGE_SUCCESS_VERIFY_EMAIL            = "success verify email"
	MESSAGE_SUCCESS_RESET_PASSWORD          = "success reset password"
	MESSAGE_SUCCESS_FORGET_PASSWORD         = "success handle forget password"
)

var (
	ErrCreateUser             = errors.New("failed to create user")
	ErrGetAllUser             = errors.New("failed to get all user")
	ErrGetUserById            = errors.New("failed to get user by id")
	ErrGetUserByEmail         = errors.New("failed to get user by email")
	ErrEmailAlreadyExists     = errors.New("email already exist")
	ErrUpdateUser             = errors.New("failed to update user")
	ErrUserNotAdmin           = errors.New("user not admin")
	ErrUserNotFound           = errors.New("user not found")
	ErrEmailNotFound          = errors.New("email not found")
	ErrDeleteUser             = errors.New("failed to delete user")
	ErrPasswordNotMatch       = errors.New("password not match")
	ErrEmailOrPassword        = errors.New("wrong email or password")
	ErrAccountNotVerified     = errors.New("account not verified")
	ErrTokenInvalid           = errors.New("token invalid")
	ErrTokenExpired           = errors.New("token expired")
	ErrAccountAlreadyVerified = errors.New("account already verified")
	ErrHashPasswordFailed     = errors.New("failed to hash password")
	ErrParseUUID              = errors.New("error parsing uuid")
	ErrUserIdEmpty            = errors.New("user id empty")
	ErrRoleNotAllowed         = errors.New("role not allowed: ")
)

type (
	UserCreateRequest struct {
		Name             string `json:"name" form:"name" binding:"required,max=700"`
		Email            string `json:"email" form:"email" binding:"required,max=100"`
		Password         string `json:"password" form:"password" binding:"required,max=100"`
		Institution      string `json:"institusi" form:"institusi" binding:"required,max=400"`
		TelpNumber       string `json:"no_telp" form:"no_telp" binding:"required,max=30"`
		InfoFrom         string `json:"info_from" form:"info_from" binding:"required,max=200"`
		Province         string `json:"asal_provinsi" form:"asal_provinsi" binding:"required,max=100"`
		EducationalLevel string `json:"jenjang" form:"jenjang" binding:"required,max=100"`
	}

	UserResponse struct {
		Name          string  `json:"name" form:"name" binding:"required"`
		Email         string  `json:"email" form:"email" binding:"required"`
		Instansi      string  `json:"instansi" form:"instansi" binding:"required"`
		NoTelp        string  `json:"no_telp" form:"no_telp" binding:"required"`
		InfoFrom      string  `json:"info_from" form:"info_from" binding:"required"`
		Provinsi      string  `json:"asal_provinsi" form:"asal_provinsi" binding:"required"`
		Jenjang       string  `json:"jenjang" form:"jenjang" binding:"required"`
		OCId          *string `json:"oc"`
		OCStatus      *string `json:"oc_status"`
		WSNId         *string `json:"wsn"`
		WSNStatus     *string `json:"wsn_status"`
		WSNCode       string  `json:"wsn_code"`
		WSNCodeStatus string  `json:"wsn_code_status"`
		ILJMajor      string  `json:"ilj_major"`
		ILJSubmajor   string  `json:"ilj_submajor"`
		Role          string  `json:"role"`
	}

	UserPaginationResponse struct {
		Data []UserResponse `json:"data"`
		pagination.Meta
	}

	GetAllUserRepositoryResponse struct {
		Users []entity.User
		pagination.Meta
	}

	UserUpdateRequest struct {
		Name          string `json:"name" form:"name" binding:"max=700"`
		Instansi      string `json:"instansi" form:"instansi" binding:"max=400"`
		No_Telp       string `json:"no_telp" form:"no_telp" binding:"max=100"`
		Info_From     string `json:"info_from" form:"info_from" binding:"max=100"`
		Asal_Provinsi string `json:"asal_provinsi" form:"asal_provinsi" binding:"max=100"`
		Jenjang       string `json:"jenjang" form:"jenjang" binding:"max=100"`
	}

	UserUpdateResponse struct {
		Name          string `json:"name" form:"name"`
		Email         string `json:"email" form:"email"`
		Instansi      string `json:"instansi" form:"instansi"`
		No_Telp       string `json:"no_telp" form:"no_telp"`
		Info_From     string `json:"info_from" form:"info_from"`
		Asal_Provinsi string `json:"asal_provinsi" form:"asal_provinsi"`
		Jenjang       string `json:"jenjang" form:"jenjang"`
	}

	SendVerificationEmailRequest struct {
		Email string `json:"email" form:"email" binding:"required,max=100"`
	}

	VerifyEmailRequest struct {
		Token string `json:"token" form:"token" binding:"required"`
	}

	VerifyEmailResponse struct {
		Email      string `json:"email"`
		IsVerified bool   `json:"is_verified"`
	}

	UserLoginRequest struct {
		Email    string `json:"email" form:"email" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	UserLoginResponse struct {
		Token string `json:"token"`
		Role  string `json:"role"`
	}

	UpdateStatusIsVerifiedRequest struct {
		UserId     string `json:"user_id" form:"user_id" binding:"required"`
		IsVerified bool   `json:"is_verified" form:"is_verified"`
	}

	ResetPasswordRequest struct {
		NewPassword string `json:"new_password" binding:"required"`
	}

	ForgetPasswordRequest struct {
		Email string `json:"email"`
	}
)
