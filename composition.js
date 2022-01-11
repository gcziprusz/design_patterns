function swimmer({ name }) {
  return {
    swim() {
      console.log(`${name} is swimming`);
    },
  };
}
function flyer({ name }) {
  return {
    fly() {
      console.log(`${name} is flying`);
    },
  };
}
function walker({ name }) {
  return {
    walk() {
      console.log(`${name} is walking`);
    },
  };
}
function attacker({ name }) {
  return {
    attack() {
      console.log(`${name} is attacking`);
    },
  };
}

function makeCreature(name) {
  return { name };
}
function createSwimmerAttacker(name) {
  let creature = makeCreature(name);
  return {
    ...creature,
    ...swimmer(creature),
    ...attacker(creature),
  };
}
function createSwimmerWalkerAttacker(name) {
  let creature = makeCreature(name);
  return {
    ...creature,
    ...swimmer(creature),
    ...attacker(creature),
    ...walker(creature),
  };
}
function createFlyer(name) {
  let creature = makeCreature(name);
  return {
    ...creature,
    ...flyer(creature),
  };
}

let shark = createSwimmerAttacker("big blue");
shark.swim();
shark.attack();

let bear = createSwimmerWalkerAttacker("grumpy salmon snatcher");
bear.swim();
bear.attack();
bear.walk();

let bird = createFlyer("Redhair Cardy");
bird.fly();
try {
  bird.attack();
} catch (e) {
  console.error(
    `${bird.name} is a friendly creature it won't attack! ERROR: ${e}`
  );
}
