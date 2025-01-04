package config

import (
	"fmt"
	"time"

	"blockchain.com/bc-67/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("sa.db?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected database")
	db = database
}

func SetupDatabase() {
	db.AutoMigrate(
		&entity.Patient{},
		&entity.Appointment{},
		&entity.Departments{},
		&entity.Room{}, // เพิ่มการ migrate ตาราง Like
	)

	// ผู้ป่วย
	patients := []entity.Patient{
		{
			NationalID: "1111111111",
			FirstName:  "J",
			LastName:   "P",
			Gender:     "หญิง",
			Age:        21,
		},
		{
			NationalID: "2222222222222",
			FirstName:  "A",
			LastName:   "P",
			Gender:     "หญิง",
			Age:        21,
		},
	}

	for _, patient := range patients {
		db.FirstOrCreate(&patient, entity.Patient{NationalID: patient.NationalID})
	}

	// นัดหมาย
	appointments := []entity.Appointment{
		{
			Date:         time.Now(),
			Time:         time.Now(),
			Illness:      "ปวดหัว",
			DepartmentID: 1,
			PatientID:    1,
		},
		{
			Date:         time.Now(),
			Time:         time.Now(),
			Illness:      "ปวดตัว",
			DepartmentID: 2,
			PatientID:    2,
		},
	}

	for _, appointment := range appointments {
		db.FirstOrCreate(&appointment, entity.Appointment{Illness: appointment.Illness})
	}

	//แผนก
	departments := []entity.Departments{
		{
			Department:  "แผนกกุมารเวชกรรม",
			Description: "ดูแลและรักษาโรคในเด็ก ตั้งแต่แรกเกิดจนถึงวัยรุ่น",
			HeadOfDept:  "นพ.สมชาย เดชานุกูล",
		},
		{
			Department:  "แผนกทันตกรรม",
			Description: "ให้บริการตรวจและรักษาโรคเกี่ยวกับฟันและช่องปาก",
			HeadOfDept:  "ทพญ.ปรียาภรณ์ สุขสันต์",
		},
		{
			Department:  "แผนกศัลยกรรม",
			Description: "รับผิดชอบการผ่าตัดเพื่อรักษาโรคและอาการบาดเจ็บ",
			HeadOfDept:  "นพ.กิตติพงศ์ ศิริธรรม",
		},
		{
			Department:  "แผนกศัลยกรรมและกระดูก",
			Description: "ให้บริการตรวจและรักษาโรคกระดูกและข้อต่อ",
			HeadOfDept:  "นพ.อนันต์ ชาญชลวิทย์",
		},
		{
			Department:  "แผนกหู คอ จมูก",
			Description: "ดูแลการรักษาโรคเกี่ยวกับหู คอ จมูก และการได้ยิน",
			HeadOfDept:  "พญ.วิไลลักษณ์ แก้วกาญจน์",
		},
		{
			Department:  "แผนกอายุรกรรม",
			Description: "ดูแลและรักษาโรคในผู้ใหญ่ เช่น โรคเบาหวานและความดันโลหิตสูง",
			HeadOfDept:  "นพ.จิตติ ภัทราวงศ์",
		},
		{
			Department:  "แผนกผิวหนัง",
			Description: "ดูแลและรักษาโรคที่เกี่ยวกับผิวหนังและโรคภูมิแพ้",
			HeadOfDept:  "พญ.วรัญญา ผ่องใส",
		},
	}

	for _, department := range departments {
		db.FirstOrCreate(&department, entity.Departments{Department: department.Department})
	}

	// ห้องตรวจ
	rooms := []entity.Room{
		// แผนกกุมารเวชกรรม
		{
			Name:         "ห้องตรวจสำหรับเด็ก 1",
			DepartmentID: 1,
		},
		{
			Name:         "ห้องตรวจสำหรับเด็ก 2",
			DepartmentID: 1,
		},
		{
			Name:         "ห้องตรวจสำหรับเด็ก 3",
			DepartmentID: 1,
		},
		// แผนกทันตกรรม
		{
			Name:         "ห้องตรวจฟัน 1",
			DepartmentID: 2,
		},
		{
			Name:         "ห้องตรวจฟัน 2",
			DepartmentID: 2,
		},
		{
			Name:         "ห้องตรวจฟัน 3",
			DepartmentID: 2,
		},
		// แผนกศัลยกรรม
		{
			Name:         "ห้องตรวจศัลยกรรม 1",
			DepartmentID: 3,
		},
		{
			Name:         "ห้องตรวจศัลยกรรม 2",
			DepartmentID: 3,
		},
		{
			Name:         "ห้องตรวจศัลยกรรม 3",
			DepartmentID: 3,
		},
		// แผนกศัลยกรรมและกระดูก
		{
			Name:         "ห้องตรวจกระดูก 1",
			DepartmentID: 4,
		},
		{
			Name:         "ห้องตรวจกระดูก 2",
			DepartmentID: 4,
		},
		{
			Name:         "ห้องตรวจกระดูก 3",
			DepartmentID: 4,
		},
		// แผนกหู คอ จมูก
		{
			Name:         "ห้องตรวจหู คอ จมูก 1",
			DepartmentID: 5,
		},
		{
			Name:         "ห้องตรวจหู คอ จมูก 2",
			DepartmentID: 5,
		},
		{
			Name:         "ห้องตรวจหู คอ จมูก 3",
			DepartmentID: 5,
		},
		// แผนกอายุรกรรม
		{
			Name:         "ห้องตรวจอายุรกรรม 1",
			DepartmentID: 6,
		},
		{
			Name:         "ห้องตรวจอายุรกรรม 2",
			DepartmentID: 6,
		},
		{
			Name:         "ห้องตรวจอายุรกรรม 3",
			DepartmentID: 6,
		},
		// แผนกผิวหนัง
		{
			Name:         "ห้องตรวจผิวหนัง 1",
			DepartmentID: 7,
		},
		{
			Name:         "ห้องตรวจผิวหนัง 2",
			DepartmentID: 7,
		},
		{
			Name:         "ห้องตรวจผิวหนัง 3",
			DepartmentID: 7,
		},
	}

	// เพิ่มห้องในฐานข้อมูล
	for _, room := range rooms {
		db.FirstOrCreate(&room, entity.Room{DepartmentID: room.DepartmentID, Name: room.Name})
	}

}
