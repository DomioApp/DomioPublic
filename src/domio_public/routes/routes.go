package routes

import (
    "net/http"
    "domio_public/handlers/home_page_handler"
    "domio_public/handlers/login_page_handler"
    "domio_public/handlers/profile_page_handler"
    "domio_public/handlers/user_domains_page_handler"
    "domio_public/handlers/add_domain_page_handler"
    "domio_public/handlers/signup_page_handler"
    "domio_public/handlers/logout_page_handler"
    "domio_public/handlers/domains_page_handler"
    "domio_public/handlers/domain_info_page_handler"
    "domio_public/handlers/domain_edit_page_handler"
    "domio_public/handlers/rent_domain_page_handler"
    "domio_public/handlers/add_payment_source_page_handler"
    "domio_public/handlers/user_payment_sources_page_handler"
    "domio_public/handlers/user_payment_source_page_handler"
    "domio_public/handlers/subscriptions_page_handler"
    "domio_public/handlers/subscription_page_handler"
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
        "DomainsPage",
        "GET",
        domains_page_handler.GetUrl(),
        domains_page_handler.DomainsPageHandler,
    },
    Route{
        "DomainInfoPage",
        "GET",
        domain_info_page_handler.GetUrl(),
        domain_info_page_handler.DomainInfoPageHandler,
    },
    Route{
        "RentDomainPage",
        "GET",
        rent_domain_page_handler.GetUrl(),
        rent_domain_page_handler.RentDomainPageHandler,
    },
    Route{
        "DomainEditPage",
        "GET",
        domain_edit_page_handler.GetUrl(),
        domain_edit_page_handler.DomainEditPageHandler,
    },
    Route{
        "LoginPage",
        "GET",
        login_page_handler.GetUrl(),
        login_page_handler.LoginPageHandler,
    },
    Route{
        "LogoutPage",
        "GET",
        logout_page_handler.GetUrl(),
        logout_page_handler.LogoutPageHandler,
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
        "AddPaymentSourcePage",
        "GET",
        add_payment_source_page_handler.GetUrl(),
        add_payment_source_page_handler.AddPaymentSourcePageHandler,
    },
    Route{
        "UserPaymentSourcesPage",
        "GET",
        user_payment_sources_page_handler.GetUrl(),
        user_payment_sources_page_handler.UserPaymentSourcesPageHandler,
    },
    Route{
        "UserPaymentSourcePage",
        "GET",
        user_payment_source_page_handler.GetUrl(),
        user_payment_source_page_handler.UserPaymentSourcePageHandler,
    },
    Route{
        "UserDomainsPage",
        "GET",
        user_domains_page_handler.GetUrl(),
        user_domains_page_handler.UserDomainsPageHandler,
    },
    Route{
        "SubscriptionsPage",
        "GET",
        subscriptions_page_handler.GetUrl(),
        subscriptions_page_handler.SubscriptionsPageHandler,
    },
    Route{
        "SubscriptionPage",
        "GET",
        subscription_page_handler.GetUrl(),
        subscription_page_handler.SubscriptionPageHandler,
    },
    Route{
        "AddDomainPage",
        "GET",
        add_domain_page_handler.GetUrl(),
        add_domain_page_handler.AddDomainPageHandler,
    },
}

