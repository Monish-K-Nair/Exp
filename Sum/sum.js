process.stdin.resume();
process.stdin.setEncoding('ascii');

var input_stdin = "";
var input_stdin_array = "";
var input_currentline = 0;

process.stdin.on('data', function (data) {
    input_stdin += data;
});

process.stdin.on('end', function () {
    input_stdin_array = input_stdin.split("\n");
    main();    
});

function main() {
    let a = parseInt(input_stdin_array[input_currentline++]);
    let b = parseInt(input_stdin_array[input_currentline++]);
    console.log(a+b);
}