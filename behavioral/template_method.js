let datastore = {
  process: function () {
    this.open();
    this.log();
    this.select();
    this.log();
    this.close();
  },
  log() {
    // optional hook
  },
  select() {
    console.log("Run default implementation of select operation");
  },
};

let inherit = function (proto) {
  let F = function () {};
  F.prototype = proto;
  return new F();
};

let mySQL = inherit(datastore);

mySQL.open = function () {
  console.log("Run open operation in MySQL flavor");
};
mySQL.select = function () {
  console.log("Run select operation in MySQL flavor");
};
mySQL.close = function () {
  console.log("Run close operation in MySQL flavor");
};
let oracle = inherit(datastore);

oracle.open = function () {
  console.log("Run open operation in Oracle flavor");
};
oracle.close = function () {
  console.log("Run close operation in Oracle flavor");
};
oracle.log = function () {
  console.log("Run log operation in Oracle flavor");
};

function clientCode(db) {
  db.process();
}

clientCode(mySQL);
console.log("--------------------------\n");
clientCode(oracle);
