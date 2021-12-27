interface Mediator {
    notify(sender: Component, event : string):void
}

class LoginFormMediator implements Mediator {
    protected formFields: Component[]
    constructor(...components: Component[]){
        this.formFields=components;
        this.formFields.forEach(f => f.setMediator(this))
    }
    notify(sender: Component, event : "click"|"hover"|"input") {
        switch (event) {
            case "click": {
                console.log(`\tExecute ${event} actions of Mediator sent by ${sender.getName()}\n`);
                console.log(`\tValidate all formFields ${event} actions of Mediator sent by ${sender.getName()}\n`);
                this.formFields.forEach(ff => ff.validate())
                break;
            }
            case "hover": {
                console.log(`\tExecute ${event} actions of Mediator sent by ${sender.getName()}\n`);
                break;
            }
            case "input": {
                console.log(`\tExecute ${event} actions of Mediator sent by ${sender.getName()}\n`);
                break;
            }
        }
    }
}

abstract class Component {
    protected mediator?: Mediator
    protected name: string
    constructor(name:string){
        this.name=name;
    }
    setMediator(mediator: Mediator) {
        this.mediator = mediator
    }
    getName() {return this.name}
    validate() {
        console.log(`\tComponent ${this.getName()} validating ... `)
        return true;
    }
}

class Button extends Component {
    onClick(){
        console.log(`Execute onClick actions of ${this.getName()} Button`);
        this.mediator?.notify(this,'click');
    }
    onHover(){
        console.log(`Execute onHover actions of ${this.getName()} Button`);
        this.mediator?.notify(this,"hover");
    }

} 
class TextField extends Component {
    onInput(){
        console.log(`Execute onInput actions of TextField ${this.getName()} Component`);
        this.mediator?.notify(this,"input");
    }
} 

let button = new Button("Submit",)
let emailField = new TextField("Email")
let fnameField = new TextField("First Name")

let mediator = new LoginFormMediator(button,emailField,fnameField);

emailField.onInput()
fnameField.onInput()
fnameField.onInput()
button.onHover()
button.onClick()



