package main

import (
	"fmt"
	"github.com/dongmyd/Code_Zinx/myDemo/protobufDemo/pb"
	"github.com/golang/protobuf/proto"
)

func main() {
	//定义一个Person结构对象
	person := &pb.Person{
		Name:   "Wuyazi",
		Age:    16,
		Emails: []string{"danbing.at@gmail.com", "danbing_at@163.com"},
		Phones: []*pb.PhoneNumber{
			&pb.PhoneNumber{
				Number: "1313131333",
				Type:   pb.PhoneType_MOBILE,
			},
			&pb.PhoneNumber{
				Number: "1313131333",
				Type:   pb.PhoneType_HOME,
			},
			&pb.PhoneNumber{
				Number: "1313131333",
				Type:   pb.PhoneType_WORK,
			},
		},
	}

	//编码
	//将person对象，就是将protobuf的message进行序列化，得到一个二进制文件
	data, err := proto.Marshal(person)
	//data 就是我们要进行网络传输的数据，对端需要按照message person格式进行解析
	if err != nil {
		fmt.Println("marshal err: ", err)
	}

	//解码
	newdata := &pb.Person{}
	err = proto.Unmarshal(data, newdata)
	if err != nil {
		fmt.Println("Unmarshal err: ", err)
	}

	fmt.Println("源数据：", person)
	fmt.Println("解码后数据：", newdata)
}
