function resultBox(data) {
    return {
        element: "div",
        id: "resultBox",
        class: "alert alert-success animated bounceInRight",
        children: [
            {
                element: "h4",
                innerHTML: "Result is: "
            },
            {
                element: "p",
                innerHTML: "> " + data
            }
        ]
    }
}