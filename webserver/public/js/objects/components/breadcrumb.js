function Breadcrumb() {
    return {         //addBreadcumb
        element: "nav",
        ariaLabel: "breadcrumb",
        children: [
            {
                element: "ol",
                class: "breadcrumb animated bounceInUp",
                children: [
                    {
                        element: "li",
                        class: "breadcrumb-item",
                        innerHTML: loginCookies.get("company")
                    },
                    {
                        element: "li",
                        class: "breadcrumb-item",
                        innerHTML: loginCookies.get("department")
                    },
                    {
                        element: "li",
                        class: "breadcrumb-item",
                        innerHTML: loginCookies.get("name")
                    },
                ]
            }
        ]
    }
}