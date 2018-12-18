function scriptSearch(scripts, name) {
    for (var script in scripts ) {
        if (scripts[script].filename == name){
            return scripts[script];
        }
    }
}