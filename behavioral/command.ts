
function add(a,b) {return a+b}
function sub(a,b) {return a-b}
function div(a,b) {return a/b}
function mul(a,b) {return a*b}

let Command = function (execute, undo,value) {
    this.execute = execute;
    this.undo = undo;
    this.value = value;
}

function AddCommand(value) {
    return new Command(add,sub,value);
}
function SubCommand(value) {
    return new Command(sub,add,value);
}
function DivCommand(value) {
    return new Command(div,mul,value);
}
function MulCommand(value) {
    return new Command(mul,div,value);
}

class Calculator {
    protected current:number =0;
    protected commands = new Array();
    private action(command) {
        var name = command.execute.toString().substr(9, 3);
        return name.charAt(0).toUpperCase() + name.slice(1);
    }
    execute(command){
        this.current = command.execute(this.current,command.value);
        this.commands.push(command);
        console.log(this.action(command) + ": " + command.value);
    }
    undo(){
        let command = this.commands.pop();
        this.current = command.undo(this.current,command.value);
        console.log("Undo :"+ this.action(command) + ": " + command.value);
    }
    getCurrentValue(){
        return this.current;
    }
}


/** Client code **/
var calculator = new Calculator();

// issue commands

calculator.execute(new AddCommand(100));
calculator.execute(new SubCommand(24));
calculator.execute(new MulCommand(6));
calculator.execute(new DivCommand(2));

// reverse last two commands

calculator.undo();
calculator.undo();

console.log("\nValue: " + calculator.getCurrentValue());

