function fn() {

    console.log(Array.from(arguments));
}


a = [1, 2, 3];
fn(...a);
