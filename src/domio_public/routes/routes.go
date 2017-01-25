package routes

import (
    "net/http"
    "domio_public/handlers/signup_page_handler"
    "domio_public/handlers/home_page_handler"
    "domio_public/handlers/login_page_handler"
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
        home_page_handler.HomePageHandler,
    },
    Route{
        "LoginPage",
        "GET",
        login_page_handler.GetUrl(),
        login_page_handler.LoginPageHandler,
    },
    Route{
        "SignupPage",
        "GET",
        signup_page_handler.GetUrl(),
        signup_page_handler.SignupPageHandler,
    },
}