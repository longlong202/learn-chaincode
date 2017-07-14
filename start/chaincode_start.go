/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	// "encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
	// "strings"
)

// HelloChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// type (
// 	Infoo struct {
// 		Transfer string `json:"transfer"`
// 		Receiver string `json:"receiver"`
// 		Current  string `json:"current"`
// 	}

// 	Status struct {
// 		Money   int    `json:"money"`
// 		Message string `json:"message"`
// 		Info    Infoo  `json:"info"`
// 	}
// )

// func (st *Status) ShowStatus() {
// 	fmt.Println("show status :")
// 	fmt.Println("\tmoney\t:", st.Money)
// 	fmt.Println("\tmessage\t:", st.Message)
// 	fmt.Println("\tinfo\n\t\ttransfer\t:", st.Info.Transfer)
// 	fmt.Println("\t\treceiver\t:", st.Info.Receiver)
// 	fmt.Println("\t\tcurrent\t:", st.Info.Current)
// }

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	fmt.Printf("HelloWorld - Init called with function %s!\n", function)
	fmt.Println("args:", args)

	return nil, nil
}

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	// Handle different functions
	// if function == "echo" {
	// 	str := strings.Join(args, " ")
	// 	fmt.Printf("echo: %s\n", str)
	// 	return []byte(str), nil
	// }

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}

	if function == "init" { //参数形式如下：args["init","zhangsan","1000"]
		//args:["zhangsan","1000"]
		fmt.Println("资金初始化")
		//字符串转换成int，使用strconv.Atoi
		val, err := strconv.Atoi(args[1])
		if err != nil {
			return nil, err
		}

		// str, err := t.PackInitJson(stub, args[0], val)
		// if err != nil {
		// 	return nil, err
		// }

		//将该json存入数据库，然后再从里面读取出数据
		fmt.Println("将状态写入数据库中.....")
		err = stub.PutState(args[0], []byte(strconv.Itoa(val)))
		if err != nil {
			return nil, err
		}

		//--------------------------第一阶段---------------------------

		//将资产交易后的数据填充到新的状态

		return nil, nil

	}

	// if function == "transfer" { //参数如下形式如下：["transfer","zhangsan","Lisi","50"]
	// 	byteval, err := t.transfer(stub, args)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return byteval, nil
	// }

	// newmoney := newst.Money + 200 //获取到的是int值，可直接进行资产交易
	//把int值转换成字符串
	// fmt.Println("资产交易后的钱：" + strconv.Itoa(newmoney))

	fmt.Println("invoke did not find func: " + function) //error

	return nil, nil
}

//需要对应的fuc名字对应该方法
// func (t *SimpleChaincode) transfer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
// 	//args:["zhangsan","Lisi","50"]
// 	/**
// 	判断参数情况在此写！
// 	**/
// 	fmt.Println("资金交易开始.....")

// 	//从数据库中读取该用户的状态
// 	Aval, err := stub.GetState(args[0]) //取出要转账的用户的状态
// 	if err != nil {
// 		return nil, errors.New("Failed to get state")
// 	}
// 	if Aval == nil {
// 		return nil, errors.New("Entity not found")
// 	}

// 	fmt.Println("从数据库中获取转账人的状态...")

// 	//得到的是字节数组，再从字节数组中获取相应的值进行资产运算，以及填充新的JSON
// 	cha := make(chan string, 1)          //创建有缓存通道
// 	go func(c chan string, str string) { //形参
// 		c <- str
// 	}(cha, string(Aval)) //实参

// 	strData := <-cha
// 	fmt.Println("----------------------")
// 	newst := &Status{} //创建封装格式的结构对象

// 	err = json.Unmarshal([]byte(strData), &newst) //反序列化会返回一个错误判断
// 	if err != nil {
// 		fmt.Println("反序列化失败！")
// 	}
// 	fmt.Println("转账人状态反序列化成功！")
// 	fmt.Println("资产交易前的转账人状态：")
// 	newst.ShowStatus() //展示从数据库中获取的数据的对象状态

// 	//以上代码取出了转账人的资金

// 	//取出收款人的状态
// 	Bval, err := stub.GetState(args[1]) //取出收款人的状态
// 	if err != nil {
// 		return nil, errors.New("Failed to get state")
// 	}
// 	if Bval == nil {
// 		return nil, errors.New("Entity not found")
// 	}

// 	fmt.Println("从数据库中获取收款人状态...")

// 	//得到的是字节数组，再从字节数组中获取相应的值进行资产运算，以及填充新的JSON
// 	chb := make(chan string, 1)          //创建有缓存通道
// 	go func(c chan string, str string) { //形参
// 		c <- str
// 	}(chb, string(Bval)) //实参

// 	strDatab := <-chb
// 	fmt.Println("----------------------")
// 	newstb := &Status{} //创建封装格式的结构对象

