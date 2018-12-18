function sendAndThen(method, object, path, action) {
    sendAjax(method, object, path).then(data => action(data));
}

function sendAjax(method, obj, url) {
    return new Promise((resolve, reject) => {
        $.ajax({
            type: method,

            //aniadido por mi
            contentType: false,
            processData: false,

            url: url,
            data: obj,
            
            success: function (data) {
                resolve(data);
            }
        });
    });
}
