package main

import (
	"errors"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Account struct {
	ID      uint
	Balance int
}

type Transaction struct {
	ID            uint
	FromAccountId uint
	ToAccountId   uint
	Amount        int
}

func main() {
	dsn := "root:rootroot@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接mysql失败: %v", err)
	}

	err = db.AutoMigrate(&Account{}, &Transaction{})
	if err != nil {
		log.Fatalf("建表失败, %v", err)
	}

	//定义 转账账户ID、转账金额
	fromAccountId := uint(1)
	toAccountId := uint(2)
	amount := 100

	//新增两条记录
	accountA := &Account{ID: fromAccountId, Balance: 200}
	accountB := &Account{ID: toAccountId, Balance: 10}
	accounts := []*Account{accountA, accountB}
	db.Create(accounts)

	db.Transaction(func(tx *gorm.DB) error {
		var account1 Account
		db.Debug().First(&account1, fromAccountId)

		if account1.Balance < amount {
			return errors.New("余额不足，无法转账")
		}

		//更新A的余额
		db.Debug().Model(&account1).Update("balance", account1.Balance-amount)

		//更新B的余额
		var account2 Account
		db.Debug().First(&account2, toAccountId)
		db.Debug().Model(&account2).Update("balance", account2.Balance+amount)

		//保存转账信息
		transaction := Transaction{
			FromAccountId: fromAccountId,
			ToAccountId:   toAccountId,
			Amount:        amount,
		}
		db.Debug().Create(&transaction)
		return nil
	})

}
