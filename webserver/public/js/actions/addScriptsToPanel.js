//this one has the mission to put all the scripts in the html
//is the general js script which goes on the window.onload

//the main page needs to be saved in order to be restablised when the user logs out
// presentation = $("#presentationContainer").html();
// var presen = document.createTextNode(presentation);

async function removePresentation() {
    presentationContainer.setAttribute("class", "container-fluid animated slideOutLeft");

    (async function(){
        await sleep(3000);
        presentationContainer.setAttribute("class", "container-fluid");
        presentationContainer.style.display = "none";
        presentationContainer.remove();
    }())
}

function addScriptsToPanel(scripts) {
    userScripts = scripts;

    createPanel(scripts);
}

async function createPanel(scripts){
    const numberOfScripts   = Object.keys(scripts).length;
    const numberOfRows      = Math.ceil((numberOfScripts / 4));

    const iterator = makeIterator(scripts);

    const scriptPanel = createComponent(ScriptPanel());
    
    for(let i = 0; i < numberOfRows; i++) {
        const row = createComponent({
            element: "div",
            class: "row"
        });

        for(let j = 0; j < NUM_COLS; j++) {           
            const col = createComponent({
                element: "div",
                class: "col-sm"
            });

            if(iterator.hasNext()){
                const script = iterator.next();
                const card = createComponent(Card(script));

                col.appendChild(card);
            }
            row.appendChild(col);
        }
        scriptPanel.appendChild(row);
    }

    scriptPanel.appendChild(document.createElement("br"));
    main.appendChild(scriptPanel);
}