function Card(script){
    return {      //object card represents a script card on the panel
        element: "div",
        class: "card animated bounceIn",
        children: [
            {
                element: "img",
                class: "card-img-top",
                src: "./../../img/langs/" + langHash.get(script.language) + ".png" // dynamic item
            },
            {
                element: "div",
                class: "card-body",    
                children: [
                    {
                        element: 'h5',
                        class: "card-title",
                        children: [
                            {
                                element: 'p',
                                class: "card-text",
                                innerHTML: script.filename // dynamic item
                            }
                        ]
                    }
                ]
            },
            {
                element: "ul",
                class: "list-group list-group-flush",
                children: [
                    {
                        element: "li",
                        class: "list-group-item",
                        innerHTML: script.description
                    }
                ]
            },
            {
                element: "div",
                class: "card-body",
                children: [
                    {
                        element: "button",
                        type: "button",
                        class: "btn btn-outline-primary",
                        innerHTML: "Run!",
                        onclick: () => {
                            createRunModal(scriptSearch(userScripts, script.filename));
                            showRunModal();
                        }
                    }
                ]
            }
        ]
    }
}