// 	err = json.Unmarshal([]byte(strDatab), &newstb) //反序列化会返回一个错误判断
// 	if err != nil {
// 		fmt.Println("反序列化失败！")
// 	}
// 	fmt.Println("反序列化成功！")
// 	fmt.Println("资产交易前的收款人状态：")
// 	newstb.ShowStatus() //展示从数据库中获取的数据的对象状态

// 	//进行资产交易
// 	moneyval, err := strconv.Atoi(args[2])
// 	if err != nil {
// 		return nil, err
// 	}
// 	if newst.Money < moneyval || newst.Money < 0 {
// 		return nil, errors.New("资金不足，无法完成交易！")
// 	}

// 	newst.Money = newst.Money - moneyval //转账人扣除相应金额

// 	newstb.Money = newstb.Money + moneyval //收款人增加相应的金额

// 	//封装转账人以及收款人
// 	Ajson, err := t.PackTransferJson(stub, newst.Money, args[0], args[1], moneyval) //封装转账人
// 	err = stub.PutState(args[0], Ajson)
// 	if err != nil {
// 		return nil, err
// 	}

// 	Bjson, err := t.PackTransferJson(stub, newstb.Money, args[0], args[1], moneyval) //封装收款人
// 	err = stub.PutState(args[1], Bjson)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return nil, nil
// }

// func (t *SimpleChaincode) PackInitJson(stub shim.ChaincodeStubInterface, name string, money int) ([]byte, error) {

// 	txid := stub.GetTxID()
// 	fmt.Printf("txid: %s\n", txid)

// 	st := &Status{
// 		money,
// 		name + " initial",
// 		Infoo{
// 			txid,
// 			txid,
// 			txid,
// 		},
// 	}
// 	jobj, err := json.Marshal(st)
// 	if err != nil {
// 		fmt.Println("编码失败！")
// 		return nil, err
// 	} else {
// 		fmt.Println("编码数据：")
// 		// fmt.Println(jobj)//字节数组格式
// 		fmt.Println(string(jobj)) //转换成字符串之后的才是所需要的格式：{"name":100,"message":"A initial","info":{"transfer":"uuid-transfer","receiver":"uuid-receiver"}}
// 		return jobj, nil
// 	}
// }

//获取用户资产交易前的状态里的st.Info.Current
// func (t *SimpleChaincode) GetUUID(stub shim.ChaincodeStubInterface, name string) (string, error) {
// 	//取出用户的状态
// 	val, err := stub.GetState(name) //取出收款人的状态
// 	if err != nil {
// 		return "", errors.New("Failed to get state")
// 	}
// 	if val == nil {
// 		return "", errors.New("Entity not found")
// 	}

// 	fmt.Println("从数据库中获取用户状态...")

// 	//得到的是字节数组，再从字节数组中获取相应的值进行资产运算，以及填充新的JSON
// 	ch := make(chan string, 1)           //创建有缓存通道
// 	go func(c chan string, str string) { //形参
// 		c <- str
// 	}(ch, string(val)) //实参

// 	strData := <-ch
// 	fmt.Println("----------------------")
// 	newst := &Status{} //创建封装格式的结构对象

// 	err = json.Unmarshal([]byte(strData), &newst) //反序列化会返回一个错误判断
// 	if err != nil {
// 		fmt.Println("反序列化失败！")
// 	}
// 	fmt.Println("反序列化成功！")
// 	// fmt.Println("资产交易前的收款人状态：")
// 	newst.ShowStatus() //展示从数据库中获取的数据的对象状态
// 	fmt.Println("获取用户资产交易前的UUID：")
// 	fmt.Println(newst.Info.Current)
// 	return newst.Info.Current, nil

// }

// func (t *SimpleChaincode) PackTransferJson(stub shim.ChaincodeStubInterface, lastmoney int, transfer string, receive string, money int) ([]byte, error) {

// 	txid := stub.GetTxID()
// 	fmt.Printf("txid: %s\n", txid)

// 	//获取transfer用户的st.Info.Current
// 	transferUUID, err := t.GetUUID(stub, transfer)
// 	if err != nil {
// 		return nil, errors.New("Entity not found")
// 	}
// 	//获取receive用户的st.Info.Current
// 	receiveUUID, err := t.GetUUID(stub, receive)
// 	if err != nil {
// 		return nil, errors.New("Entity not found")
// 	}

// 	st := &Status{
// 		lastmoney,
// 		transfer + " transfer to " + receive + " " + strconv.Itoa(money),
// 		Infoo{
// 			transferUUID,
// 			receiveUUID,
// 			txid,
// 		},
// 	}
// 	jobj, err := json.Marshal(st)
// 	if err != nil {
// 		fmt.Println("编码失败！")
// 		return nil, err
// 	} else {
// 		fmt.Println("编码数据：")
// 		// fmt.Println(jobj)//字节数组格式
// 		fmt.Println(string(jobj)) //转换成字符串之后的才是所需要的格式：{"name":100,"message":"A initial","info":{"transfer":"uuid-transfer","receiver":"uuid-receiver"}}
// 		return jobj, nil
// 	}
// }

