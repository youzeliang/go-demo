package main

//
import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/balancer/base"
	"strconv"
	"time"
	"github.com/sony/sonyflake"

)


type FrameWorkOrderRelation struct {
	ID           int64     `gorm:"column:id;primary_key" json:"id"`
	WorkOrderID  string    `gorm:"column:work_order_id" json:"workOrderID"`
	WarehouseID  string    `gorm:"column:warehouse_id" json:"warehouseID"`
	ActionWaveID string    `gorm:"column:action_wave_id" json:"actionWaveID"`
	FrameID      string    `gorm:"column:frame_id" json:"frameID"`
	Status       int       `gorm:"column:status" json:"status"`
	TargetType   int       `gorm:"column:target_type" json:"targetType"`
	CreateTime   time.Time `gorm:"column:create_time" json:"-"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"-"`
}

var flake = sonyflake.NewSonyflake(sonyflake.Settings{})


func GetUuidUInt64() uint64 {
	uid, _ := flake.NextID()
	return uid
}

func GetUuidString() string {
	return strconv.FormatUint(GetUuidUInt64(), 10)
}

func main() {

	r := gin.New()

	userName:="root"
	password:="Lyz123456"
	addr:="rm-bp1lck87254dnsaelvo.mysql.rds.aliyuncs.com"
	name:="warelogic"

	config := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		userName,
		password,
		addr,
		name,
		true,
		// "Asia%2FShanghai",  // 必须是 url.QueryEscape 的
		"Local",
	)
	_, err := gorm.Open("mysql", config)
	if err != nil {
		logrus.Fatalf("数据库连接失败. 数据库名字: %s. 错误信息: %s", name, err)
	} else {
		logrus.Infof("数据库连接成功, 数据库名字: %s", name)
	}

	for range time.Tick(time.Millisecond*100){
		//fmt.Println("Hello TigerwolfC")



		now := time.Now()
		relation := FrameWorkOrderRelation{
			WarehouseID:  GetUuidString(),
			ActionWaveID: GetUuidString(),
			FrameID:      GetUuidString(),
			Status:       1,
			TargetType:   1,
			CreateTime:   now,
			UpdateTime:   now,
		}

		err = relation.Inserta(ctx)
		if err != nil {
			return err
		}



	}


	endless.ListenAndServe(":"+"8080", r)


}
//
//func pump1(ch chan int) {
//	for i := 0; ; i++ {
//		ch <- i * 2
//		time.Sleep(time.Duration(time.Second))
//	}
//}
//
//func pump2(ch chan string) {
//	for i := 0; ; i++ {
//		ch <- strconv.Itoa(i+5)
//		time.Sleep(time.Duration(time.Second))
//	}
//}
//
//func suck(ch1 chan int, ch2 chan string) {
//	chRate := time.Tick(time.Duration(time.Second*5)) // 定时器
//	for {
//		select {
//		case v := <-ch1:
//			fmt.Printf("Received on channel 1: %d\n", v)
//		case v := <-ch2:
//			fmt.Printf("Received on channel 2: %s\n", v)
//		case <-chRate:
//			fmt.Printf("Log log...\n")
//		}
//	}
//}