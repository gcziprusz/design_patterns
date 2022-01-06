interface State {
    do1():void;
    do2():void;
}

class PendingState implements State{
    do1(){
        console.log(`PendingState doing one`)
    }
    do2(){
        console.log(`PendingState doing two`)
    }
}
class InProgressState implements State{
    do1(){
        console.log(`InProgressState doing one`)
    }
    do2(){
        console.log(`InProgressState doing two`)
    }
}
class DoneState implements State{
    do1(){
        console.log(`DoneState done one`)
    }
    do2(){
        console.log(`DoneState done two`)
    }
}


class Thing {
    state: State;
    constructor(s: State){
        this.state =s;
    }
    transition(s:State) {
        this.state =s;
    }
    request1(){
        this.state.do1();
    }
    request2(){
        this.state.do2();
    }
}

let t = new Thing(new PendingState());

t.request1()
t.request2()
console.log('--------------------------------')
t.transition(new InProgressState())
t.request1()
t.request2()
console.log('--------------------------------')
t.transition(new DoneState())
t.request1()
t.request2()






