// https://refactoring.guru/design-patterns/facade
// https://wwthis.dofactory.com/javascript/design-patterns/facade


// Facade
interface iWalletFacade {
    account:      Account;
    wallet:       Wallet;
    securityCode: SecurityCode;
    notification: Note;
    ledger:       Ledger;
}
class WalletFacade {
    protected walletFacade: iWalletFacade;
    constructor(accountID:string, code:number){
        this.walletFacade = {
            account:  new Account(accountID),
            wallet:       new Wallet(),
            securityCode: new SecurityCode(code),
            notification: new Note(),
            ledger:       new Ledger()
        }
        console.log("Account created")
    }
    addMoneyToWallet(accountID:string,securityCode:number,amount:number){
        console.log("Starting add money to wallet")
        let err = this.walletFacade.account.checkAccount(accountID)
        if (err != null) {
            return err
        }
        err = this.walletFacade.securityCode.checkCode(securityCode)
        if (err != null) {
            return err
        }
        this.walletFacade.wallet.creditBalance(amount)
        this.walletFacade.notification.sendWalletCreditNotification()
        this.walletFacade.ledger.makeEntry(accountID, "credit", amount)
    }

    deductMoneyFromWallet(accountID:string,securityCode:number,amount:number){
        console.log("Starting debit money from wallet")
        let err = this.walletFacade.account.checkAccount(accountID)
        if (err != null) {
            return err
        }

        err = this.walletFacade.securityCode.checkCode(securityCode)
        if (err != null) {
            return err
        }
        err = this.walletFacade.wallet.debitBalance(amount)
        if (err != null) {
            return err
        }
        this.walletFacade.notification.sendWalletDebitNotification()
        this.walletFacade.ledger.makeEntry(accountID, "credit", amount)
    }
}


// subsystem parts
class Account{
    protected name :string;
    constructor(name:string){
        this.name = name;
    }
    checkAccount(accountName: string) {
        if (this.name != accountName) {
            return console.error("Account Name is incorrect")
        }
        console.log("Account Verified")
    }
}
class SecurityCode{
    protected code :number;
    constructor(code:number){
        this.code =code;
    }
   checkCode(incomingCode: number)  {
        if (this.code != incomingCode) {
            return console.error("Security Code is incorrect")
        }
        console.log("SecurityCode Verified")
    }
}
class Wallet {
    protected balance:number
    constructor(){
        this.balance=0;
    }
    creditBalance(amount :number) {
        this.balance += amount
        console.log("Wallet balance added successfully")
    }
    
    debitBalance(amount :number) {
        if (this.balance < amount) {
            return console.error("Balance is not sufficient")
        }
        console.log("Wallet balance is Sufficient")
        this.balance = this.balance - amount
    }
}
class Ledger {
    constructor(){}
    makeEntry(accountID:string, txnType: string, amount: number) {
        console.log(`Make ledger entry for accountId ${accountID} with txnType ${txnType} for amount ${amount}\n`)
    }
}
class Note {
    constructor(){}
    sendWalletCreditNotification() {
        console.log("Sending wallet credit notification")
    }
    sendWalletDebitNotification() {
        console.log("Sending wallet debit notification")
    }
}

// client code 
const walletFacade = new WalletFacade("abc", 1234)
let err = walletFacade.addMoneyToWallet("abc", 1234, 10)
if (err) {
    console.error(`Error: ${err}`)
}
err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
if (err) {
    console.error(`Error: ${err}`)
}
