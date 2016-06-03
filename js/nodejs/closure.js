function add() {
    var total = 0; 

    for (var i=0 ; i < arguments.length && i < 2 ; i++){
        if (typeof arguments[i] != "number") return undefined;
    }

    if ( arguments.length == 1 ){
        total = arguments[0];
        return function(num){
            return typeof num == "number" ? total+num : undefined;
        };
    }else{
        return arguments[0]+arguments[1]; 
    }
}

add(2,3);

