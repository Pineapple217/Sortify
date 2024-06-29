package handler

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/Pineapple217/Sortify/web/pkg/auth"
	"github.com/Pineapple217/Sortify/web/pkg/database"
	v "github.com/Pineapple217/Sortify/web/pkg/validate"
	"github.com/Pineapple217/Sortify/web/pkg/view"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
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
	// If user is logged in, redirect, TODO: should be a middleware
	if auth.GetAuth(c.Request().Context()).Check() {
		return c.Redirect(http.StatusSeeOther, "/")
	}
	return render(c, view.LoginIndex(view.LoginIndexPageData{}))
}

func (h *Handler) LoginUser(c echo.Context) error {
	// If user is logged in, redirect
	if auth.GetAuth(c.Request().Context()).Check() {
		return c.Redirect(http.StatusSeeOther, "/")
	}

	var values view.LoginFormValues
	errors, ok := v.Request(c.Request(), &values, authSchema)
	if !ok {
		return render(c, view.LoginForm(values, errors))
	}

	user, err := h.DB.GetUser(c.Request().Context(), values.Username)
	if err != nil {
		errors.Add("credentials", "invalid credentials")
		return render(c, view.LoginForm(values, errors))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(values.Password))
	if err != nil {
		errors.Add("credentials", "invalid credentials")
		return render(c, view.LoginForm(values, errors))
	}

	s, err := h.DB.CreateSession(c.Request().Context(), database.CreateSessionParams{
		UserID: user.ID,
		Token:  uuid.New().String(),
		ExpiresAt: pgtype.Timestamptz{
			Time:  time.Now().Add(time.Hour * time.Duration(sessionExpiry)),
			Valid: true,
		},
		IpAddress: pgtype.Text{String: c.RealIP(), Valid: true},
		UserAgent: pgtype.Text{String: c.Request().UserAgent(), Valid: true},
	})
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
	// If user is logged in, redirect
	if auth.GetAuth(c.Request().Context()).Check() {
		return c.Redirect(http.StatusSeeOther, "/")
	}
	return render(c, view.SignupIndex(view.SignupIndexPageData{}))
}

func (h *Handler) SignupUser(c echo.Context) error {
	// If user is logged in, redirect
	if auth.GetAuth(c.Request().Context()).Check() {
		return c.Redirect(http.StatusSeeOther, "/")
	}

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
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		// unique_violation
		if pgErr.Code == "23505" {
			form_errors.Add("username", "username is not available")
			return render(c, view.SignupForm(values, form_errors))
		}
		return err
	}

	s, err := h.DB.CreateSession(c.Request().Context(), database.CreateSessionParams{
		UserID: user.ID,
		Token:  uuid.New().String(),
		ExpiresAt: pgtype.Timestamptz{
			Time:  time.Now().Add(time.Hour * time.Duration(sessionExpiry)),
			Valid: true,
		},
		IpAddress: pgtype.Text{String: c.RealIP(), Valid: true},
		UserAgent: pgtype.Text{String: c.Request().UserAgent(), Valid: true},
	})
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
	// If user isn't logged in, redirect
	if !auth.GetAuth(c.Request().Context()).Check() {
		return c.Redirect(http.StatusSeeOther, "/")
	}

	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}

	// This should happen even if the db delete fails
	defer func() {
		sess.Values = map[any]any{}
		sess.Save(c.Request(), c.Response())
	}()

	err = h.DB.DeleteSessionByToken(c.Request().Context(), sess.Values["sessionToken"].(string))
	if err != nil {
		return err
	}

	c.Response().Header().Set("HX-Refresh", "true")
	return c.NoContent(http.StatusOK)
}

func createUserFromFormValues(ctx context.Context, db *database.Queries, values view.SignupFormValues) (database.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(values.Password), bcrypt.DefaultCost)
	if err != nil {
		return database.User{}, err
	}

	user, err := db.CreateUser(ctx, database.CreateUserParams{
		Username:     values.Username,
		PasswordHash: string(hash),
	})
	if err != nil {
		return user, err
	}
	return user, nil
}
