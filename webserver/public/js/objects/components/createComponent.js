// function createComponent(object) {
//     let component = document.createElement(object.element);

//     if(object.hasOwnProperty("id")){
//         component.setAttribute("id", object.id);
//     }

//     if(object.hasOwnProperty("class")){
//         component.setAttribute("class" , object.class);
//     }

//     if(object.hasOwnProperty("children")){
//         for (var index in object.children) {
//             component.appendChild(createComponent(object.children[index]));
//         }
//     }

//     if(object.hasOwnProperty("innerHTML")){
//         component.innerHTML = object.innerHTML;
//     }

//     if(object.hasOwnProperty("src")){
//         component.setAttribute("src", object.src);
//     }

//     if(object.hasOwnProperty("type")){
//         component.setAttribute("type", object.type);
//     }

//     if(object.hasOwnProperty("ariaLabel")){
//         component.setAttribute("aria-label", object.ariaLabel);
//     }

//     if(object.hasOwnProperty("onclick")){
//         component.addEventListener("click", object.onclick)
//     }

//     return component;
// }

function createComponent(object) {
    let component;

    const hash = {
        "element"   : () => {component = document.createElement(object.element)},
        
        "id"        : () => {component.setAttribute("id", object.id)},
        "class"     : () => {component.setAttribute("class" , object.class)},
        "src"       : () => {component.setAttribute("src", object.src)},
        "type"      : () => {component.setAttribute("type", object.type)},
        "for"       : () => {component.setAttribute("for", object.for)},
        "ariaLabel" : () => {component.setAttribute("ariaLabel", object.ariaLabel)},

        "innerHTML" : () => {component.innerHTML = object.innerHTML},

        "value"     : () => {component.setAttribute("value", object.value)},

        "dataToggle": () => {component.setAttribute("data-toggle", object.dataToggle)},
        "dataTarget": () => {component.setAttribute("data-target", object.dataTarget)},

        "dataDismiss": () => {component.setAttribute("data-dismiss", object.dataDismiss)},

        "children"  : () => {
            for (let index in object.children) {
                component.appendChild(createComponent(object.children[index]));
            }
        },

        "script" : () => {component.script =  object.script},

        "onclick"   : () => {component.addEventListener("click", object.onclick)}
    }

    const properties = Object.keys(object);

    for(let key in properties) {
        hash[properties[key]]();
    }

    return component;
}