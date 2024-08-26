package web_controllers

import (
	"errors"
	"fiber_blog/app/models"
	"fiber_blog/env"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jhonoryza/inertia-fiber"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginForm() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return inertia.Render(c, http.StatusOK, "Auth/Login", fiber.Map{
			"canResetPassword": true,
		})
	}
}

func registerValidate() (*validator.Validate, ut.Translator) {
	validate := validator.New()

	// Mengatur penerjemah (translator) untuk pesan error kustom
	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	trans, _ := uni.GetTranslator("en")

	// Mendaftarkan terjemahan default bahasa Inggris
	_ = en_translations.RegisterDefaultTranslations(validate, trans)

	// Menambahkan pesan error kustom untuk tag 'required'
	_ = validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is required", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	// Menambahkan pesan error kustom untuk tag 'email'
	_ = validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} must be a valid email address!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})

	// Menambahkan pesan error kustom untuk tag 'min'
	_ = validate.RegisterTranslation("min", trans, func(ut ut.Translator) error {
		return ut.Add("min", "{0} minimum must be {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", fe.Field(), fe.Param())
		return t
	})

	// Menambahkan pesan error kustom untuk tag 'max'
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
		validationErrors := err.(validator.ValidationErrors)
		var errMessage string
		for _, fieldError := range validationErrors {
			errMessage = errMessage + fmt.Sprintf("%v,", fieldError.Translate(trans))
		}
		return nil, errors.New(errMessage)
	}
	return &loginData, nil
}

func Login(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		loginData, err := validateLoginData(c)
		if err != nil {
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
			return inertia.Render(c, http.StatusBadRequest, "Auth/Login", fiber.Map{
				"errors": map[string]string{
					"email": "invalid credential",
				},
			})
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
		if err != nil {
			return inertia.Render(c, http.StatusBadRequest, "Auth/Login", fiber.Map{
				"errors": map[string]string{
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
			return inertia.Render(c, http.StatusInternalServerError, "Auth/Login", fiber.Map{
				"message": "internal server error",
			})
		}

		// Set cookie dengan token
		c.Cookie(&fiber.Cookie{
			Name:     env.GetEnv().GetString("COOKIE_NAME"),
			Value:    tokenString,
			Expires:  timeExpire,
			HTTPOnly: true,
			Secure:   true,
			SameSite: "Strict",
		})

		return c.RedirectToRoute("dashboard", fiber.Map{})
	}
}

func Logout() fiber.Handler {
	return func(c *fiber.Ctx) error {

		// Menghapus cookie dengan mengatur waktu kedaluwarsa di masa lalu
		c.Cookie(&fiber.Cookie{
			Name:     env.GetEnv().GetString("COOKIE_NAME"),
			Value:    "",
			Expires:  time.Now().Add(-time.Hour), // Set waktu kedaluwarsa di masa lalu
			HTTPOnly: true,
			Secure:   true,
			SameSite: "Strict",
		})

		return c.RedirectToRoute("login.form", fiber.Map{})
	}
}