//只查询钱数
// func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, name string) ([]byte, error) {
// 	//取出用户的状态
// 	val, err := stub.GetState(name) //取出收款人的状态
// 	if err != nil {
// 		return nil, errors.New("Failed to get state")
// 	}
// 	if val == nil {
// 		return nil, errors.New("Entity not found")
// 	}

// 	fmt.Println("从数据库中获取用户状态...")

// 	//得到的是字节数组，再从字节数组中获取相应的值进行资产运算，以及填充新的JSON
// 	ch := make(chan string, 1)           //创建有缓存通道
// 	go func(c chan string, str string) { //形参
// 		c <- str
// 	}(ch, string(val)) //实参

// 	strData := <-ch
// 	fmt.Println("----------------------")
// 	newst := &Status{} //创建封装格式的结构对象

// 	err = json.Unmarshal([]byte(strData), &newst) //反序列化会返回一个错误判断
// 	if err != nil {
// 		fmt.Println("反序列化失败！")
// 	}
// 	fmt.Println("反序列化成功！")
// 	// fmt.Println("资产交易前的收款人状态：")
// 	newst.ShowStatus() //展示从数据库中获取的数据的对象状态
// 	fmt.Println("获取" + name + "的余额：")
// 	fmt.Println(newst.Money)
// 	return []byte(strconv.Itoa(newst.Money)), nil
// }

// func (t *SimpleChaincode) uuid(stub shim.ChaincodeStubInterface, name string) ([]byte, error) {
// 	//取出用户的状态
// 	val, err := stub.GetState(name) //取出收款人的状态
// 	if err != nil {
// 		return nil, errors.New("Failed to get state")
// 	}
// 	if val == nil {
// 		return nil, errors.New("Entity not found")
// 	}

// 	fmt.Println("从数据库中获取用户状态...")

// 	//得到的是字节数组，再从字节数组中获取相应的值进行资产运算，以及填充新的JSON
// 	ch := make(chan string, 1)           //创建有缓存通道
// 	go func(c chan string, str string) { //形参
// 		c <- str
// 	}(ch, string(val)) //实参

// 	strData := <-ch
// 	fmt.Println("----------------------")
// 	newst := &Status{} //创建封装格式的结构对象

// 	err = json.Unmarshal([]byte(strData), &newst) //反序列化会返回一个错误判断
// 	if err != nil {
// 		fmt.Println("反序列化失败！")
// 	}
// 	fmt.Println("反序列化成功！")
// 	// fmt.Println("资产交易前的收款人状态：")
// 	newst.ShowStatus() //展示从数据库中获取的数据的对象状态
// 	fmt.Println("获取用户的UUID：")
// 	fmt.Println(newst.Info.Current)
// 	return []byte(newst.Info.Current), nil
// }

// Query is our entry point for queries
// func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
// 	// fmt.Println("query is running " + function)
// 	if function == "query" { //args:[name]
// 		Avalbytes, err := t.query(stub, args[0])
// 		if err != nil {
// 			jsonResp := "{\"Error\":\"Failed to get state for " + args[0] + "\"}"
// 			return nil, errors.New(jsonResp)
// 		}

// 		return Avalbytes, nil
// 	}

// 	if function == "uuid" { //args:[name]
// 		Avalbytes, err := t.uuid(stub, args[0])
// 		if err != nil {
// 			jsonResp := "{\"Error\":\"Failed to get state for " + args[0] + "\"}"
// 			return nil, errors.New(jsonResp)
// 		}

// 		return Avalbytes, nil
// 	}

// Handle different functions
// if function == "txid" {
// 	txid := stub.GetTxID()
// 	fmt.Printf("txid: %s\n", txid)
// 	return []byte(txid), nil
// }

// if function == "args" {
// 	arg := stub.GetArgs()
// 	fmt.Println("args:", arg)
// 	return nil, nil
// }

// if function == "strings" {
// 	str := stub.GetStringArgs()
// 	fmt.Println("sargs:", str)
// 	return []byte(strings.Join(str, " ")), nil
// }

// if function == "meta" {
// 	meta, err := stub.GetCallerMetadata()
// 	fmt.Println("metadata:" + string(meta))
// 	return meta, err
// }

// if function == "bind" {
// 	bind, err := stub.GetBinding()
// 	fmt.Println("binding:" + string(bind))
// 	return bind, err
// }

// if function == "payload" {
// 	payload, err := stub.GetPayload()
// 	fmt.Println("binding:" + string(payload))
// 	return payload, err
// }

// if function == "timestamp" {
// 	time, err := stub.GetTxTimestamp()
// 	fmt.Println("timestamp:" + time.String())
// 	return []byte(time.String()), err
// }

// fmt.Println("query did not find func: " + function) //error

// 	return nil, errors.New("Received unknown function query: " + function)
// }

// Query callback representing the query of a chaincode
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function != "query" {
		return nil, errors.New("Invalid query function name. Expecting \"query\"")
	}
	var A string // Entities
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return nil, errors.New(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return nil, errors.New(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return Avalbytes, nil
}
