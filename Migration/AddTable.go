package Migration

import "github.com/qiankaihua/ginDemo/Model"

func AddTable()  {
	InitMigration()
	AddMigration("user", &Model.Test{})
	AddMigration("game", &Model.Test{})
	AddMigration("download", &Model.Test{})
}
