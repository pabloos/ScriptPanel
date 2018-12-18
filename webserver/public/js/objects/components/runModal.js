function runModal(script) {
    return {
        element: "div",
        id: "runModal",
        class: "modal fade",
        children: [
            {
                element: "div",
                class: "modal-dialog",
                children: [
                    {
                        element: "div",
                        class: "modal-content",
                        children: [
                            {
                                element: "div",
                                class: "modal-header",
                                children: [
                                    {
                                        element: "h5",
                                        class: "modal-title",
                                        innerHTML: script.filename
                                    },
                                    {
                                        element: "button",
                                        type: "button",
                                        class: "close",
                                        dataDismiss: "modal",
                                        children: [
                                            {
                                                element: "span",
                                                innerHTML: "&times;"
                                            }
                                        ]
                                    }
                                ]
                            },
                            {
                                element: "div",
                                class: "modal-body",
                                children: [
                                    {
                                        element: "form",
                                        class: "px-4 py-3",
                                        id: "flagsForm",
                                        children: [
                                            {
                                                element: "div",
                                                class: "form-group",
                                                children: [
                                                    {
                                                        element: "fake",
                                                        innerHTML: "flag"
                                                    },
                                                    {
                                                        element: "select",
                                                        class: "form-control",
                                                        id: "flagsSelect"
                                                    }
                                                ]
                                            }
                                        ]
                                    },
                                    {
                                        element: "form",
                                        class: "px-4 py-3",
                                        id: "argsForm",
                                        children: [
                                            
                                        ]
                                    }
                                ]
                            },{
                                element: "div",
                                class: "modal-footer",
                                children: [
                                    {
                                        element: "button",
                                        class: "btn btn-primary",
                                        id: "runButton",
                                        innerHTML: "Run!",
                                        script: script,
                                        onclick: runScript

                                    }
                                ]
                            }
                        ]
                    }
                ]
            }
        ]
    };
}