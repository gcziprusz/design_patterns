interface Handler {
    setNext(handler:Handler):Handler;
    handle(request:string):string;
}
abstract class AbstractHandler implements Handler {
    private nextHandler: Handler|null = null;
    public setNext(handler:Handler):Handler{
        this.nextHandler=handler;
        return this.nextHandler;
    }
    public handle(request:string):string{
        if(!this.nextHandler) return "";
        return this.nextHandler.handle(request)
    }
}

/******** Concrete Handlers **********/
class MonkeyHandler extends AbstractHandler {
    handle(request: string):string {
        if(request === 'banana') return `\tMonkey: I will handle the ${request}!`
        return super.handle(request);
    }
}
class CatHandler extends AbstractHandler {
    handle(request: string):string {
        if(request === 'fish') return `\tCat: I will handle the ${request}!`
        return super.handle(request);
    }
}
class ChipmunkHandler extends AbstractHandler {
    handle(request: string):string {
        if(request === 'acorn') return `\tChipmunk: I will handle the ${request}!`
        return super.handle(request);
    }
}

/******** CLIENT CODE **********/
function clientCode(handler:Handler)
{
    const foods = ['banana','acorn','tea','fish'];
    for(let f of foods){
        console.log(`\nWho wants to eat the ${f}?`)
        let res = handler.handle(f);
        if(!res) 
        {
            console.log(`\t${f} was not handled by anyone!`);
        } else {
            console.log(`\t${res}`)
        }

    }

}


const monkey = new MonkeyHandler()
const cat = new CatHandler()
const chipmunk = new ChipmunkHandler()

monkey.setNext(cat).setNext(chipmunk);

console.log('\n\nChain: Monkey > Cat > Chipmunk\n');
clientCode(monkey)

console.log('\n\nChain: Cat > Chipmunk\n');
clientCode(cat)

console.log('\n\nChain: Chipmunk\n');
clientCode(chipmunk)
