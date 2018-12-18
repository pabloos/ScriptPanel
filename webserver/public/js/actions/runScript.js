function showRunModal() {
    $('#runModal').modal('show');
}

function createRunModal(script) {    
    if($("#runModal").length) {     //remove if there's a runModal configured
        $("#runModal").remove();
    }

    const modal = createComponent(runModal(script));

    document.getElementsByTagName("body")[0].appendChild(modal);

    const select = document.getElementById("flagsSelect");

    for(let flag in script.flags) {
        const option = createComponent({
            element: "option",
            innerHTML: script.flags[flag]
        });

        select.appendChild(option);
    }

    for(let arg in script.args) { 
        const argsFormGroup = createComponent({
            element: "div",
            class: "form-check form-check-inline",
            children: [
                {
                    element: "input",
                    class: "form-check-input",
                    type: "checkbox",
                    value: script.args[arg],
                    id: script.args[arg],
                },
                {
                    element: "label",
                    class: "form-check-label",
                    for: script.args[arg],
                    innerHTML: script.args[arg]
                }  
            ]
        });
        
        argsForm.appendChild(argsFormGroup);
    }
}

function runScript() {    
    let config = {};

    config.flag = getFormObject(flagsForm).flag;    //take the args and flags from the form modal
    config.args = getArgs(argsForm);
    
    runRequest = {};
    runRequest.script = this.script;
    runRequest.config = config;

    const init = {
        method: 'POST',
        body: JSON.stringify(runRequest),
        headers:{
            'Content-Type': 'application/json'
        }
    }

    var request = new Request("/runScript", init)

    fetch(request)
        .then(response => response = response.json())
        .catch(error => console.error('Error:', error))
        .then(response => {
            if(false) { //login fails case
                alert("RunScript has failed");
            } else {
                processResult(response)
            }
        })
}

function getArgs(form) {
    let array = Array.from(form.children);
    let args = [];

    for (let i in array) {
        input = array[i].children[0];
        
        if(input.checked){
            args.push(input.value);
        }
    }

    return args;
}

function processResult(data){    
    $('#runModal').modal('hide');

    $("#resultBox").remove();       //we need to remove the previous resultBox
    
    scriptPanel.appendChild(createComponent(resultBox(data)));
}