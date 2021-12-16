interface Missile {
    fire():void;
}

// SUBJECT
class Rocket implements Missile{
    public fire(){
        console.log('A rocket fired!')
    }
}

// PROXY 
class RocketProxy implements Missile{
    private rocket: Rocket
    constructor(rocket: Rocket){
        this.rocket = rocket
    }
    public fire(){
        this.checkAccess()
        this.rocket.fire()
        this.logAccess()
    }
    private checkAccess():boolean {
        console.log('Checking access to Rocket')
        return true;
    }
    private logAccess():void{
        console.log('Logging access to Rocket')
    }
}


// CLIENT 
function clientCode(missile: Missile){
    missile.fire();
}

console.log('\nDirect Fire\n');
let r = new Rocket();
clientCode(r)

console.log('\nProxy Fire\n');
let auditedRocket = new RocketProxy(r);
clientCode(auditedRocket)
