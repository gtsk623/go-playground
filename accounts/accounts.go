package accounts

import "errors"

//Account struct
//Func는 대문자로 시작해서 외부에서 접근이 가능하나, 안에 값들은 소문자로 시작하기 때문에 export 불가능
type Account struct {
	owner   string
	balance int
}

//error를 관리하기 위해 변수를 선언해서 return 시킴
var errNoMoney = errors.New("Can't withdraw")

//아래와 같은 코드를 Receiver라고 함
//아래와 같이 작성하면 account는 Deposit func를 가지게 됨
//a의 타입은 Account
//receiver의 이름을 정할 때는 struct의 첫글자의 소문자로 지어야 함
//아래에서 receiver를 선언 할 때 Account가 아닌 *Account를 쓰면 복사본이 아닌 Deposit method를 호출한 Account를 사용하라는 뜻
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

//balance만 출력하는 func
func (a Account) Balance() int {
	return a.balance
}

//에러를 핸들링할 수 있는 func 생성
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

//아래에서 *Account는 Account를 pointer하고 있다는 뜻
//NewAccount : 새로운 계좌를 생성하는 func
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}
