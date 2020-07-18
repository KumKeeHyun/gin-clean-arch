package orm

import (
	"github.com/KumKeeHyun/gin-clean-arch/app/domain/model"
	"github.com/KumKeeHyun/gin-clean-arch/app/setting"
	"github.com/jinzhu/gorm"
)

var dbConn *gorm.DB

func Setup() {
	driver, conn := setting.Databasesetting.MakeConnection()
	dbConn, _ = gorm.Open(driver, conn)

	dbConn.AutoMigrate(
		&model.Node{},
		&model.Sensor{},
		&model.NodeSensor{},
		&model.SensorValue{},
	)
	dbConn.Model(&model.NodeSensor{}).AddForeignKey("node_uuid", "nodes(uuid)", "CASCADE", "CASCADE")
	dbConn.Model(&model.NodeSensor{}).AddForeignKey("sensor_uuid", "sensors(uuid)", "CASCADE", "CASCADE")
	dbConn.Model(&model.SensorValue{}).AddForeignKey("sensor_uuid", "sensors(uuid)", "CASCADE", "CASCADE")
}
