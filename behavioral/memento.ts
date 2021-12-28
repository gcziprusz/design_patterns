class Originator {
    private state:string;
    constructor(state:string){
        this.state=state;
    }
    createMemento(){
     return new Memento(this.state);
    }
    restoreMemento(m: Memento){
        this.state = m.get();
    }
    setState(state: string){
        this.state=state;
    }
    getState(){
        return this.state;
    }
}


class Memento {
    private state: string;
    constructor(state:string){
        this.state =state;
    }
    get(){
        return this.state;
    }
}

class CareTaker {
    private mementos = new Array<Memento>();
    private originator: Originator;
    constructor( originator: Originator){
        this.originator = originator;
    }
    save(m: Memento){
        this.mementos.push(m)
    }
    undo(){
        if (this.mementos.length < 1 ) return;
        this.originator.restoreMemento(this.mementos.pop() as Memento);
    }
}


let o = new Originator("A");
console.log(`Originator STATE ${o.getState()}`);
let careTaker = new CareTaker(o);


careTaker.save(o.createMemento());
o.setState("B");
console.log(`Originator STATE ${o.getState()}`);


careTaker.save(o.createMemento());
o.setState("C")
console.log(`Originator STATE ${o.getState()}`);


console.log(`undo ${o.getState()}`);
careTaker.undo();
console.log(`Originator STATE ${o.getState()}`);



console.log(`undo ${o.getState()}`);
careTaker.undo();
console.log(`Originator STATE ${o.getState()}`);
careTaker.undo();
console.log(`Originator STATE ${o.getState()}`);
careTaker.undo();
console.log(`Originator STATE ${o.getState()}`);


