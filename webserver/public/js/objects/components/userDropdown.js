function userDropdown() {
    return {
        element: "div",
        class: "dropdown-menu dropdown-menu-right",
        children: [
            {
                element: "a",
                class: "dropdown-item",
                dataToggle: "modal",
                dataTarget: "#uploadScriptModal",
                innerHTML: "Upload a script"
            },
            {
                element: "div",
                class: "dropdown-divider"
            },
            {
                element: "a",
                class: "dropdown-item",
                innerHTML: "Log out",
                onclick: unset
            }
        ]
    }
}