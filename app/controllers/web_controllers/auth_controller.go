package web_controllers

import (
	"errors"
	"fiber_blog/app/models"
	"fiber_blog/env"
	"fiber_blog/utils"
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/session"
	"net/http"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jhonoryza/inertia-fiber"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginForm() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies(env.GetEnv().GetString("COOKIE_NAME"))
		if cookie != "" {
			return inertia.RedirectToRoute(c, "auth.dashboard", fiber.Map{})
		}

		return inertia.Render(c, http.StatusOK, "Auth/Login", fiber.Map{
			"canResetPassword": true,
		})
	}
}

func registerValidate() (*validator.Validate, ut.Translator) {
	validate := validator.New()

	// set translator for custom error message
	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	trans, _ := uni.GetTranslator("en")

	// register default translator using en
	_ = entranslations.RegisterDefaultTranslations(validate, trans)

	// add custom error for tag 'required'
	_ = validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is required", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	// add custom error for tag 'email'
	_ = validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} must be a valid email address!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})

	// add custom error for tag 'min'
	_ = validate.RegisterTranslation("min", trans, func(ut ut.Translator) error {
		return ut.Add("min", "{0} minimum must be {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", fe.Field(), fe.Param())
		return t
	})

	// add custom error for tag 'max'
	_ = validate.RegisterTranslation("max", trans, func(ut ut.Translator) error {
		return ut.Add("max", "{0} maximum must be {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("max", fe.Field(), fe.Param())
		return t
	})
	return validate, trans
}

type LoginData struct {
	Email    string `json:"email" validate:"required,email,min=1,max=255"`
	Password string `json:"password" validate:"required,min=5,max=255"`
	Remember bool   `json:"remember"`
}

func validateLoginData(c *fiber.Ctx) (*LoginData, error) {
	var loginData LoginData
	err := c.BodyParser(&loginData)
	if err != nil {
		return nil, errors.New("bad request")
	}

	validate, trans := registerValidate()
	err = validate.Struct(loginData)
	if err != nil {
		var validationErrors validator.ValidationErrors
		errors.As(err, &validationErrors)
		var errMessage string
		for _, fieldError := range validationErrors {
			errMessage = errMessage + fmt.Sprintf("%v,", fieldError.Translate(trans))
		}
		return nil, errors.New(errMessage)
	}
	return &loginData, nil
}

func Login(db *gorm.DB, store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		loginData, err := validateLoginData(c)
		if err != nil {
			inertia.Share(c, fiber.Map{
				"flash": utils.FlashMessage{
					Message: err.Error(),
					Type:    "danger",
				},
			})

			return inertia.Render(c, http.StatusBadRequest, "Auth/Login", fiber.Map{
				"message": err.Error(),
			})
		}

		email := loginData.Email
		pass := loginData.Password
		isRemember := loginData.Remember

		timeExpire := time.Now().Add(time.Hour * 1)
		unixExpire := time.Now().Add(time.Hour * 1).Unix()
		if isRemember {
			timeExpire = time.Now().Add(time.Hour * 2)
			unixExpire = time.Now().Add(time.Hour * 2).Unix()
		}

		// Validate email and password
		var user models.User
		tx := db.Where("email = ?", email).First(&user)
		if tx.Error != nil {
			inertia.Share(c, fiber.Map{
				"flash": utils.FlashMessage{
					Message: "Invalid Credential.",
					Type:    "danger",
				},
			})

			return inertia.Render(c, http.StatusBadRequest, "Auth/Login", fiber.Map{
				"errors": map[string]string{
					"email":    "invalid credential",
					"password": "invalid credential",
				},
			})
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
		if err != nil {
			inertia.Share(c, fiber.Map{
				"flash": utils.FlashMessage{
					Message: "Invalid Credential.",
					Type:    "danger",
				},
			})

			return inertia.Render(c, http.StatusBadRequest, "Auth/Login", fiber.Map{
				"errors": map[string]string{
					"email":    "invalid credential",
					"password": "invalid credential",
				},
			})
		}

		// Create the Claims
		claims := jwt.MapClaims{
			"name":  user.Name,
			"email": user.Email,
			"exp":   unixExpire,
		}

		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		tokenString, err := token.SignedString([]byte(env.GetEnv().GetString("JWT_SECRET")))
		if err != nil {
			inertia.Share(c, fiber.Map{
				"flash": utils.FlashMessage{
					Message: "internal server error",
					Type:    "danger",
				},
			})

			return inertia.Render(c, http.StatusInternalServerError, "Auth/Login", fiber.Map{
				"message": "internal server error",
			})
		}

		// Set cookie using jwt token
		c.Cookie(&fiber.Cookie{
			Name:     env.GetEnv().GetString("COOKIE_NAME"),
			Value:    tokenString,
			Expires:  timeExpire,
			HTTPOnly: true,
			Secure:   true,
			SameSite: "Strict",
		})

		utils.SessionFlash(store, c, fiber.Map{
			"message": "login success",
			"type":    "success",
		})

		return inertia.RedirectToRoute(c, "auth.dashboard", fiber.Map{})
	}
}

func Logout(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {

		// delete cookie with expired time
		c.Cookie(&fiber.Cookie{
			Name:     env.GetEnv().GetString("COOKIE_NAME"),
			Value:    "",
			Expires:  time.Now().Add(-time.Hour), // Set expired time
			HTTPOnly: true,
			Secure:   true,
			SameSite: "Strict",
		})

		utils.SessionFlash(store, c, fiber.Map{
			"message": "logout success",
			"type":    "success",
		})

		return inertia.RedirectToRoute(c, "login", fiber.Map{})
	}
}
