package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/Pineapple217/Sortify/web/ent"
	DBSession "github.com/Pineapple217/Sortify/web/ent/session"
	"github.com/Pineapple217/Sortify/web/ent/user"
	v "github.com/Pineapple217/Sortify/web/pkg/validate"
	"github.com/Pineapple217/Sortify/web/pkg/view"
	"github.com/google/uuid"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

const sessionExpiry = 48

var authSchema = v.Schema{
	"username": v.Rules(v.Required),
	"password": v.Rules(v.Required),
}

func (h *Handler) LoginIndex(c echo.Context) error {
	return render(c, view.LoginIndex(view.LoginIndexPageData{}))
}

func (h *Handler) LoginUser(c echo.Context) error {

	var values view.LoginFormValues
	errors, ok := v.Request(c.Request(), &values, authSchema)
	if !ok {
		return render(c, view.LoginForm(values, errors))
	}
	user, err := h.DB.User.Query().
		Where(user.Username(values.Username)).
		Only(c.Request().Context())
	if err != nil {
		errors.Add("credentials", "invalid credentials")
		return render(c, view.LoginForm(values, errors))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(values.Password))
	if err != nil {
		errors.Add("credentials", "invalid credentials")
		return render(c, view.LoginForm(values, errors))
	}

	s, err := h.DB.Session.Create().
		SetUserID(user.ID).
		SetToken(uuid.New().String()).
		SetExpiresAt(time.Now().Add(time.Hour * time.Duration(sessionExpiry))).
		SetIPAddress(c.RealIP()).
		SetUserAgent(c.Request().UserAgent()).
		Save(c.Request().Context())
	if err != nil {
		return err
	}

	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	sess.Values["sessionToken"] = s.Token
	err = sess.Save(c.Request(), c.Response().Writer)
	if err != nil {
		return err
	}

	c.Response().Header().Set("HX-Redirect", "/")
	return c.NoContent(http.StatusOK)
}

var signupSchema = v.Schema{
	"password": v.Rules(
		v.ContainsSpecial,
		v.ContainsUpper,
		v.Min(7),
		v.Max(50),
	),
	"username": v.Rules(v.Min(3), v.Max(50)),
}

func (h *Handler) SignupForm(c echo.Context) error {
	return render(c, view.SignupIndex(view.SignupIndexPageData{}))
}

func (h *Handler) SignupUser(c echo.Context) error {
	var values view.SignupFormValues
	form_errors, ok := v.Request(c.Request(), &values, signupSchema)
	if !ok {
		return render(c, view.SignupForm(values, form_errors))
	}

	if values.Password != values.PasswordConfirm {
		form_errors.Add("password", "passwords do not match")
		return render(c, view.SignupForm(values, form_errors))
	}
	user, err := createUserFromFormValues(c.Request().Context(), h.DB, values)
	if ent.IsConstraintError(err) {
		form_errors.Add("username", "username is not available")
		return render(c, view.SignupForm(values, form_errors))
	} else if err != nil {
		panic(err)
	}

	s, err := h.DB.Session.Create().
		SetUserID(user.ID).
		SetToken(uuid.New().String()).
		SetExpiresAt(time.Now().Add(time.Hour * time.Duration(sessionExpiry))).
		SetIPAddress(c.RealIP()).
		SetUserAgent(c.Request().UserAgent()).
		Save(c.Request().Context())
	if err != nil {
		return err
	}

	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	sess.Values["sessionToken"] = s.Token
	err = sess.Save(c.Request(), c.Response().Writer)
	if err != nil {
		return err
	}

	c.Response().Header().Set("HX-Redirect", "/")
	return c.NoContent(http.StatusOK)
}

func (h *Handler) LogoutUser(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}

	// This should happen even if the db delete fails
	defer func() {
		sess.Values = map[any]any{}
		sess.Save(c.Request(), c.Response())
	}()

	_, err = h.DB.Session.Delete().
		Where(DBSession.Token(sess.Values["sessionToken"].(string))).
		Exec(c.Request().Context())
	if err != nil {
		return err
	}

	c.Response().Header().Set("HX-Refresh", "true")
	return c.NoContent(http.StatusOK)
}

func createUserFromFormValues(ctx context.Context, db *ent.Client, values view.SignupFormValues) (*ent.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(values.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user, err := db.User.Create().
		SetUsername(values.Username).
		SetPasswordHash(string(hash)).
		Save(ctx)

	if err != nil {
		return user, err
	}
	return user, nil
}
