class Manager {
    public name:string
    protected salary:number
    protected vacation:number
    constructor(name:string,salary:number,vacation:number){
        this.name = name
        this.salary = salary
        this.vacation = vacation
    }
    public accept(v: Visitor){
        v.executeOnManager(this)
    }
    public getSalary() : number {
        return this.salary
    }
    public setSalary(v : number) {
        this.salary = v;
    }
    public getVacation() : number {
        return this.vacation
    }
    public setVacation(v : number) {
        this.vacation = v;
    }
}
class Employee {
    public name:string
    protected salary:number
    protected vacation:number
    constructor(name:string,salary:number,vacation:number){
        this.name = name
        this.salary = salary
        this.vacation = vacation
    }
    public accept(v: Visitor){
        v.executeOnEmployee(this)
    }
    public getSalary() : number {
        return this.salary
    }
    public setSalary(v : number) {
        this.salary = v;
    }
    public getVacation() : number {
        return this.vacation
    }
    public setVacation(v : number) {
        this.vacation = v;
    }
}


interface Visitor{
    executeOnEmployee(e:Employee):void
    executeOnManager(m:Manager):void
}

class SalaryGiver implements Visitor{
    public executeOnEmployee(e: Employee){
        e.setSalary(e.getSalary()*2)
        console.log(`Employee ${e.name} visited by SalaryGiver`)
    }
    public executeOnManager(m :Manager){
        m.setSalary(m.getSalary()*3)
        console.log(`Manager ${m.name} visited by SalaryGiver`)
    }
}
class VacationSucker implements Visitor{
    public executeOnEmployee(e: Employee){
        e.setVacation(e.getVacation()-2)
        console.log(`Employee ${e.name} visited by VacationSucker`)
    }
    public executeOnManager(m :Manager){
        m.setVacation(m.getVacation()-3)
        console.log(`Manager ${m.name} visited by VacationSucker`)
    }
}

let salaryGiver = new SalaryGiver()
let vacationReaper = new VacationSucker()

let workers = [new Employee("Jon",50000,30)
,new Employee("Bob",150000,32)
,new Employee("Mary",100000,10)
,new Manager("Jennifer",50000000,20)];


for (let e of workers){
    console.log(`${e.name} salary was ${e.getSalary()}`);
    e.accept(salaryGiver)
    console.log(`${e.name} salary is ${e.getSalary()}`);
}
console.log("---------------------");
for (let e of workers){
    console.log(`${e.name} vacationd was ${e.getVacation()}`);
    e.accept(vacationReaper)
    console.log(`${e.name} vacation was ${e.getVacation()}`);
}

