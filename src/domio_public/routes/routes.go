package routes

import (
    "net/http"
    "domio_public/handlers/home_page_handler"
    "domio_public/handlers/login_page_handler"
    "domio_public/handlers/profile_page_handler"
    "domio_public/handlers/user_domains_page_handler"
    "domio_public/handlers/add_domain_page_handler"
    "domio_public/handlers/signup_page_handler"
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
        "HomePage",
        "GET",
        home_page_handler.GetUrl(),
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
    Route{
        "ProfilePage",
        "GET",
        profile_page_handler.GetUrl(),
        profile_page_handler.ProfilePageHandler,
    },
    Route{
        "UserDomainsPage",
        "GET",
        user_domains_page_handler.GetUrl(),
        user_domains_page_handler.UserDomainsPageHandler,
    },
    Route{
        "AddDomainPage",
        "GET",
        add_domain_page_handler.GetUrl(),
        add_domain_page_handler.AddDomainPageHandler,
    },
}

