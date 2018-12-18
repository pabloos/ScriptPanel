function userButton(user){
    return {
        element: "button",
        id: "userButton",
        class: "btn btn-success dropdown-toggle", //convertimos el botón de login en un dropdown
        dataTarget: "",                           //para que no nos despliegue el formulario
        dataToggle: "dropdown",
        innerHTML: user
    }
}