package routes

import (
    "net/http"
    "domio_public/handlers/login_page_handler"
    "domio_public/handlers/signup_page_handler"
    "domio_public/handlers/home_handler"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var RoutesList = Routes{
    Route{
        "Index",
        "GET",
        "/",
        home_handler.HomeHandler,
    },
    Route{
        "LoginPage",
        "GET",
        "/login",
        login_page_handler.LoginPageHandler,
    },
    Route{
        "SignupPage",
        "GET",
        "/signup",
        signup_page_handler.SignupPageHandler,
    },
}