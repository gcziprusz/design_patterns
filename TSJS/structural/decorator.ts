// https://refactoring.guru/design-patterns/decorator/typescript/example 
// https://www.dofactory.com/javascript/design-patterns/decorator 

// Component interface
interface IDryer {
    dry(): string
  }
  
  // concrete component 
  class Dryer implements IDryer {
    public dry() {
      return "simple dry."
    }
  }
  
  // Base decoraotor class
  class DryerDecorator implements IDryer {
     protected dryer: Dryer;
  
     constructor(dryer: Dryer){
       this.dryer = dryer;
     }
  
    public dry(){
       return this.dryer.dry();
    }
  }
  
  // Concrete Decorators 
  class ElectricDryer extends DryerDecorator {
    public dry(){
      return "dry using electicity." + super.dry()
    }
  }
  class HeatDryer extends DryerDecorator {
    public dry(){
      return "dry with heath." + super.dry()
    }
  }
  class SteamDryer extends DryerDecorator {
    public dry(){
      return "dry with steam." + super.dry()
    }
  }
  
  // client code
  function clientCode(dryer: IDryer) {
      console.log(`RESULT: ${dryer.dry()}`);
  }
  
  const simple = new Dryer();
  clientCode(simple);
  
  const electricDryer = new ElectricDryer(simple);
  const electricHeatDryer = new HeatDryer(electricDryer);
  const electricHeatSteamDryer = new SteamDryer(electricHeatDryer);
  clientCode(electricHeatSteamDryer);


