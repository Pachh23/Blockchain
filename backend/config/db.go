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
		&entity.Departments{},
		&entity.Room{}, // เพิ่มการ migrate ตาราง Like
		&entity.Times{},
		&entity.Appointment{},
	)

	// ผู้ป่วย
	patients := []entity.Patient{
		{
			NationalID: "11111",
			FirstName:  "สมหญิง",
			LastName:   "มีสุข",
			Gender:     "หญิง",
			Age:        21,
		},
		{
			NationalID: "22222",
			FirstName:  "สมชาย",
			LastName:   "ใจดี",
			Gender:     "ชาย",
			Age:        60,
		},
		{
			NationalID: "33333",
			FirstName:  "สมศักดิ๋",
			LastName:   "แสงชัย",
			Gender:     "ชาย",
			Age:        49,
		},
		{
			NationalID: "44444",
			FirstName:  "เจน",
			LastName:   "แก้วแก้ม",
			Gender:     "หญิง",
			Age:        25,
		},
		{
			NationalID: "55555",
			FirstName:  "นุ่น",
			LastName:   "สุทธิ",
			Gender:     "หญิง",
			Age:        29,
		},
		{
			NationalID: "66666",
			FirstName:  "โบว์",
			LastName:   "แสงจ้า",
			Gender:     "หญิง",
			Age:        30,
		},
		{
			NationalID: "77777",
			FirstName:  "บี",
			LastName:   "ศรีสุข",
			Gender:     "หญิง",
			Age:        31,
		},
		{
			NationalID: "88888",
			FirstName:  "เอ",
			LastName:   "แสงดาว",
			Gender:     "หญิง",
			Age:        21,
		},
		{
			NationalID: "99999",
			FirstName:  "เอเค",
			LastName:   "สุขสันต์",
			Gender:     "ชาย",
			Age:        44,
		},
		{
			NationalID: "00000",
			FirstName:  "ยู",
			LastName:   "มนตรี",
			Gender:     "ชาย",
			Age:        50,
		},
	}

	for _, patient := range patients {
		db.FirstOrCreate(&patient, entity.Patient{NationalID: patient.NationalID})
	}

	// นัดหมาย
	appointments := []entity.Appointment{
		{
			Date:         time.Now(),
			TimeID:       1,
			Illness:      "ปวดหัว",
			DepartmentID: 1,
			RoomID:       1,
		},
		{
			Date:         time.Now(),
			TimeID:       1,
			Illness:      "ปวดตัว",
			DepartmentID: 2,
			RoomID:       1,
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
			Room:         "ห้องตรวจสำหรับเด็ก 1",
			DepartmentID: 1,
		},
		{
			Room:         "ห้องตรวจสำหรับเด็ก 2",
			DepartmentID: 1,
		},
		{
			Room:         "ห้องตรวจสำหรับเด็ก 3",
			DepartmentID: 1,
		},
		// แผนกทันตกรรม
		{
			Room:         "ห้องตรวจฟัน 1",
			DepartmentID: 2,
		},
		{
			Room:         "ห้องตรวจฟัน 2",
			DepartmentID: 2,
		},
		{
			Room:         "ห้องตรวจฟัน 3",
			DepartmentID: 2,
		},
		// แผนกศัลยกรรม
		{
			Room:         "ห้องตรวจศัลยกรรม 1",
			DepartmentID: 3,
		},
		{
			Room:         "ห้องตรวจศัลยกรรม 2",
			DepartmentID: 3,
		},
		{
			Room:         "ห้องตรวจศัลยกรรม 3",
			DepartmentID: 3,
		},
		// แผนกศัลยกรรมและกระดูก
		{
			Room:         "ห้องตรวจกระดูก 1",
			DepartmentID: 4,
		},
		{
			Room:         "ห้องตรวจกระดูก 2",
			DepartmentID: 4,
		},
		{
			Room:         "ห้องตรวจกระดูก 3",
			DepartmentID: 4,
		},
		// แผนกหู คอ จมูก
		{
			Room:         "ห้องตรวจหู คอ จมูก 1",
			DepartmentID: 5,
		},
		{
			Room:         "ห้องตรวจหู คอ จมูก 2",
			DepartmentID: 5,
		},
		{
			Room:         "ห้องตรวจหู คอ จมูก 3",
			DepartmentID: 5,
		},
		// แผนกอายุรกรรม
		{
			Room:         "ห้องตรวจอายุรกรรม 1",
			DepartmentID: 6,
		},
		{
			Room:         "ห้องตรวจอายุรกรรม 2",
			DepartmentID: 6,
		},
		{
			Room:         "ห้องตรวจอายุรกรรม 3",
			DepartmentID: 6,
		},
		// แผนกผิวหนัง
		{
			Room:         "ห้องตรวจผิวหนัง 1",
			DepartmentID: 7,
		},
		{
			Room:         "ห้องตรวจผิวหนัง 2",
			DepartmentID: 7,
		},
		{
			Room:         "ห้องตรวจผิวหนัง 3",
			DepartmentID: 7,
		},
	}

	// เพิ่มห้องในฐานข้อมูล
	for _, room := range rooms {
		db.FirstOrCreate(&room, entity.Room{DepartmentID: room.DepartmentID, Room: room.Room})
	}
	// แผนก
	times := []entity.Times{
		{
			Time: "10.00 - 11.00",
		},
		{
			Time: "11.00 - 12.00",
		},
		{
			Time: "12.00 - 13.00",
		},
		{
			Time: "13.00 - 14.00",
		},
		{
			Time: "14.00 - 15.00",
		},
		{
			Time: "15.00 - 16.00",
		},
	}

	for _, time := range times {
		db.FirstOrCreate(&time, entity.Times{Time: time.Time})
	}

}
