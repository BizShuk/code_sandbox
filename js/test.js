var non_flatten = [[1,2,[3]],4];
var test = 1;

function flatten(t){
    if (typeof t.length == 'undefined') {
        return []
    }

    flatten_array = [];
    for (var i=0;i<t.length;i++){
        if (typeof t[i].length != "undefined" ) {   // the elem is Array
            flatten_array = flatten_array.concat(flatten(t[i]));
        }else{                                      // the elem is Number
            flatten_array.push(t[i]);
        }
    }
    return flatten_array
}

flatten = flatten(non_flatten);
console.log("non flatten",non_flatten)
console.log("flatten:",flatten)





function IsPrime(n){
    if (typeof n != "number" || n % 1 != 0) {
        return false;
    }
    for (var i = 2, len = n; i < len; i++) {
        if (n%i===0) {
            return false;
        }
    }
    return true;
}


console.log("IsPrime:");
console.log("30 is Prime ?",IsPrime(30));
console.log("7 is Prime ?",IsPrime(7));
console.log("aaa is Prime ?",IsPrime("aaa"));



// TODO: Primise version
function DescendOfThree(n) {
    if (typeof n != "number" || n % 1 != 0) {
        return false;
    }

    for (var i = n, len = 3; i >= len; i--) {
        console.log(i,n);
        if (i%3==0) {
            tmp = i;
            do {
                console.log(tmp);
                tmp -= 3;
            } while (tmp%3==0 && tmp>0);
            break;
        }
    }
    return 
}

DescendOfThree(200);




