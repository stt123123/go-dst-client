package proto

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"testing"
)

type writer struct {
	w *bufio.Writer
	
}

//func (w *writer) writeArgs()  {
//	var req = &Request{
//		RequestId:
//
//	}
//
//}

//func generateId(){
//	t := time.Now().Unix()
//}

func  Test_a_test(t *testing.T) {

	conn,err := grpc.Dial("127.0.0.1:8080",grpc.WithInsecure())
	if err != nil{
		fmt.Println(err)
	}
	defer conn.Close()
	//client := proto.NewHelloClient(conn)    //  自动生成方法, 因为proto文件中service的名字是hello
	//name := "jhonsmith"
	//result, err := client.Hello(context.Background(), &proto.Request{Name:name})   // 调用grpc方法, 对服务端进行通讯
	//if err != nil{
	//	fmt.Println(err)
	//}
	//fmt.Println("I see, u see. I say ", name, "u say ", result.Msg)

	client := NewDstStringServiceClient(conn)
	p,err := client.Put(context.Background(),&PutRequest{Key:"a",Value:"avalue"})
	if err != nil{
		fmt.Println(err)
	}
	a,err := json.Marshal(p)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(a)

}